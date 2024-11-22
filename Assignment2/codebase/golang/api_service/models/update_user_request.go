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

type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`

	Email string `json:"email,omitempty" validate:"regexp=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,}$"`
}
