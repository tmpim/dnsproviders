// Package linodev4 adapts the lego LinodeV4 DNS
// provider for Casket. Importing this package plugs it in.
package linodev4

import (
	"errors"

	"github.com/tmpim/casket/caskettls"
	"github.com/go-acme/lego/v3/providers/dns/linodev4"
)

func init() {
	caskettls.RegisterDNSProvider("linodev4", NewDNSProvider)
}

// NewDNSProvider returns a new LinodeV4 DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = access token (API token)
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return linodev4.NewDNSProvider()
	case 1:
		config := linodev4.NewDefaultConfig()
		config.Token = credentials[0]
		return linodev4.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
