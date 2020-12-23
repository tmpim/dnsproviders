// Package acmedns adapts the lego exec DNS
// provider for Casket. Importing this package plugs it in.
package acmedns

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/acmedns"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("acmedns", NewDNSProvider)
}

// NewDNSProvider returns a new acmedns DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return acmedns.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
