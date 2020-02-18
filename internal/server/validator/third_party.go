package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/prices/pkg/thirdpartyrpc"
	"github.com/twitchtv/twirp"
)

type ThirdPartyServerValidator struct {}

func NewThirdPartyServerValidator() ThirdPartyServerValidator {
	return ThirdPartyServerValidator{}
}


func (ThirdPartyServerValidator) CreateThirdParty(params *thirdpartyrpc.CreateThirdPartyParams) error {
	//if validator.IsEmpty(params.GoogleBillingServiceSku) {
	//	return twirp.RequiredArgumentError("google_billing_service_sku")
	//}
	if validator.IsEmpty(params.InAppPurchaseSku) {
		return twirp.RequiredArgumentError("in_app_purchase_sku")
	}
	if params.StartAtDay > 31 || params.StartAtDay < 0 {
		return twirp.InvalidArgumentError("start_at_day", "invalid start_at_day")
	}
	if params.StartAtMonth > 12 || params.StartAtMonth < 0 {
		return twirp.InvalidArgumentError("start_at_month", "invalid start_at_month")
	}
	return nil
}

func (v ThirdPartyServerValidator) GetThirdParty(params *thirdpartyrpc.GetThirdPartyParams) error {
	if validator.IsEmpty(params.GoogleBillingServiceSku) &&
		validator.IsEmpty(params.InAppPurchaseSku) && params.Id == 0 {
		return twirp.RequiredArgumentError("id_or_googlesku_or_applesku")
	}
	return nil
}