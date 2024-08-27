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

type SplitTestsUpdateExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.SplitTestsUpdateRequest
}

func (e *SplitTestsUpdateExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.SplitTestsUpdateRequest{
		AccountId:       int64(0),
		SplitTestStatus: model.SplitTestStatus_SUSPEND,
		SplitTestId:     int64(0),
	}
}

func (e *SplitTestsUpdateExample) RunExample() (model.SplitTestsUpdateResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.SplitTests().Update(ctx, e.Data)
}

func main() {
	e := &SplitTestsUpdateExample{}
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
