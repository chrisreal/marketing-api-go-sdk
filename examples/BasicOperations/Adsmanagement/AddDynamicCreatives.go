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

type DynamicCreativesAddExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.DynamicCreativesAddRequest
}

func (e *DynamicCreativesAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.DynamicCreativesAddRequest{
		AccountId: int64(0),
		DynamicCreativeElements: &model.DynamicCreativeElements{
			ImageOptions:       &[]string{"YOUR AD IMAGE"},
			TitleOptions:       &[]string{"YOUR AD TEXT 1", "YOUR AD TEXT 2"},
			DescriptionOptions: &[]string{"YOUR AD DESCRIPTION"},
		},
		PageType: model.DestinationType_DEFAULT,
		PageSpec: &model.DynamicCreativePageSpec{
			PageUrl: "YOUR AD PAGE URL",
		},
		CampaignType:              model.CampaignType_NORMAL,
		PromotedObjectType:        model.PromotedObjectType_LINK,
		SiteSet:                   &[]string{"SITE_SET_MOBILE_INNER"},
		DynamicCreativeTemplateId: int64(0),
		DynamicCreativeName:       "SDK动态创意5ede25293ba60",
	}
}

func (e *DynamicCreativesAddExample) RunExample() (model.DynamicCreativesAddResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.DynamicCreatives().Add(ctx, e.Data)
}

func main() {
	e := &DynamicCreativesAddExample{}
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
