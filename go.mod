module github.com/uptimerobot/certmagic

go 1.24.0

require (
	github.com/klauspost/cpuid v1.2.5
	github.com/libdns/libdns v0.2.0
	github.com/mholt/acmez v0.1.3
	github.com/miekg/dns v1.1.30
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.45.0
	golang.org/x/net v0.47.0
)

require (
	go.uber.org/atomic v1.6.0 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
)

replace github.com/caddyserver/certmagic => github.com/uptimerobot/certmagic v0.12.0
