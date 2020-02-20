package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/planrepo"
	"github.com/pepeunlimited/prices/pkg/planrpc"
	"github.com/twitchtv/twirp"
)

type PlanServerValidator struct {}

func (v PlanServerValidator) CreatePlan(params *planrpc.CreatePlanParams) error {
	if params.Length == 0 {
		return twirp.RequiredArgumentError("length")
	}
	if validator.IsEmpty(params.Unit) {
		return twirp.RequiredArgumentError("unit")
	}
	fromString := planrepo.PlanUnitFromString(params.Unit)
	if fromString.String() == "UNKNOWN" {
		return twirp.InvalidArgumentError("unit", "unknown_unit")
	}
	return nil
}

func (v PlanServerValidator) GetPlan(params *planrpc.GetPlanParams) error {
	if params.PlanId == 0 {
		return twirp.RequiredArgumentError("plan_id")
	}
	return nil
}

func NewPlanServerValidator() PlanServerValidator {
	return PlanServerValidator{}
}