package api

import (
	"integrations/api/wrapper"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *API) getShopifyOrders(requestCTX *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("client_id")
	ClientId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		requestCTX.SetErr(err, 400)
		return
	}

	res, err := a.App.GetShopifyOrders(ClientId)
	if err != nil {
		requestCTX.SetErr(err, 400)
		return
	}

	requestCTX.SetAppResponse(res, 200)
}
