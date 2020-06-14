package sonos

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
)

// Result contains the data from the API http response in case it's required
type Result struct {
	Status        int           // HTTP Status Code
	Body          []byte        // HTTP Response Body
	Headers       http.Header   // HTTP Response Headers
	RefeshedToken *oauth2.Token // None nil value if token was refreshed
}

func (api *API) apiCall(method string, path string, body []byte, tok oauth2.Token) (*Result, error) {
	log.Printf("apiCall called")

	var err error // Local Error Variable
	var result = &Result{}

	// Construct new API call URL
	u := url.URL{
		Scheme: "https",
		Path:   "api.ws.sonos.com/control/api/v1/" + path,
	}

	// Refesh Token if required
	if tok.Expiry.After(time.Now().Add(time.Minute * -15)) {
		// Get new token
		result.RefeshedToken, err = api.OAuth2.TokenSource(api.ctx, &tok).Token()
		if err != nil {
			log.Printf("Error refreshing oauth2 token")
			return nil, err
		}
		tok = *result.RefeshedToken // Set old token to new one for this call
	}

	// Create a new oauth2 client
	client := api.OAuth2.Client(api.ctx, &tok)
	// Do the request
	var resp *http.Response
	switch method {
	case http.MethodGet:
		resp, err = client.Get(u.String())
	case http.MethodPost:
		resp, err = client.Post(u.String(), "application/json", bytes.NewReader(body))
	}

	// Save values from response if present
	if resp != nil {
		result.Status = resp.StatusCode
		result.Headers = resp.Header
	}
	if err != nil {
		return result, err
	}

	// Get the response body as a byte array for ease
	result.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// Return results
	return result, nil
}
