package model

// Shopify
type Shopify struct {
	GetShopifyDetails GetShopifyDetails `json:"shopify,omitempty" bson:"shopify,omitempty"`
}

// GetShopifyDetails contains schema for shopify data in client collection
type GetShopifyDetails struct {
	Shop        string `json:"name,omitempty" bson:"name,omitempty"`
	AccessToken string `json:"access_token,omitempty" bson:"access_token,omitempty"`
	Scope       string `json:"scope,omitempty" bson:"scope,omitempty"`
}
