// Package vscale adapts the lego Vscale DNS
// provider for Casket. Importing this package plugs it in.
package vscale

import (
	"errors"

	"github.com/tmpim/casket/caskettls"
	"github.com/go-acme/lego/v3/providers/dns/vscale"
)

func init() {
	caskettls.RegisterDNSProvider("vscale", NewDNSProvider)
}

// NewDNSProvider returns a new Vscale DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment (https://godoc.org/github.com/go-acme/lego/providers/dns/vscale)
// len(1): credentials[0] = Token
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return vscale.NewDNSProvider()
	case 1:
		config := vscale.NewDefaultConfig()
		config.Token = credentials[0]
		return vscale.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
