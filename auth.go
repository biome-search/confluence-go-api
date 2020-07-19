package goconfluence

import (
	"net/http"
)

// Auth implements basic auth
func (a *API) Auth(req *http.Request) {
	tok, _ := a.tokenSource.Token()

	tok.SetAuthHeader(req)
}
