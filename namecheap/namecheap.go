// Package namecheap adapts the lego NameCheap DNS provider
// for Casket. Importing this package plugs it in.
package namecheap

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/namecheap"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("namecheap", NewDNSProvider)
}

// NewDNSProvider returns a new NameCheap DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = API user
//         credentials[1] = API key
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return namecheap.NewDNSProvider()
	case 2:
		config := namecheap.NewDefaultConfig()
		config.APIUser = credentials[0]
		config.APIKey = credentials[1]
		return namecheap.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
