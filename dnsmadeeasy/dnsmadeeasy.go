// Package dnsimple adapts the lego DNS Made Easy DNS provider
// for Casket. Importing this package plugs it in.
package dnsmadeeasy

import (
	"errors"

	"github.com/tmpim/casket/caskettls"
	"github.com/go-acme/lego/v3/providers/dns/dnsmadeeasy"
)

func init() {
	caskettls.RegisterDNSProvider("dnsmadeeasy", NewDNSProvider)
}

// NewDNSProvider returns a new DNS Made Easy DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(3): credentials[0] = API Endpoint
//         credentials[1] = API key
//         credentials[2] = API secret
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return dnsmadeeasy.NewDNSProvider()
	case 3:
		config := dnsmadeeasy.NewDefaultConfig()
		config.BaseURL = credentials[0]
		config.APIKey = credentials[1]
		config.APISecret = credentials[2]
		return dnsmadeeasy.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
