// Package httpreq adapts the lego httpreq DNS
// provider for Casket. Importing this package plugs it in.
package httpreq

import (
	"errors"
	"net/url"

	"github.com/go-acme/lego/v4/providers/dns/httpreq"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("httpreq", NewDNSProvider)
}

// NewDNSProvider returns a new httpreq DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = Endpoint
//         credentials[1] = Mode
//         credentials[2] = Username
//         credentials[3] = Password
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return httpreq.NewDNSProvider()
	case 4:
		config := httpreq.NewDefaultConfig()

		endpoint, err := url.Parse(credentials[0])
		if err != nil {
			return nil, errors.New("endpoint is not a valid URL")
		}

		config.Endpoint = endpoint
		config.Mode = credentials[1]
		config.Username = credentials[2]
		config.Password = credentials[3]
		return httpreq.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
