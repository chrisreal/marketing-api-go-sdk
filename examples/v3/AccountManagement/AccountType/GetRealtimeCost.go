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

type RealtimeCostGetExample struct {
	TAds                *ads.SDKClient
	AccessToken         string
	AccountId           int64
	Level               string
	Date                string
	RealtimeCostGetOpts *api.RealtimeCostGetOpts
}

func (e *RealtimeCostGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.Level = "ADVERTISER"
	e.Date = "YOUR DATE"
	e.RealtimeCostGetOpts = &api.RealtimeCostGetOpts{

		Fields: optional.NewInterface([]string{"campaign_id", "adgroup_id", "cost"}),
	}
}

func (e *RealtimeCostGetExample) RunExample() (model.RealtimeCostGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.RealtimeCost().Get(ctx, e.AccountId, e.Level, e.Date, e.RealtimeCostGetOpts)
}

func main() {
	e := &RealtimeCostGetExample{}
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
