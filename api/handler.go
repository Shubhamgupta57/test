package api

import (
	"net/http"

	"integrations/api/wrapper"
)

func (a *API) requestHandler(h func(c *wrapper.RequestContext, w http.ResponseWriter, r *http.Request)) http.Handler {
	return &wrapper.Request{
		HandlerFunc: h,
	}
}
