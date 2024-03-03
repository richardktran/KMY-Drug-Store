package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/KMY-Drug-Store/app/services/contracts"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type ProductController struct {
	productService contracts.IProductService
}

func NewProductController(
	productService contracts.IProductService,
) ProductController {
	return ProductController{
		productService: productService,
	}
}

func (ctl *ProductController) GetProduct() func(*gin.Context) {
	return func(c *gin.Context) {
		productName := c.DefaultQuery("name", "")

		if productName == "" {
			app.ResponseBadRequest(
				app.ThrowBadRequestError(errors.New("product_name_is_required"), "product_name_is_required"),
			).Context(c)

			return
		}

		product, err := ctl.productService.GetProductByName(productName)

		if err != nil {
			app.ResponseNotFound(
				app.ThrowNotFoundError(err, "product_not_found"),
			).Context(c)

			return
		}

		app.ResponseSuccess(product).Context(c)
	}
}
