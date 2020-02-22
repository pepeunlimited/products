package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/plan"
	"github.com/pepeunlimited/prices/pkg/rpc/plan"
	"github.com/twitchtv/twirp"
)

type PlanServerValidator struct {}

func (v PlanServerValidator) CreatePlan(params *plan.CreatePlanParams) error {
	if params.Length == 0 {
		return twirp.RequiredArgumentError("length")
	}
	if validator.IsEmpty(params.Unit) {
		return twirp.RequiredArgumentError("unit")
	}
	fromString := plan.PlanUnitFromString(params.Unit)
	if fromString.String() == "UNKNOWN" {
		return twirp.InvalidArgumentError("unit", "unknown_unit")
	}
	return nil
}

func (v PlanServerValidator) GetPlan(params *plan.GetPlanParams) error {
	if params.PlanId == 0 {
		return twirp.RequiredArgumentError("plan_id")
	}
	return nil
}

func (v PlanServerValidator) GetPlans(params *plan.GetPlansParams) error {
	return nil
}

func NewPlanServerValidator() PlanServerValidator {
	return PlanServerValidator{}
}