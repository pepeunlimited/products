package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/prices/pkg/rpc/product"
	"github.com/twitchtv/twirp"
)

type ProductValidator struct {

}

func (v ProductValidator) CreateProduct(params *product.CreateProductParams) error {
	if validator.IsEmpty(params.Sku) {
		return twirp.RequiredArgumentError("sku")
	}
	return nil
}

func (v ProductValidator) GetProduct(params *product.GetProductParams) error {
	if params.ProductId == 0 && validator.IsEmpty(params.Sku) {
		return twirp.RequiredArgumentError("at_least_product_id")
	}
	return nil
}

func (v ProductValidator) GetProducts(params *product.GetProductsParams) error {
	if params.PageSize == 0 {
		return twirp.RequiredArgumentError("page_size")
	}
	return nil
}

func NewProductValidator() ProductValidator {
	return ProductValidator{}
}
