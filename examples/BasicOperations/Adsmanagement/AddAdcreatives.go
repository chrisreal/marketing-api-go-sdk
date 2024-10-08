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

type AdcreativesAddExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.AdcreativesAddRequest
}

func (e *AdcreativesAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.AdcreativesAddRequest{
		AccountId:            int64(0),
		AdcreativeTemplateId: int64(0),
		AdcreativeElements: &model.AdcreativeCreativeElementsMp{
			Title:       "YOUR AD TEXT",
			Description: "YOUR AD DESCRIPTION",
		},
		PageType: model.DestinationType_DEFAULT,
		PageSpec: &model.PageSpec{
			PageUrl: "YOUR AD PAGE URL",
		},
		PromotedObjectType: model.PromotedObjectType_LINK,
		AdcreativeName:     "SDK广告创意5ede2529029e9",
		SiteSet:            &[]string{"SITE_SET_QZONE"},
		CampaignId:         int64(0),
	}
}

func (e *AdcreativesAddExample) RunExample() (model.AdcreativesAddResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Adcreatives().Add(ctx, e.Data)
}

func main() {
	e := &AdcreativesAddExample{}
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
