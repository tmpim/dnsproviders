// Package auroradns adapts the lego AuroraDNS DNS
// provider for Casket. Importing this package plugs it in.
package auroradns

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/auroradns"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("auroradns", NewDNSProvider)
}

// NewDNSProvider returns a new AuroraDNS DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(3): credentials[0] = Base URL
//         credentials[1] = User ID
//         credentials[2] = Key
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return auroradns.NewDNSProvider()
	case 3:
		config := auroradns.NewDefaultConfig()
		config.BaseURL = credentials[0]
		config.UserID = credentials[1]
		config.Key = credentials[2]
		return auroradns.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
