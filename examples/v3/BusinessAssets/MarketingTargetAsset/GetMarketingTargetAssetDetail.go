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

	"github.com/chrisreal/marketing-api-go-sdk/pkg/ads/v3"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/api/v3"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/config/v3"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/errors"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/model/v3"
)

type MarketingTargetAssetDetailGetExample struct {
	TAds                              *ads.SDKClient
	AccessToken                       string
	MarketingAssetId                  int64
	MarketingTargetType               string
	MarketingTargetAssetDetailGetOpts *api.MarketingTargetAssetDetailGetOpts
}

func (e *MarketingTargetAssetDetailGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.MarketingAssetId = 789
	e.MarketingTargetType = "marketingTargetType_example"
	e.MarketingTargetAssetDetailGetOpts = &api.MarketingTargetAssetDetailGetOpts{}
}

func (e *MarketingTargetAssetDetailGetExample) RunExample() (model.MarketingTargetAssetDetailGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.MarketingTargetAssetDetail().Get(ctx, e.MarketingAssetId, e.MarketingTargetType, e.MarketingTargetAssetDetailGetOpts)
}

func main() {
	e := &MarketingTargetAssetDetailGetExample{}
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
