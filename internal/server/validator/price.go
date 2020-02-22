package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/products/pkg/rpc/price"
	"github.com/twitchtv/twirp"
)

type PriceServerValidator struct {}

func (v PriceServerValidator) CreatePrice(params *price.CreatePriceParams) error {
	if params.Price == 0 {
		return twirp.RequiredArgumentError("price")
	}
	//if params.Discount == 0 {
	//	return twirp.RequiredArgumentError("discount")
	//}
	if params.ProductId == 0 {
		return twirp.RequiredArgumentError("product_id")
	}
	return nil
}

func (v PriceServerValidator) EndPrice(params *price.EndPriceAtParams) error {
	if params.EndAtDay == 0 {
		return twirp.RequiredArgumentError("end_at_day")
	}
	if params.EndAtMonth == 0 {
		return twirp.RequiredArgumentError("end_at_month")
	}
	if params.Params == nil {
		return twirp.RequiredArgumentError("params")
	}
	return v.GetPrice(params.Params)
}

func (v PriceServerValidator) GetPrice(params *price.GetPriceParams) error {
	if params.ProductId  == 0 &&
		params.PriceId == 0 &&
		validator.IsEmpty(params.ProductSku) {
		return twirp.RequiredArgumentError("at_least_price_id")
	}
	return nil
}

func NewPriceServerValidator() PriceServerValidator {
	return PriceServerValidator{}
}