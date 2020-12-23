// Package linode adapts the lego Linode DNS
// provider for Casket. Importing this package plugs it in.
package linode

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/linode"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("linode", NewDNSProvider)
}

// NewDNSProvider returns a new Linode DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = access token (API key)
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return linode.NewDNSProvider()
	case 1:
		config := linode.NewDefaultConfig()
		config.Token = credentials[0]
		return linode.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
