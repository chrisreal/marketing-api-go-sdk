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
	"github.com/chrisreal/marketing-api-go-sdk/pkg/config"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/errors"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/model"
)

type AdcreativesUpdateExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.AdcreativesUpdateRequest
}

func (e *AdcreativesUpdateExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.AdcreativesUpdateRequest{
		AccountId: int64(0),
		AdcreativeElements: &model.AdcreativeCreativeElementsMp{
			Title:       "YOUR AD TEXT",
			Description: "YOUR AD DESCRIPTION",
		},
		PageSpec: &model.PageSpec{
			PageUrl: "YOUR AD PAGE URL",
		},
		AdcreativeId: int64(0),
	}
}

func (e *AdcreativesUpdateExample) RunExample() (model.AdcreativesUpdateResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Adcreatives().Update(ctx, e.Data)
}

func main() {
	e := &AdcreativesUpdateExample{}
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
