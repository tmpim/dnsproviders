package stackpath

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/stackpath"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("stackpath", NewDNSProvider)
}

// NewDNSProvider returns a new Stackpath DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(3): credentials[0] = client id
//         credentials[1] = client secret
//         credentials[2] = stack id
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return stackpath.NewDNSProvider()
	case 3:
		config := stackpath.NewDefaultConfig()
		config.ClientID = credentials[0]
		config.ClientSecret = credentials[1]
		config.StackID = credentials[2]
		return stackpath.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
