package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	thirdpartypricerepo "github.com/pepeunlimited/products/internal/pkg/mysql/thirdpartyprice"
	"github.com/pepeunlimited/products/pkg/rpc/thirdpartyprice"
	"github.com/twitchtv/twirp"
)

type ThirdPartyServerValidator struct {}

func NewThirdPartyServerValidator() ThirdPartyServerValidator {
	return ThirdPartyServerValidator{}
}


func (ThirdPartyServerValidator) CreateThirdParty(params *thirdpartyprice.CreateThirdPartyPriceParams) error {
	//if validator.IsEmpty(params.GoogleBillingServiceSku) {
	//	return twirp.RequiredArgumentError("google_billing_service_sku")
	//}
	if validator.IsEmpty(params.InAppPurchaseSku) {
		return twirp.RequiredArgumentError("in_app_purchase_sku")
	}
	if validator.IsEmpty(params.Type) {
		return twirp.RequiredArgumentError("type")
	}
	if thirdpartypricerepo.ThirdPartyPriceTypeFromString(params.Type) == thirdpartypricerepo.Unknown {
		return twirp.InvalidArgumentError("type", "unknown")
	}
	return nil
}

func (v ThirdPartyServerValidator) GetThirdParty(params *thirdpartyprice.GetThirdPartyPriceParams) error {
	if validator.IsEmpty(params.GoogleBillingServiceSku) &&
		validator.IsEmpty(params.InAppPurchaseSku) && params.ThirdPartyPriceId == 0 {
		return twirp.RequiredArgumentError("at_least_third_party_id")
	}
	return nil
}

func (v ThirdPartyServerValidator) EndThirdParty(params *thirdpartyprice.EndThirdPartyPriceParams) error {
	if params.Params == nil {
		return twirp.RequiredArgumentError("params")
	}
	err := v.GetThirdParty(params.Params)
	if err != nil {
		return err
	}
	if params.EndAtDay == 0 {
		return twirp.RequiredArgumentError("end_at_day")
	}
	if params.EndAtMonth == 0 {
		return twirp.RequiredArgumentError("end_at_month")
	}
	return nil
}

func (v ThirdPartyServerValidator) GetThirdPartyPrices(params *thirdpartyprice.GetThirdPartyPricesParams) error {
	return nil
}