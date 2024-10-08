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

type DynamicAdVideoAddExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.DynamicAdVideoAddRequest
}

func (e *DynamicAdVideoAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.DynamicAdVideoAddRequest{
		ProductCatalogId:    int64(0),
		VideoMaxDuration:    int64(0),
		AccountId:           int64(0),
		ProductMode:         model.AdNum_SINGLE,
		ProductSource:       "YOUR PRODUCT ID",
		DynamicAdTemplateId: int64(0),
	}
}

func (e *DynamicAdVideoAddExample) RunExample() (model.DynamicAdVideoAddResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.DynamicAdVideo().Add(ctx, e.Data)
}

func main() {
	e := &DynamicAdVideoAddExample{}
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
