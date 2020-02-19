package errorz

import (
	"github.com/pepeunlimited/prices/internal/pkg/mysql/thirdpartyrepo"
	"github.com/twitchtv/twirp"
	"log"
)

type ThirdPartyErrorz struct {}

func (ThirdPartyErrorz) IsThirdPartyError(err error) error {
	switch err {
	case thirdpartyrepo.ErrThirdPartyNotExist:
		return twirp.NotFoundError("third_party_not_found")
	case thirdpartyrepo.ErrGoogleBillingServiceSkuExist:
		return twirp.NewError(twirp.AlreadyExists, "google_billing_service_sku_already_exist")
	case thirdpartyrepo.ErrInAppPurchaseSkuExist:
		return twirp.NewError(twirp.AlreadyExists, "in_app_purchase_sku_already_exist")
	}
	log.Print("third-party-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}

func NewThirdPartyErrorz() ThirdPartyErrorz {
	return ThirdPartyErrorz{}
}