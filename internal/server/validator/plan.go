package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	planrepo "github.com/pepeunlimited/products/internal/pkg/mysql/plan"
	"github.com/pepeunlimited/products/pkg/rpc/plan"
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
	fromString := planrepo.PlanUnitFromString(params.Unit)
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
	if params.ProductId == 0 && validator.IsEmpty(params.ProductSku) {
		return twirp.RequiredArgumentError("at_least_product_id")
	}
	return nil
}

func (v PlanServerValidator) EndPlanAt(params *plan.EndPlanAtParams) error {
	if params.EndAtDay == 0 {
		return twirp.RequiredArgumentError("end_at_day")
	}
	if params.EndAtMonth == 0 {
		return twirp.RequiredArgumentError("end_at_month")
	}
	if params.PlanId == 0 {
		return twirp.RequiredArgumentError("plan_id")
	}
	return nil
}

func NewPlanServerValidator() PlanServerValidator {
	return PlanServerValidator{}
}