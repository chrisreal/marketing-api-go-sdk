/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/ads"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/api"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/config"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/errors"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/model"
)

type ProductItemsGetExample struct {
	TAds                *ads.SDKClient
	AccessToken         string
	AccountId           int64
	ProductCatalogId    int64
	ProductItemsGetOpts *api.ProductItemsGetOpts
}

func (e *ProductItemsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.ProductCatalogId = int64(0)
	e.ProductItemsGetOpts = &api.ProductItemsGetOpts{

		Fields: optional.NewInterface([]string{"product_outer_id", "product_name", "product_image_url"}),
	}
}

func (e *ProductItemsGetExample) RunExample() (model.ProductItemsGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.ProductItems().Get(ctx, e.AccountId, e.ProductCatalogId, e.ProductItemsGetOpts)
}

func main() {
	e := &ProductItemsGetExample{}
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
