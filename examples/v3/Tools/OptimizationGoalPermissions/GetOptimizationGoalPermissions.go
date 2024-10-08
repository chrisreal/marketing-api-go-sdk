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

type OptimizationGoalPermissionsGetExample struct {
	TAds                               *ads.SDKClient
	AccessToken                        string
	AccountId                          int64
	SiteSet                            []string
	MarketingGoal                      string
	MarketingSubGoal                   string
	MarketingCarrierType               string
	MarketingTargetType                string
	OptimizationGoalPermissionsGetOpts *api.OptimizationGoalPermissionsGetOpts
}

func (e *OptimizationGoalPermissionsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.SiteSet = []string{}
	e.MarketingGoal = "marketingGoal_example"
	e.MarketingSubGoal = "marketingSubGoal_example"
	e.MarketingCarrierType = "marketingCarrierType_example"
	e.MarketingTargetType = "marketingTargetType_example"
	e.OptimizationGoalPermissionsGetOpts = &api.OptimizationGoalPermissionsGetOpts{}
}

func (e *OptimizationGoalPermissionsGetExample) RunExample() (model.OptimizationGoalPermissionsGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.OptimizationGoalPermissions().Get(ctx, e.AccountId, e.SiteSet, e.MarketingGoal, e.MarketingSubGoal, e.MarketingCarrierType, e.MarketingTargetType, e.OptimizationGoalPermissionsGetOpts)
}

func main() {
	e := &OptimizationGoalPermissionsGetExample{}
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
