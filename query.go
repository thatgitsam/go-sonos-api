package sonos

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

// GetHouseholds returns all Households linked to the supplied Token
func (api *API) GetHouseholds(tok *oauth2.Token) ([]Household, *Result, error) {

	// Local structure for JSON dump
	type account struct {
		Household []Household `json:"households,omitempty"`
	}
	var a *account

	r, err := api.apiCall(http.MethodGet, "households", nil, *tok)
	if err != nil {
		return nil, r, err
	}
	err = json.Unmarshal(r.Body, &a)
	if err != nil {
		return nil, r, err
	}

	for i, h := range a.Household {
		a.Household[i].Name = "Household-" + h.ID[6:10] // Set a friendly name based on ID start

		// Get list of players and drop into description
		x, tok, err := api.GetHousehold(h.ID, tok)
		if err != nil {
			return nil, tok, err
		}
		// Store details
		a.Household[i].Groups = x.Groups
		a.Household[i].Players = x.Players
		// Generate Description
		for _, p := range a.Household[i].Players {
			a.Household[i].Desc = a.Household[i].Desc + "[" + p.Name + "]"
		}
	}

	return a.Household, r, nil
}

// GetHousehold returns Households by ID linked to the supplied Token
func (api *API) GetHousehold(id string, tok *oauth2.Token) (*Household, *Result, error) {

	var h *Household

	r, err := api.apiCall(http.MethodGet, "households/"+id+"/groups", nil, *tok)
	if err != nil {
		return nil, r, err
	}
	err = json.Unmarshal(r.Body, &h)
	if err != nil {
		log.Println(err)
		return nil, r, err
	}

	return h, r, nil
}
