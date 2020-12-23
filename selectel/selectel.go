// Package selectel adapts the lego Selectel DNS
// provider for Casket. Importing this package plugs it in.
package selectel

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/selectel"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("selectel", NewDNSProvider)
}

// NewDNSProvider returns a new Selectel DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment (https://godoc.org/github.com/go-acme/lego/providers/dns/selectel)
// len(1): credentials[0] = Token
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return selectel.NewDNSProvider()
	case 1:
		config := selectel.NewDefaultConfig()
		config.Token = credentials[0]
		return selectel.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
