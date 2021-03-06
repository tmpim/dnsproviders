// Package gandiv5 adapts the lego Gandiv5 DNS provider
// for Casket. Importing this package plugs it in.
package gandiv5

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/gandiv5"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("gandiv5", NewDNSProvider)
}

// NewDNSProvider returns a new Gandiv5 DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use API key from `GANDIV5_API_KEY` env var
// len(1): credentials[0] = API key
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return gandiv5.NewDNSProvider()
	case 1:
		config := gandiv5.NewDefaultConfig()
		config.APIKey = credentials[0]
		return gandiv5.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
