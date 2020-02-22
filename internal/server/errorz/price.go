package errorz

import (
	"github.com/pepeunlimited/products/internal/pkg/mysql/price"
	"github.com/twitchtv/twirp"
	"log"
)

func Price(err error) error {
	switch err {
	case price.ErrInvalidEndAt:
		return twirp.InvalidArgumentError("end_at", "invalid_end_at")
	case price.ErrInvalidStartAt:
		return twirp.InvalidArgumentError("start_at", "invalid_start_at")
	case price.ErrInvalidProduct:
		return twirp.InvalidArgumentError("product_id", "invalid_product_id")
	case price.ErrPriceNotExist:
		return twirp.NotFoundError("price_not_found")
	}
	log.Print("price-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}
