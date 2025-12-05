package desec

import (
	"context"
	"fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/certmagic"
	"github.com/libdns/desec"
	"github.com/libdns/libdns"
)

func init() {
	caddy.RegisterModule(Provider{})
}

type Provider struct {
	Token string `json:"token,omitempty"`
}

func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.desec",
		New: func() caddy.Module { return new(Provider) },
	}
}

//
// Caddy <-> CertMagic DNS Provider Implementation
//

func (p *Provider) Acquire(ctx context.Context, fqdn, token string) error {
	return p.Present(ctx, fqdn, token)
}

func (p *Provider) Release(ctx context.Context, fqdn, token string) error {
	return p.CleanUp(ctx, fqdn, token)
}

//
// libdns Provider Wrapper
//

func (p *Provider) Present(ctx context.Context, fqdn, token string) error {
	provider := &desec.Provider{Token: p.Token}

	record := libdns.TXT{
		Name: libdns.NameFromFQDN(fqdn),
		Text: token,
	}

	_, err := provider.AppendRecords(ctx, record.Zone(), []libdns.Record{record})
	return err
}

func (p *Provider) CleanUp(ctx context.Context, fqdn, token string) error {
	provider := &desec.Provider{Token: p.Token}

	record := libdns.TXT{
		Name: libdns.NameFromFQDN(fqdn),
		Text: token,
	}

	_, err := provider.DeleteRecords(ctx, record.Zone(), []libdns.Record{record})
	return err
}

var _ certmagic.DNSProvider = (*Provider)(nil)
