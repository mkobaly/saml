package samlsp

import (
	"errors"
	"net/http"

	"github.com/mkobaly/saml"
)

var ErrSPNotFound = errors.New("Service Provider not found")

// Resolver is used to identify the correct Service Provider settings
// for a given web request. This allows you to support multiple
// Identity Providers since the Idp settings are stored with the Serivce Provider
type Resolver interface {
	LookupSP(r *http.Request) (*saml.ServiceProvider, error)
}
