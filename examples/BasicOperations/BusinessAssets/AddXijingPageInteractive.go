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
)

type XijingPageInteractiveAddExample struct {
	TAds                         *ads.SDKClient
	AccessToken                  string
	AccountId                    int64
	IsAutoSubmit                 int64
	PageType                     string
	InteractivePageType          string
	PageTitle                    string
	PageName                     string
	MobileAppId                  string
	XijingPageInteractiveAddOpts *api.XijingPageInteractiveAddOpts
}

func (e *XijingPageInteractiveAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.IsAutoSubmit = 789
	e.PageType = "pageType_example"
	e.InteractivePageType = "interactivePageType_example"
	e.PageTitle = "pageTitle_example"
	e.PageName = "pageName_example"
	e.MobileAppId = "mobileAppId_example"
	e.XijingPageInteractiveAddOpts = &api.XijingPageInteractiveAddOpts{}
}

func (e *XijingPageInteractiveAddExample) RunExample() (interface{}, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.XijingPageInteractive().Add(ctx, e.AccountId, e.IsAutoSubmit, e.PageType, e.InteractivePageType, e.PageTitle, e.PageName, e.MobileAppId, e.XijingPageInteractiveAddOpts)
}

func main() {
	e := &XijingPageInteractiveAddExample{}
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
