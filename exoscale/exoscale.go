// Package exoscale adapts the lego Exoscale DNS provider
// for Casket. Importing this package plugs it in.
package exoscale

import (
	"errors"

	"github.com/tmpim/casket/caskettls"
	"github.com/go-acme/lego/v3/providers/dns/exoscale"
)

func init() {
	caskettls.RegisterDNSProvider("exoscale", NewDNSProvider)
}

// NewDNSProvider returns a new Exoscale DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = API Key
//         credentials[1] = API Secret
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return exoscale.NewDNSProvider()
	case 2:
		config := exoscale.NewDefaultConfig()
		config.APIKey = credentials[0]
		config.APISecret = credentials[1]
		return exoscale.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
