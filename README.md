# Caddy DNS Provider for DeSEC (ACME DNS-01)

This plugin adds support for the ACME DNS-01 challenge in Caddy using the DeSEC DNS service via the libdns/desec provider.

With this plugin, Caddy can automatically create and clean up the required TXT records for certificate issuance through DeSEC.

---

## 🔧 Installation

Build Caddy with this plugin using `xcaddy`:

```bash
xcaddy build \
  --with github.com/soju841/caddy-dns-desec \
  --output ./caddy
