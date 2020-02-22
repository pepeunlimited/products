package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/plan"
	"github.com/pepeunlimited/prices/internal/server/errorz"
	"github.com/pepeunlimited/prices/internal/server/validator"
	"github.com/pepeunlimited/prices/pkg/planrpc"
	"github.com/twitchtv/twirp"
	"time"
)

type PlanServer struct {
	plans plan.PlanRepository
	valid validator.PlanServerValidator
}

func (server PlanServer) CreatePlan(ctx context.Context, params *planrpc.CreatePlanParams) (*planrpc.Plan, error) {
	err := server.valid.CreatePlan(params)
	if err != nil {
		return nil, err
	}
	create, err := server.plans.Create(ctx, params.TitleI18NId, uint8(params.Length), plan.PlanUnitFromString(params.Unit))
	if err != nil {
		return nil, errorz.Plan(err)
	}
	return ToPlan(create), nil
}

func (server PlanServer) GetPlans(ctx context.Context, params *planrpc.GetPlansParams) (*planrpc.GetPlansResponse, error) {
	err := server.valid.GetPlans(params)
	if err != nil {
		return nil, err
	}
	plans, err := server.plans.GetPlans(ctx, params.Show)
	if err != nil {
		return nil, errorz.Plan(err)
	}
	return &planrpc.GetPlansResponse{Plans:ToPlans(plans)}, nil
}

func (server PlanServer) GetPlan(ctx context.Context, params *planrpc.GetPlanParams) (*planrpc.Plan, error) {
	err := server.valid.GetPlan(params)
	if err != nil {
		return nil, err
	}
	plan, err := server.plans.GetPlanByID(ctx, int(params.PlanId))
	if err != nil {
		return nil, errorz.Plan(err)
	}
	return ToPlan(plan), nil
}

func (server PlanServer) endAt(length int, startAt time.Time, unit plan.Unit) (time.Time, error) {
	switch unit {
	case plan.Days:
		return startAt.AddDate(0, 0, length).UTC(), nil
	case plan.Weeks:
		return startAt.AddDate(0, 0, length * 7).UTC(), nil
	case plan.Months:
		return startAt.AddDate(0, length, 0).UTC(),nil
	case plan.Years:
		return startAt.AddDate(length, 0, 0).UTC(),nil
	default:
		return time.Time{}, twirp.InvalidArgumentError("unit", "unknown")
	}
}

func NewPlanServer(client *ent.Client) PlanServer {
	return PlanServer{
		plans: plan.NewPlanRepository(client),
		valid: validator.NewPlanServerValidator(),
	}
}