/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type SubcustomerTransferAddRequest struct {
	AccountId      *int64         `json:"account_id,omitempty"`
	Amount         *int64         `json:"amount,omitempty"`
	ToAccountId    *int64         `json:"to_account_id,omitempty"`
	FundType       AccountTypeMap `json:"fund_type,omitempty"`
	ExternalBillNo *string        `json:"external_bill_no,omitempty"`
	Memo           *string        `json:"memo,omitempty"`
}
