// Package vultr adapts the lego Vultr DNS provider
// for Casket. Importing this package plugs it in.
package vultr

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/vultr"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("vultr", NewDNSProvider)
}

// NewDNSProvider returns a new Vultr DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = API key
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return vultr.NewDNSProvider()
	case 1:
		config := vultr.NewDefaultConfig()
		config.APIKey = credentials[0]
		return vultr.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
