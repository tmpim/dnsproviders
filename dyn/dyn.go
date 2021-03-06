// Package dyn adapts the lego Dyn DNS provider
// for Casket. Importing this package plugs it in.
package dyn

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/dyn"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("dyn", NewDNSProvider)
}

// NewDNSProvider returns a new Dyn DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(3): credentials[0] = customer name
//         credentials[1] = username
//         credentials[2] = password
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return dyn.NewDNSProvider()
	case 3:
		config := dyn.NewDefaultConfig()
		config.CustomerName = credentials[0]
		config.UserName = credentials[1]
		config.Password = credentials[2]
		return dyn.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
