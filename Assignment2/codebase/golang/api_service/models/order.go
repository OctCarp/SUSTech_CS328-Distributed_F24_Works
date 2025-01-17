/*
 * SUSTech Store API
 *
 * API service for SUSTech Store
 *
 * API version: 0.1.0
 * Contact: 12110304@mail.sustech.edu.cn
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Order struct {
	Id int32 `json:"id,omitempty"`

	UserId int32 `json:"user_id,omitempty"`

	ProductId int32 `json:"product_id,omitempty"`

	Quantity int32 `json:"quantity,omitempty"`

	TotalPrice float64 `json:"total_price,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`

	ProductName string `json:"product_name,omitempty"`

	UserName string `json:"user_name,omitempty"`
}
