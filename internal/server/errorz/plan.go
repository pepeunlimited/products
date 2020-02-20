package errorz

import (
	"github.com/pepeunlimited/prices/internal/pkg/mysql/planrepo"
	"github.com/twitchtv/twirp"
	"log"
)

func IsPlanError(err error) error {
	switch err {
	case planrepo.ErrPlanNotExist:
		return twirp.NotFoundError("plan_not_found")
	case planrepo.ErrUnknownPlanUnit:
		return twirp.InvalidArgumentError("unit","unknown_plan_unit")
	}
	log.Print("plan-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}