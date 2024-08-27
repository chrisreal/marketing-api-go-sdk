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

type UserActionSetsGetExample struct {
	TAds                  *ads.SDKClient
	AccessToken           string
	AccountId             int64
	UserActionSetsGetOpts *api.UserActionSetsGetOpts
}

func (e *UserActionSetsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.UserActionSetsGetOpts = &api.UserActionSetsGetOpts{

		Fields: optional.NewInterface([]string{"user_action_set_id", "type", "mobile_app_id", "name", "description", "activate_status", "created_time"}),
	}
}

func (e *UserActionSetsGetExample) RunExample() (model.UserActionSetsGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.UserActionSets().Get(ctx, e.AccountId, e.UserActionSetsGetOpts)
}

func main() {
	e := &UserActionSetsGetExample{}
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
