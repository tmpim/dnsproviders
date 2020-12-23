// Package azure adapts the lego azure DNS
// provider for Casket. Importing this package plugs it in.
package azure

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/azure"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("azure", NewDNSProvider)
}

// NewDNSProvider returns a new azure DNS challenge provider.
// The credentials are detected automatically; see underlying
// package docs for details:
// https://godoc.org/github.com/go-acme/lego/providers/dns/azure
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return azure.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
