// Package dnsimple adapts the lego DNSimple DNS provider
// for Casket. Importing this package plugs it in.
package dnsimple

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/dnsimple"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("dnsimple", NewDNSProvider)
}

// NewDNSProvider returns a new DNSimple DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = email
//         credentials[1] = API key
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return dnsimple.NewDNSProvider()
	case 2:
		config := dnsimple.NewDefaultConfig()
		config.AccessToken = credentials[1]
		return dnsimple.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
