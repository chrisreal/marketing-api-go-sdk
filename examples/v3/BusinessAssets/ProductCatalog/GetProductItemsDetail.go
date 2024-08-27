/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/ads/v3"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/api/v3"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/config/v3"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/errors"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/model/v3"
)

type ProductItemsDetailGetExample struct {
	TAds                      *ads.SDKClient
	AccessToken               string
	AccountId                 int64
	ProductCatalogId          int64
	ProductOuterId            string
	ProductItemsDetailGetOpts *api.ProductItemsDetailGetOpts
}

func (e *ProductItemsDetailGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.ProductCatalogId = int64(0)
	e.ProductOuterId = "YOUR PRODUCT ID"
	e.ProductItemsDetailGetOpts = &api.ProductItemsDetailGetOpts{

		Fields: optional.NewInterface([]string{"feed_id", "system_status", "reject_message", "product_item_spec"}),
	}
}

func (e *ProductItemsDetailGetExample) RunExample() (model.ProductItemsDetailGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.ProductItemsDetail().Get(ctx, e.AccountId, e.ProductCatalogId, e.ProductOuterId, e.ProductItemsDetailGetOpts)
}

func main() {
	e := &ProductItemsDetailGetExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}
