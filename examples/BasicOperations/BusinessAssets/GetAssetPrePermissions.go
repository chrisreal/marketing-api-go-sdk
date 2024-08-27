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

	"github.com/chrisreal/marketing-api-go-sdk/pkg/ads"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/api"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/config"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/errors"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/model"
)

type AssetPrePermissionsGetExample struct {
	TAds                       *ads.SDKClient
	AccessToken                string
	AccountId                  int64
	AssetType                  string
	AssetPrePermissionsGetOpts *api.AssetPrePermissionsGetOpts
}

func (e *AssetPrePermissionsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.AssetType = "assetType_example"
	e.AssetPrePermissionsGetOpts = &api.AssetPrePermissionsGetOpts{}
}

func (e *AssetPrePermissionsGetExample) RunExample() (model.AssetPrePermissionsGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.AssetPrePermissions().Get(ctx, e.AccountId, e.AssetType, e.AssetPrePermissionsGetOpts)
}

func main() {
	e := &AssetPrePermissionsGetExample{}
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
