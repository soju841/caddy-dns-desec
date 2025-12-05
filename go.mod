module github.com/soju841/caddy-dns-desec

go 1.21

require (
	github.com/caddyserver/caddy/v2 v2.10.2
	github.com/libdns/desec v0.2.1
)

require github.com/libdns/libdns v0.2.1 // indirect

replace github.com/caddyserver/caddy/v2 => github.com/caddyserver/caddy/v2 v2.10.2
