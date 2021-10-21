package app

import (
	"context"
	"encoding/json"
	"integrations/model"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (a *App) GetShopifyOrders(cID primitive.ObjectID) (map[string]interface{}, error) {
	var shopify model.Shopify
	opts := options.FindOne().SetProjection(bson.M{"_id": 0, "shopify": 1})
	err := a.MongoDB.Database.Collection(model.ClientColl).FindOne(context.TODO(), bson.M{"_id": cID}, opts).Decode(&shopify)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to query database")
	}

	url := "https://" + shopify.GetShopifyDetails.Shop + "/admin/shop.json"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-Shopify-Access-Token", shopify.GetShopifyDetails.AccessToken)

	client := &http.Client{}

	resp, err := client.Do(req) //making request to shopify for access token
	if err != nil {
		return nil, errors.Wrapf(err, "Not able to make request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't read response body")
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result) //translating json response into a map

	return result, nil
}
