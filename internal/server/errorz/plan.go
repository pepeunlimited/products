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
	case plan.ErrInvalidStartAt:
		return twirp.InvalidArgumentError("start_at","invalid")
	case plan.ErrInvalidEndAt:
		return twirp.InvalidArgumentError("end_at","invalid")
	case plan.ErrInvalidStartAtEndAt:
		return twirp.InvalidArgumentError("start_at_end_at","equal")
	}
	log.Print("plan-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}