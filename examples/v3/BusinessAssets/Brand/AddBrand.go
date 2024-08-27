/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/chrisreal/marketing-api-go-sdk/pkg/ads/v3"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/config/v3"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/errors"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/model/v3"
)

type BrandAddExample struct {
	TAds           *ads.SDKClient
	AccessToken    string
	AccountId      int64
	Name           string
	BrandImageFile *os.File
}

func (e *BrandAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.Name = "YOUR BRAND NAME"
	file, err := os.Open("YOUR FILE PATH")
	if err != nil {
		e.BrandImageFile = file
	}
}

func (e *BrandAddExample) RunExample() (model.BrandAddResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Brand().Add(ctx, e.AccountId, e.Name, e.BrandImageFile)
}

func main() {
	e := &BrandAddExample{}
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
