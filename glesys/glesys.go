// Package glesys adapts the lego GleSYS DNS provider
// for Casket. Importing this package plugs it in.
package glesys

import (
	"errors"

	"github.com/tmpim/casket/caskettls"
	"github.com/go-acme/lego/v3/providers/dns/glesys"
)

func init() {
	caskettls.RegisterDNSProvider("glesys", NewDNSProvider)
}

// NewDNSProvider returns a new GleSYS DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = API user
//         credentials[1] = API key
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return glesys.NewDNSProvider()
	case 2:
		config := glesys.NewDefaultConfig()
		config.APIUser = credentials[0]
		config.APIKey = credentials[1]
		return glesys.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
