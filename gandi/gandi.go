// Package gandi adapts the lego Gandi DNS provider
// for Casket. Importing this package plugs it in.
package gandi

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/gandi"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("gandi", NewDNSProvider)
}

// NewDNSProvider returns a new Gandi DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = API key
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return gandi.NewDNSProvider()
	case 1:
		config := gandi.NewDefaultConfig()
		config.APIKey = credentials[0]
		return gandi.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
