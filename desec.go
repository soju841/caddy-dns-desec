package desec

import (
	"context"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/certmagic"
	desecapi "github.com/libdns/desec"
	"github.com/libdns/libdns"
)

func init() {
	caddy.RegisterModule(Provider{})
}

// Provider implements the Caddy DNS provider interface for deSEC
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
// certmagic.DNSProvider interface (Caddy uses Present/CleanUp)
//

func (p *Provider) Present(ctx context.Context, fqdn, value string) error {
	api := &desecapi.Provider{Token: p.Token}

	name, zone := splitFQDN(fqdn)

	record := libdns.TXT{
		Name: name,
		Text: value,
	}

	_, err := api.AppendRecords(ctx, zone, []libdns.Record{record})
	return err
}

func (p *Provider) CleanUp(ctx context.Context, fqdn, value string) error {
	api := &desecapi.Provider{Token: p.Token}

	name, zone := splitFQDN(fqdn)

	record := libdns.TXT{
		Name: name,
		Text: value,
	}

	_, err := api.DeleteRecords(ctx, zone, []libdns.Record{record})
	return err
}

//
// Helper: split FQDN into name + zone
//

func splitFQDN(fqdn string) (name, zone string) {
	fqdn = strings.TrimSuffix(fqdn, ".")

	parts := strings.Split(fqdn, ".")
	if len(parts) < 3 {
		return "@", fqdn
	}

	zone = strings.Join(parts[len(parts)-2:], ".")
	name = strings.Join(parts[:len(parts)-2], ".")

	if name == "" {
		name = "@"
	}

	return name, zone
}

// interface assertion for Caddy
var _ certmagic.DNSProvider = (*Provider)(nil)
