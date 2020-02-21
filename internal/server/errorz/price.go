package errorz

import (
	"github.com/pepeunlimited/prices/internal/pkg/mysql/pricerepo"
	"github.com/twitchtv/twirp"
	"log"
)

func Price(err error) error {
	switch err {
	case pricerepo.ErrInvalidEndAt:
		return twirp.InvalidArgumentError("end_at", "invalid_end_at")
	case pricerepo.ErrInvalidStartAt:
		return twirp.InvalidArgumentError("start_at", "invalid_start_at")
	case pricerepo.ErrPriceNotExist:
		return twirp.NotFoundError("price_not_found")
	}
	log.Print("price-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}
