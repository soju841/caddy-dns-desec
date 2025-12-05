module github.com/soju841/caddy-dns-desec

go 1.21

require (
    github.com/caddyserver/caddy/v2 v2.10.2
    github.com/libdns/desec v0.0.0-20240422000000-3cac5c97ba68
)

replace github.com/libdns/desec => github.com/libdns/desec v0.0.0-20240422000000-3cac5c97ba68
