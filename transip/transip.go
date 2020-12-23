// Package transip adapts the lego TransIP DNS
// provider for Casket. Importing this package plugs it in.
package transip

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/transip"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("transip", NewDNSProvider)
}

// NewDNSProvider returns a new TransIP DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = Account Name
//         credentials[1] = Private Key Path
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return transip.NewDNSProvider()
	case 2:
		config := transip.NewDefaultConfig()
		config.AccountName = credentials[0]
		config.PrivateKeyPath = credentials[1]
		return transip.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
