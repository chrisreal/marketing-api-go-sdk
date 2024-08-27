/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/antihax/optional"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/errors"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/model"
	. "github.com/chrisreal/marketing-api-go-sdk/pkg/model"
)

// Linger please
var (
	_ context.Context
)

type VideomakerVideocapturesApiService service

/*
VideomakerVideocapturesApiService 生成视频封面图
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param optional nil or *VideomakerVideocapturesAddOpts - Optional Parameters:
     * @param "VideoId" (optional.String) -
     * @param "VideoFile" (optional.Interface of *os.File) -
     * @param "Signature" (optional.String) -
     * @param "Number" (optional.Int64) -
     * @param "ReturnImageIds" (optional.Bool) -

@return VideomakerVideocapturesAddResponse
*/

type VideomakerVideocapturesAddOpts struct {
	VideoId        optional.String
	VideoFile      optional.Interface
	Signature      optional.String
	Number         optional.Int64
	ReturnImageIds optional.Bool
}

func (a *VideomakerVideocapturesApiService) Add(ctx context.Context, accountId int64, localVarOptionals *VideomakerVideocapturesAddOpts) (VideomakerVideocapturesAddResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue VideomakerVideocapturesAddResponseData
		localVarResponse    VideomakerVideocapturesAddResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/videomaker_videocaptures/add"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFormParams.Add("account_id", parameterToString(accountId, ""))
	if localVarOptionals != nil && localVarOptionals.VideoId.IsSet() {
		localVarFormParams.Add("video_id", parameterToString(localVarOptionals.VideoId.Value(), ""))
	}
	var localVarFile *os.File
	localVarFileKey = "video_file"
	if localVarOptionals != nil && localVarOptionals.VideoFile.IsSet() {
		localVarFileOk := false
		localVarFile, localVarFileOk = localVarOptionals.VideoFile.Value().(*os.File)
		if !localVarFileOk {
			return localVarReturnValue, nil, reportError("videoFile should be *os.File")
		}
	}
	if localVarFile != nil {
		localVarNewFile, localErr := os.Open(localVarFile.Name())
		if localErr != nil {
			return localVarReturnValue, nil, localErr
		}
		defer localVarNewFile.Close()
		fbs, localErr := ioutil.ReadAll(localVarNewFile)
		if localErr != nil {
			return localVarReturnValue, nil, localErr
		}
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
	}
	if localVarOptionals != nil && localVarOptionals.Signature.IsSet() {
		localVarFormParams.Add("signature", parameterToString(localVarOptionals.Signature.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Number.IsSet() {
		localVarFormParams.Add("number", parameterToString(localVarOptionals.Number.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReturnImageIds.IsSet() {
		localVarFormParams.Add("return_image_ids", parameterToString(localVarOptionals.ReturnImageIds.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, localVarFileKey)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, nil, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarResponse, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			if *localVarResponse.Code != 0 {
				var localVarResponseErrors []model.ApiErrorStruct
				if localVarResponse.Errors != nil {
					localVarResponseErrors = *localVarResponse.Errors
				}
				err = errors.NewError(localVarResponse.Code, localVarResponse.Message, localVarResponse.MessageCn, localVarResponseErrors)
				return localVarReturnValue, localVarHttpResponse.Header, err
			}
			if localVarResponse.Data == nil {
				return localVarReturnValue, localVarHttpResponse.Header, err
			} else {
				return *localVarResponse.Data, localVarHttpResponse.Header, err
			}
		} else {
			return localVarReturnValue, localVarHttpResponse.Header, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v VideomakerVideocapturesAddResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse.Header, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse.Header, newErr
		}

		return localVarReturnValue, localVarHttpResponse.Header, newErr
	}

	return localVarReturnValue, localVarHttpResponse.Header, nil
}
