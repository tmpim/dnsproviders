// Package lightsail adapts the lego AWS Lightsail DNS
// provider for Casket. Importing this package plugs it in.
package lightsail

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/lightsail"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("lightsail", NewDNSProvider)
}

// NewDNSProvider returns a new AWS Lightsail DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return lightsail.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
