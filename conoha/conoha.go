// Package conoha adapts the lego ConoHa DNS provider
// for Casket. Importing this package plugs it in.
package conoha

import (
	"errors"

	"github.com/tmpim/casket/caskettls"
	"github.com/go-acme/lego/v3/providers/dns/conoha"
)

func init() {
	caskettls.RegisterDNSProvider("conoha", NewDNSProvider)
}

// NewDNSProvider returns a new ConoHa DNS challenge provider.
// The credentials are detected automatically; see underlying
// package docs for details:
// https://godoc.org/github.com/go-acme/lego/providers/dns/conoha
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return conoha.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
