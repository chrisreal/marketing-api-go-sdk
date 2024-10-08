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

type MarketingTargetAssetPropertyValuesApiService service

/*
MarketingTargetAssetPropertyValuesApiService 获取可用的推广内容资产属性值
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param organizationId
 * @param marketingTargetType
 * @param optional nil or *MarketingTargetAssetPropertyValuesGetOpts - Optional Parameters:
     * @param "MarketingAssetType" (optional.String) -
     * @param "MarketingAssetCategory" (optional.String) -
     * @param "PropertyName" (optional.String) -
     * @param "Filtering" (optional.Interface of []FilteringStruct) -
     * @param "Page" (optional.Int64) -
     * @param "PageSize" (optional.Int64) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return MarketingTargetAssetPropertyValuesGetResponse
*/

type MarketingTargetAssetPropertyValuesGetOpts struct {
	MarketingAssetType     optional.String
	MarketingAssetCategory optional.String
	PropertyName           optional.String
	Filtering              optional.Interface
	Page                   optional.Int64
	PageSize               optional.Int64
	Fields                 optional.Interface
}

func (a *MarketingTargetAssetPropertyValuesApiService) Get(ctx context.Context, organizationId int64, marketingTargetType string, localVarOptionals *MarketingTargetAssetPropertyValuesGetOpts) (MarketingTargetAssetPropertyValuesGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue MarketingTargetAssetPropertyValuesGetResponseData
		localVarResponse    MarketingTargetAssetPropertyValuesGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/marketing_target_asset_property_values/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("organization_id", parameterToString(organizationId, ""))
	localVarQueryParams.Add("marketing_target_type", parameterToString(marketingTargetType, ""))
	if localVarOptionals != nil && localVarOptionals.MarketingAssetType.IsSet() {
		localVarQueryParams.Add("marketing_asset_type", parameterToString(localVarOptionals.MarketingAssetType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MarketingAssetCategory.IsSet() {
		localVarQueryParams.Add("marketing_asset_category", parameterToString(localVarOptionals.MarketingAssetCategory.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PropertyName.IsSet() {
		localVarQueryParams.Add("property_name", parameterToString(localVarOptionals.PropertyName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filtering.IsSet() {
		localVarQueryParams.Add("filtering", parameterToString(localVarOptionals.Filtering.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
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
			var v MarketingTargetAssetPropertyValuesGetResponse
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
