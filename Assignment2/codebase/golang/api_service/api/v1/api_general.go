/*
 * SUSTech Store API
 *
 * API service for SUSTech Store
 *
 * API version: 0.1.0
 * Contact: 12110304@mail.sustech.edu.cn
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"octcarp/sustech/cs328/a2/api/utils"
)

type GeneralAPI struct {
}

// GetWelcomeMessage Get /
// Get welcome message
func (api *GeneralAPI) GetWelcomeMessage(c *gin.Context) {
	utils.SendCodeWithMessage(c, http.StatusOK,
		"Welcome to SUSTech Store API",
	)
}
