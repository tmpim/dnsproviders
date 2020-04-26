// Package alidns adapts the lego Ali DNS
// provider for Casket. Importing this package plugs it in.
package alidns

import (
	"errors"

	"github.com/tmpim/casket/caskettls"
	"github.com/go-acme/lego/v3/providers/dns/alidns"
)

func init() {
	caskettls.RegisterDNSProvider("alidns", NewDNSProvider)
}

// NewDNSProvider returns a new alidns DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = access token (API key)
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return alidns.NewDNSProvider()
	case 1:
		config := alidns.NewDefaultConfig()
		config.APIKey = credentials[0]
		config.SecretKey = credentials[1]
		return alidns.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
