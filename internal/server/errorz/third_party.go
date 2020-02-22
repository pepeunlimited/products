package errorz

import (
	"github.com/pepeunlimited/products/internal/pkg/mysql/thirdpartyprice"
	"github.com/twitchtv/twirp"
	"log"
)

func ThirdParty(err error) error {
	switch err {
	case thirdpartyprice.ErrThirdPartyPriceNotExist:
		return twirp.NotFoundError("third_party_price_not_found")
	case thirdpartyprice.ErrGoogleBillingServiceSkuExist:
		return twirp.NewError(twirp.AlreadyExists, "google_billing_service_sku_already_exist")
	case thirdpartyprice.ErrInAppPurchaseSkuExist:
		return twirp.NewError(twirp.AlreadyExists, "in_app_purchase_sku_already_exist")
	}
	log.Print("third-party-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}