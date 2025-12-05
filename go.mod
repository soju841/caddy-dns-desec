module github.com/soju841/caddy-dns-desec

go 1.21

require (
    github.com/caddyserver/caddy/v2 v2.10.2
    github.com/libdns/desec v0.0.0-00010101000000-000000000000
)

replace github.com/libdns/desec => github.com/libdns/desec latest
