package validator

import (
	"github.com/pepeunlimited/prices/pkg/pricerpc"
	"github.com/twitchtv/twirp"
)

type PriceServerValidator struct {}

func (v PriceServerValidator) CreatePrice(params *pricerpc.CreatePriceParams) error {
	if params.Price == 0 {
		return twirp.RequiredArgumentError("price")
	}
	if params.Discount == 0 {
		return twirp.RequiredArgumentError("discount")
	}
	if params.ProductId == 0 {
		return twirp.RequiredArgumentError("product_id")
	}
	return nil
}

func NewPriceServerValidator() PriceServerValidator {
	return PriceServerValidator{}
}