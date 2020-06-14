package sonos

import (
	"context"

	"golang.org/x/oauth2"
)

// API represents access to the sonos api
type API struct {
	OAuth2 *oauth2.Config
	ctx    context.Context
}

// NewAPI returns an API instance ready to handle requests
func NewAPI(ctx context.Context, o *oauth2.Config) *API {
	return &API{
		ctx:    ctx,
		OAuth2: o,
	}
}
