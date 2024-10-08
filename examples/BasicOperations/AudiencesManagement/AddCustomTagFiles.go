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
	"os"

	"github.com/chrisreal/marketing-api-go-sdk/pkg/ads"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/api"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/config"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/errors"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/model"
)

type CustomTagFilesAddExample struct {
	TAds                  *ads.SDKClient
	AccessToken           string
	AccountId             int64
	UserIdType            string
	TagId                 int64
	File                  *os.File
	CustomTagFilesAddOpts *api.CustomTagFilesAddOpts
}

func (e *CustomTagFilesAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.UserIdType = "QQ"
	e.TagId = int64(0)
	file, err := os.Open("YOUR FILE PATH")
	if err != nil {
		e.File = file
	}
	e.CustomTagFilesAddOpts = &api.CustomTagFilesAddOpts{}
}

func (e *CustomTagFilesAddExample) RunExample() (model.CustomTagFilesAddResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.CustomTagFiles().Add(ctx, e.AccountId, e.UserIdType, e.TagId, e.File, e.CustomTagFilesAddOpts)
}

func main() {
	e := &CustomTagFilesAddExample{}
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
