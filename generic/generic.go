// Package generic adapts the lego exec DNS
// provider for Casket. Importing this package plugs it in.
package generic

import (
	"errors"

	"github.com/go-acme/lego/v4/providers/dns/exec"
	"github.com/tmpim/casket/caskettls"
)

func init() {
	caskettls.RegisterDNSProvider("generic", NewDNSProvider)
}

// NewDNSProvider returns a new exec DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use program to run from environment variable EXEC_PATH
// len(1): credentials[0] = program to run
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return exec.NewDNSProvider()
	case 1:
		config := exec.NewDefaultConfig()
		config.Program = credentials[0]
		return exec.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
