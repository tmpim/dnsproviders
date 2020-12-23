// Package route53 adapts the lego Route53 DNS
// provider for Casket. Importing this package plugs it in.
package route53

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/route53"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("route53", NewDNSProvider)
}

// NewDNSProvider returns a new Route53 DNS challenge provider.
// The credentials are detected automatically; see underlying
// package docs for details:
// https://godoc.org/github.com/go-acme/lego/providers/dns/route53
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return route53.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
