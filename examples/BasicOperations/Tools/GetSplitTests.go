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

type SplitTestsGetExample struct {
	TAds              *ads.SDKClient
	AccessToken       string
	AccountId         int64
	SplitTestsGetOpts *api.SplitTestsGetOpts
}

func (e *SplitTestsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.SplitTestsGetOpts = &api.SplitTestsGetOpts{

		Fields: optional.NewInterface([]string{"account_id", "split_test_id", "split_test_status", "split_test_name", "begin_time", "end_time", "smart_expand_enabled", "adgroup_id_list", "recommended_rating", "recommended_adgroup_id_list"}),
	}
}

func (e *SplitTestsGetExample) RunExample() (model.SplitTestsGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.SplitTests().Get(ctx, e.AccountId, e.SplitTestsGetOpts)
}

func main() {
	e := &SplitTestsGetExample{}
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
