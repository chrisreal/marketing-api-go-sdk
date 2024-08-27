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

type UserPropertySetsGetExample struct {
	TAds                    *ads.SDKClient
	AccessToken             string
	AccountId               int64
	UserPropertySetsGetOpts *api.UserPropertySetsGetOpts
}

func (e *UserPropertySetsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.UserPropertySetsGetOpts = &api.UserPropertySetsGetOpts{

		Fields: optional.NewInterface([]string{"user_property_set_id", "name", "description", "created_time"}),
	}
}

func (e *UserPropertySetsGetExample) RunExample() (model.UserPropertySetsGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.UserPropertySets().Get(ctx, e.AccountId, e.UserPropertySetsGetOpts)
}

func main() {
	e := &UserPropertySetsGetExample{}
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
