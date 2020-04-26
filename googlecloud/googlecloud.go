// Package googlecloud adapts the lego Google Cloud DNS
// provider for Casket. Importing this package plugs it in.
package googlecloud

import (
	"errors"

	"github.com/tmpim/casket/caskettls"
	"github.com/go-acme/lego/v3/providers/dns/gcloud"
)

func init() {
	caskettls.RegisterDNSProvider("googlecloud", NewDNSProvider)
}

// NewDNSProvider returns a new Google Cloud DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = project
func NewDNSProvider(credentials ...string) (caskettls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return gcloud.NewDNSProvider()
	case 1:
		config := gcloud.NewDefaultConfig()
		config.Project = credentials[0]
		return gcloud.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
