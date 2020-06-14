package sonos

import (
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

// SubscribeHousehold requests event subscription for household with supplied ID
func (api *API) SubscribeHousehold(id string, tok *oauth2.Token) (*Result, error) {

	log.Printf("subscribeHousehold called")

	r, err := api.apiCall(http.MethodPost, "households/"+id+"/groups/subscription", nil, *tok)
	if err != nil {
		return r, err
	}
	log.Printf("subscribeHousehold Returning")
	return r, nil
}
