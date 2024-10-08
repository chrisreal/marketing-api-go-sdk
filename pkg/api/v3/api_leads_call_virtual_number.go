/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/errors"
	"github.com/chrisreal/marketing-api-go-sdk/pkg/model"
	. "github.com/chrisreal/marketing-api-go-sdk/pkg/model/v3"
)

// Linger please
var (
	_ context.Context
)

type LeadsCallVirtualNumberApiService service

/*
LeadsCallVirtualNumberApiService 获取中间号
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param caller
 * @param callee
 * @param optional nil or *LeadsCallVirtualNumberGetOpts - Optional Parameters:
     * @param "LeadsId" (optional.Int64) -
     * @param "OuterLeadsId" (optional.String) -
     * @param "RequestId" (optional.String) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return LeadsCallVirtualNumberGetResponse
*/

type LeadsCallVirtualNumberGetOpts struct {
	LeadsId      optional.Int64
	OuterLeadsId optional.String
	RequestId    optional.String
	Fields       optional.Interface
}

func (a *LeadsCallVirtualNumberApiService) Get(ctx context.Context, accountId int64, caller string, callee string, localVarOptionals *LeadsCallVirtualNumberGetOpts) (LeadsCallVirtualNumberGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue LeadsCallVirtualNumberGetResponseData
		localVarResponse    LeadsCallVirtualNumberGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/leads_call_virtual_number/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	if localVarOptionals != nil && localVarOptionals.LeadsId.IsSet() {
		localVarQueryParams.Add("leads_id", parameterToString(localVarOptionals.LeadsId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OuterLeadsId.IsSet() {
		localVarQueryParams.Add("outer_leads_id", parameterToString(localVarOptionals.OuterLeadsId.Value(), ""))
	}
	localVarQueryParams.Add("caller", parameterToString(caller, ""))
	localVarQueryParams.Add("callee", parameterToString(callee, ""))
	if localVarOptionals != nil && localVarOptionals.RequestId.IsSet() {
		localVarQueryParams.Add("request_id", parameterToString(localVarOptionals.RequestId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"text/plain"}

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
			var v LeadsCallVirtualNumberGetResponse
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
