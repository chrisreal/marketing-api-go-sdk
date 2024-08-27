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

type AdcreativeTemplatePreviewGetExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.AdcreativeTemplatePreviewGetRequest
}

func (e *AdcreativeTemplatePreviewGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.AdcreativeTemplatePreviewGetRequest{
		AccountId: int64(0),
		PreviewSpec: &model.AdcreativePreviewSpec{
			AdcreativeTemplateId: int64(133),
			SiteSet:              &[]string{"SITE_SET_WECHAT"},
			PromotedObjectType:   model.PromotedObjectType_LINK_WECHAT,
			AdcreativeElements: &model.AdcreativeCreativeElements{
				Image: "YOUR AD IMAGE ID",
			},
			PageType: model.DestinationType_DEFAULT,
			PageSpec: &model.PreviewPageSpec{
				PageUrl: "YOUR AD PAGE URL",
			},
		},
	}
}

func (e *AdcreativeTemplatePreviewGetExample) RunExample() (model.AdcreativeTemplatePreviewGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.AdcreativeTemplatePreview().Get(ctx, e.Data)
}

func main() {
	e := &AdcreativeTemplatePreviewGetExample{}
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
