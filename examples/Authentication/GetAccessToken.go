package main

import (
	"encoding/json"
	"fmt"

	"github.com/antihax/optional"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/ads"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/api"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/config"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/errors"
)

// 获取token示例
type GetAccessTokenExample struct {
	TAds           *ads.SDKClient
	AccessToken    string
	ClientId       int64
	ClientSecret   string
	GrantType      string
	OauthTokenOpts *api.OauthTokenOpts
}

// Init ...
func (e *GetAccessTokenExample) Init() {
	e.TAds = ads.Init(&config.SDKConfig{
		IsDebug: true,
	})
	// oauth/token不提供沙箱环境
	e.TAds.UseProduction()
	// YOUR CLIENT ID
	e.ClientId = int64(0)
	e.ClientSecret = "YOUR CLIENT SECRET"
	e.GrantType = "authorization_code"
	e.OauthTokenOpts = &api.OauthTokenOpts{
		AuthorizationCode: optional.NewString("YOUR AUTHORIZATION CODE"),
		RedirectUri:       optional.NewString("YOUR REDIRECT URI"),
	}
}

// RunExample ...
func (e *GetAccessTokenExample) RunExample() {
	tads := e.TAds

	// change ctx as needed
	ctx := *tads.Ctx
	response, _, err := tads.Oauth().Token(ctx, e.ClientId, e.ClientSecret, e.GrantType, e.OauthTokenOpts)

	if err != nil {
		//
		if resErr, ok := err.(errors.ResponseError); ok {
			// When Api returns an error
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Resopnse error:", string(errStr))
		} else {
			// When validation fails or other local issues
			fmt.Println("Error:", err)
		}
	} else {
		// 从返回里获得AccessToken并设置到tads中
		tads.SetAccessToken(*response.AccessToken)
	}
}

func main() {
	e := &GetAccessTokenExample{}
	e.Init()
	e.RunExample()
}
