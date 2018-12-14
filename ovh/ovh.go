// Package ovh adapts the lego OVH DNS
// provider for Caddy. Importing this package plugs it in.
package ovh

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/ovh"
)

func init() {
	caddytls.RegisterDNSProvider("ovh", NewDNSProvider)
}

// NewDNSProvider returns a new OVH DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(4): credentials[0] = API Endpoint
//         credentials[1] = Application Key
//         credentials[2] = Application Secret
//         credentials[3] = Consumer Key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return ovh.NewDNSProvider()
	case 4:
		config := ovh.NewDefaultConfig()
		config.APIEndpoint = credentials[0]
		config.ApplicationKey = credentials[1]
		config.ApplicationSecret = credentials[2]
		config.ConsumerKey = credentials[3]
		return ovh.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
