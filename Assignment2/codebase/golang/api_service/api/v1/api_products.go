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
	"octcarp/sustech/cs328/a2/api/grpc/dbclient"
	"octcarp/sustech/cs328/a2/api/models"
	"octcarp/sustech/cs328/a2/api/utils"
	dbpb "octcarp/sustech/cs328/a2/gogrpc/dbs/pb"
	"strconv"
)

type ProductsAPI struct {
}

// GetProduct Get /products/:id
// Get product by ID
func (api *ProductsAPI) GetProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequestErr(c, "Invalid product ID")
		return
	}

	req := &dbpb.GetProductRequest{
		ProductId: int32(productID),
	}
	res, err := dbclient.GetDbClient().GetProduct(c.Request.Context(), req)

	if err != nil {
		utils.SendDbErr(c, err.Error())
		return
	}

	if res == nil {
		utils.SendNotFoundErr(c)
		return
	}

	product := models.Product{
		Id:          res.Id,
		Name:        res.Name,
		Description: res.Description,
		Category:    res.Category,
		Price:       res.Price,
		Slogan:      res.Slogan,
		Stock:       res.Stock,
		CreatedAt:   res.CreatedAt,
	}

	utils.ResponseLog(c, http.StatusOK, "Get product success")
	c.JSON(http.StatusOK, product)
}

// ListProducts Get /products
// List all products
func (api *ProductsAPI) ListProducts(c *gin.Context) {
	req := &dbpb.ListProductsRequest{}

	res, err := dbclient.GetDbClient().ListProducts(c.Request.Context(), req)
	if err != nil {
		utils.SendDbErr(c, err.Error())
		return
	}

	products := make([]models.Product, 0)
	for _, p := range res.Products {
		product := models.Product{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Category:    p.Category,
			Price:       p.Price,
			Slogan:      p.Slogan,
			Stock:       p.Stock,
			CreatedAt:   p.CreatedAt,
		}
		products = append(products, product)
	}

	utils.ResponseLog(c, http.StatusOK, "List products success")
	c.JSON(http.StatusOK, products)
}
