// Package gandi adapts the lego Gandi DNS provider
// for Caddy. Importing this package plugs it in.
package gandi

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/gandi"
)

func init() {
	caddytls.RegisterDNSProvider("gandi", NewDNSProvider)
}

// NewDNSProvider returns a new Gandi DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = API key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return gandi.NewDNSProvider()
	case 1:
		config := gandi.NewDefaultConfig()
		config.APIKey = credentials[0]
		return gandi.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
