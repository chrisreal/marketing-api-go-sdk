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

type HourlyReportsGetExample struct {
	TAds                 *ads.SDKClient
	AccessToken          string
	AccountId            int64
	Level                string
	DateRange            model.DateRange
	HourlyReportsGetOpts *api.HourlyReportsGetOpts
}

func (e *HourlyReportsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.Level = "REPORT_LEVEL_ADGROUP"
	e.DateRange = model.DateRange{
		StartDate: "REPORT DATE",
		EndDate:   "REPORT DATE",
	}
	e.HourlyReportsGetOpts = &api.HourlyReportsGetOpts{

		Fields: optional.NewInterface([]string{"hour", "view_count", "valid_click_count", "ctr", "cpc", "cost"}),
	}
}

func (e *HourlyReportsGetExample) RunExample() (model.HourlyReportsGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.HourlyReports().Get(ctx, e.AccountId, e.Level, e.DateRange, e.HourlyReportsGetOpts)
}

func main() {
	e := &HourlyReportsGetExample{}
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
