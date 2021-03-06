// Package inwx adapts the lego INWX DNS
// provider for Casket. Importing this package plugs it in.
package inwx

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/inwx"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("inwx", NewDNSProvider)
}

// NewDNSProvider returns a new INWX DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = Username
//         credentials[1] = Password
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return inwx.NewDNSProvider()
	case 2:
		config := inwx.NewDefaultConfig()
		config.Username = credentials[0]
		config.Password = credentials[1]
		return inwx.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
