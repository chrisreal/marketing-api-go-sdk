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

type TargetingTagReportsApiService service

/*
TargetingTagReportsApiService 获取定向标签报表
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param type_
 * @param level
 * @param dateRange
 * @param optional nil or *TargetingTagReportsGetOpts - Optional Parameters:
     * @param "PosType" (optional.String) -
     * @param "Filtering" (optional.Interface of []FilteringStruct) -
     * @param "GroupBy" (optional.Interface of []string) -
     * @param "OrderBy" (optional.Interface of []OrderByStruct) -
     * @param "Page" (optional.Int64) -
     * @param "PageSize" (optional.Int64) -
     * @param "TimeLine" (optional.String) -
     * @param "WeixinOfficialAccountsUpgradeEnabled" (optional.Bool) -
     * @param "AdqAccountsUpgradeEnabled" (optional.Bool) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return TargetingTagReportsGetResponse
*/

type TargetingTagReportsGetOpts struct {
	PosType                              optional.String
	Filtering                            optional.Interface
	GroupBy                              optional.Interface
	OrderBy                              optional.Interface
	Page                                 optional.Int64
	PageSize                             optional.Int64
	TimeLine                             optional.String
	WeixinOfficialAccountsUpgradeEnabled optional.Bool
	AdqAccountsUpgradeEnabled            optional.Bool
	Fields                               optional.Interface
}

func (a *TargetingTagReportsApiService) Get(ctx context.Context, accountId int64, type_ string, level string, dateRange ReportDateRange, localVarOptionals *TargetingTagReportsGetOpts) (TargetingTagReportsGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue TargetingTagReportsGetResponseData
		localVarResponse    TargetingTagReportsGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/targeting_tag_reports/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	localVarQueryParams.Add("type", parameterToString(type_, ""))
	localVarQueryParams.Add("level", parameterToString(level, ""))
	localVarQueryParams.Add("date_range", parameterToString(dateRange, ""))
	if localVarOptionals != nil && localVarOptionals.PosType.IsSet() {
		localVarQueryParams.Add("pos_type", parameterToString(localVarOptionals.PosType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filtering.IsSet() {
		localVarQueryParams.Add("filtering", parameterToString(localVarOptionals.Filtering.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.GroupBy.IsSet() {
		localVarQueryParams.Add("group_by", parameterToString(localVarOptionals.GroupBy.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.OrderBy.IsSet() {
		localVarQueryParams.Add("order_by", parameterToString(localVarOptionals.OrderBy.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimeLine.IsSet() {
		localVarQueryParams.Add("time_line", parameterToString(localVarOptionals.TimeLine.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WeixinOfficialAccountsUpgradeEnabled.IsSet() {
		localVarQueryParams.Add("weixin_official_accounts_upgrade_enabled", parameterToString(localVarOptionals.WeixinOfficialAccountsUpgradeEnabled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AdqAccountsUpgradeEnabled.IsSet() {
		localVarQueryParams.Add("adq_accounts_upgrade_enabled", parameterToString(localVarOptionals.AdqAccountsUpgradeEnabled.Value(), ""))
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
			var v TargetingTagReportsGetResponse
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
