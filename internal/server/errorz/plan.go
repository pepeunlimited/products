package errorz

import (
	"github.com/pepeunlimited/products/internal/pkg/mysql/plan"
	"github.com/twitchtv/twirp"
	"log"
)

func Plan(err error) error {
	switch err {
	case plan.ErrPlanNotExist:
		return twirp.NotFoundError("plan_not_found")
	case plan.ErrUnknownPlanUnit:
		return twirp.InvalidArgumentError("unit","unknown_plan_unit")
	}
	log.Print("plan-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}