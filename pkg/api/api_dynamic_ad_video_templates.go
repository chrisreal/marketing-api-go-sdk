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
	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	. "github.com/tencentad/marketing-api-go-sdk/pkg/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type DynamicAdVideoTemplatesApiService service

/*
DynamicAdVideoTemplatesApiService 获取动态商品视频模板
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param productCatalogId
 * @param adcreativeTemplateId
 * @param productMode
 * @param optional nil or *DynamicAdVideoTemplatesGetOpts - Optional Parameters:
     * @param "SupportChannel" (optional.Bool) -
     * @param "Page" (optional.Int64) -
     * @param "PageSize" (optional.Int64) -
     * @param "TemplateIdList" (optional.Interface of []int64) -
     * @param "TemplateName" (optional.String) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return DynamicAdVideoTemplatesGetResponse
*/

type DynamicAdVideoTemplatesGetOpts struct {
	SupportChannel optional.Bool
	Page           optional.Int64
	PageSize       optional.Int64
	TemplateIdList optional.Interface
	TemplateName   optional.String
	Fields         optional.Interface
}

func (a *DynamicAdVideoTemplatesApiService) Get(ctx context.Context, accountId int64, productCatalogId int64, adcreativeTemplateId int64, productMode string, localVarOptionals *DynamicAdVideoTemplatesGetOpts) (DynamicAdVideoTemplatesGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue DynamicAdVideoTemplatesGetResponseData
		localVarResponse    DynamicAdVideoTemplatesGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/dynamic_ad_video_templates/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	localVarQueryParams.Add("product_catalog_id", parameterToString(productCatalogId, ""))
	localVarQueryParams.Add("adcreative_template_id", parameterToString(adcreativeTemplateId, ""))
	localVarQueryParams.Add("product_mode", parameterToString(productMode, ""))
	if localVarOptionals != nil && localVarOptionals.SupportChannel.IsSet() {
		localVarQueryParams.Add("support_channel", parameterToString(localVarOptionals.SupportChannel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TemplateIdList.IsSet() {
		localVarQueryParams.Add("template_id_list", parameterToString(localVarOptionals.TemplateIdList.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.TemplateName.IsSet() {
		localVarQueryParams.Add("template_name", parameterToString(localVarOptionals.TemplateName.Value(), ""))
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
				var localVarResponseErrors []ApiErrorStruct
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
			var v DynamicAdVideoTemplatesGetResponse
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
