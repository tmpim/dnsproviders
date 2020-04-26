// Package dnspod adapts the lego Dnspod DNS
// provider for Casket. Importing this package plugs it in.
package dnspod

import (
	"errors"

	"github.com/tmpim/casket/caskettls"
	"github.com/go-acme/lego/v3/providers/dns/dnspod"
)

func init() {
	caskettls.RegisterDNSProvider("dnspod", NewDNSProvider)
}

// NewDNSProvider returns a new dnspod DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = access token (API key)
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return dnspod.NewDNSProvider()
	case 1:
		config := dnspod.NewDefaultConfig()
		config.LoginToken = credentials[0]
		return dnspod.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
