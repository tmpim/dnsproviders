// Package duckdns adapts the lego duckdns DNS
// provider for Casket. Importing this package plugs it in.
package duckdns

import (
	"errors"

	"github.com/tmpim/casket/caskettls"
	"github.com/go-acme/lego/v3/providers/dns/duckdns"
)

func init() {
	caskettls.RegisterDNSProvider("duckdns", NewDNSProvider)
}

// NewDNSProvider returns a new duckdns DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = duckdns token
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return duckdns.NewDNSProvider()
	case 1:
		config := duckdns.NewDefaultConfig()
		config.Token = credentials[0]
		return duckdns.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
