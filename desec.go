package desec

import (
	"context"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/certmagic"
	"github.com/libdns/desec"
	"github.com/libdns/libdns"
)

// Ensure Provider implements interfaces
var (
	_ certmagic.ACMEDNSProvider = (*Provider)(nil)
)

// Provider wraps the libdns/desec provider to make it usable in Caddy.
type Provider struct {
	Token string `json:"token,omitempty"`
}

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.desec",
		New: func() caddy.Module { return new(Provider) },
	}
}

func (p *Provider) getProvider() *desec.Provider {
	return &desec.Provider{
		Token: p.Token,
	}
}

// Present creates a DNS TXT record to complete DNS-01 challenge.
func (p *Provider) Present(domain, token, keyAuth string) error {
	zone := libdns.ZoneFromFQDN("_acme-challenge." + domain)

	_, err := p.getProvider().AppendRecords(context.Background(), zone, []libdns.Record{
		{
			Type:  "TXT",
			Name:  "_acme-challenge",
			Value: keyAuth,
		},
	})
	return err
}

// CleanUp removes the DNS TXT challenge record.
func (p *Provider) CleanUp(domain, token, keyAuth string) error {
	zone := libdns.ZoneFromFQDN("_acme-challenge." + domain)

	_, err := p.getProvider().DeleteRecords(context.Background(), zone, []libdns.Record{
		{
			Type:  "TXT",
			Name:  "_acme-challenge",
			Value: keyAuth,
		},
	})
	return err
}
