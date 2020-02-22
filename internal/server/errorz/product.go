package errorz

import (
	"github.com/pepeunlimited/prices/internal/pkg/mysql/product"
	"github.com/twitchtv/twirp"
	"log"
)

func Product(err error) error {
	switch err {
	case product.ErrProductNotExist:
		return twirp.NotFoundError("product_not_found")
	case product.ErrProductSkuExist:
		return twirp.NewError(twirp.AlreadyExists, "product_sku_already_exist")
	case product.ErrCantFigureIsSubscription:
		return twirp.NewError(twirp.Internal, "is_production_subscription_failed")
	}
	log.Print("product-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}