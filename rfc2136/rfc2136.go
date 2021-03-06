// Package rfc2136 adapts the lego RFC 2136 dynamic update DNS
// provider for Casket. Importing this package plugs it in.
package rfc2136

import (
	"errors"
	"time"

	"github.com/go-acme/lego/v4/providers/dns/rfc2136"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("rfc2136", NewDNSProvider)
}

// NewDNSProvider returns a new RFC 2136 DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(4): credentials[0] = nameserver
//         credentials[1] = TSIG algorithm
//         credentials[2] = TSIG key
//         credentials[3] = TSIG secret
//         DNS propagation timeout uses default from github.com/go-acme/lego/providers/dns/rfc2136 (60s)
// len(5): credentials[0] = nameserver
//         credentials[1] = TSIG algorithm
//         credentials[2] = TSIG key
//         credentials[3] = TSIG secret
//         credentials[4] = DNS propagation timeout
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	var timeout time.Duration

	switch len(credentials) {
	case 0:
		return rfc2136.NewDNSProvider()
	case 5:
		var err error
		timeout, err = time.ParseDuration(credentials[4])
		if err != nil {
			return nil, errors.New("invalid DNS propagation timeout")
		}
		fallthrough
	case 4:
		config := rfc2136.NewDefaultConfig()
		config.Nameserver = credentials[0]
		config.TSIGAlgorithm = credentials[1]
		config.TSIGKey = credentials[2]
		config.TSIGSecret = credentials[3]
		config.PropagationTimeout = timeout
		return rfc2136.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
