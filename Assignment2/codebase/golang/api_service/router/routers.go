/*
 * SUSTech Store API
 *
 * API service for SUSTech Store
 *
 * API version: 0.1.0
 * Contact: 12110304@mail.sustech.edu.cn
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package router

import (
	"net/http"
	"octcarp/sustech/cs328/a2/api/api/v1"
	"octcarp/sustech/cs328/a2/api/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc

	Protected bool
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouterWithGinEngine add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	loggerMiddleware := middleware.Logger()
	authMiddleware := middleware.JWTAuth()
	timeoutMiddleware := middleware.WithTimeout(time.Second * 5)

	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}

		var handlers []gin.HandlerFunc
		handlers = append(handlers, loggerMiddleware)
		if route.Protected {
			handlers = append(handlers, authMiddleware)
		}
		handlers = append(handlers, timeoutMiddleware)
		handlers = append(handlers, route.HandlerFunc)

		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, handlers...)
		case http.MethodPost:
			router.POST(route.Pattern, handlers...)
		case http.MethodPut:
			router.PUT(route.Pattern, handlers...)
		case http.MethodPatch:
			router.PATCH(route.Pattern, handlers...)
		case http.MethodDelete:
			router.DELETE(route.Pattern, handlers...)
		}
	}

	return router
}

// DefaultHandleFunc Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the GeneralAPI part of the API
	GeneralAPI v1.GeneralAPI
	// Routes for the OrdersAPI part of the API
	OrdersAPI v1.OrdersAPI
	// Routes for the ProductsAPI part of the API
	ProductsAPI v1.ProductsAPI
	// Routes for the UsersAPI part of the API
	UsersAPI v1.UsersAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{
		{
			"GetWelcomeMessage",
			http.MethodGet,
			"/",
			handleFunctions.GeneralAPI.GetWelcomeMessage,
			false,
		},
		{
			"CancelOrder",
			http.MethodDelete,
			"/orders/:id",
			handleFunctions.OrdersAPI.CancelOrder,
			true,
		},
		{
			"GetOrder",
			http.MethodGet,
			"/orders/:id",
			handleFunctions.OrdersAPI.GetOrder,
			true,
		},
		{
			"GetUserOrdersById",
			http.MethodGet,
			"/orders/user/:id",
			handleFunctions.OrdersAPI.GetUserOrdersById,
			true,
		},
		{
			"PlaceOrder",
			http.MethodPost,
			"/orders",
			handleFunctions.OrdersAPI.PlaceOrder,
			true,
		},
		{
			"GetProduct",
			http.MethodGet,
			"/products/:id",
			handleFunctions.ProductsAPI.GetProduct,
			false,
		},
		{
			"ListProducts",
			http.MethodGet,
			"/products",
			handleFunctions.ProductsAPI.ListProducts,
			false,
		},
		{
			"DeactivateUser",
			http.MethodDelete,
			"/users/:id",
			handleFunctions.UsersAPI.DeactivateUser,
			true,
		},
		{
			"GetUser",
			http.MethodGet,
			"/users/:id",
			handleFunctions.UsersAPI.GetUser,
			true,
		},
		{
			"LoginUser",
			http.MethodPost,
			"/users/login",
			handleFunctions.UsersAPI.LoginUser,
			false,
		},
		{
			"RegisterUser",
			http.MethodPost,
			"/users/register",
			handleFunctions.UsersAPI.RegisterUser,
			false,
		},
		{
			"UpdateUser",
			http.MethodPut,
			"/users/:id",
			handleFunctions.UsersAPI.UpdateUser,
			true,
		},
	}
}
