package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/planrepo"
	"github.com/pepeunlimited/prices/internal/server/errorz"
	"github.com/pepeunlimited/prices/internal/server/validator"
	"github.com/pepeunlimited/prices/pkg/planrpc"
)

type PlanServer struct {
	plans planrepo.PlanRepository
	valid validator.PlanServerValidator
}

func (server PlanServer) CreatePlan(ctx context.Context, params *planrpc.CreatePlanParams) (*planrpc.Plan, error) {
	err := server.valid.CreatePlan(params)
	if err != nil {
		return nil, err
	}
	create, err := server.plans.Create(ctx, params.TitleI18NId, uint8(params.Length), planrepo.PlanUnitFromString(params.Unit))
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

func NewPlanServer(client *ent.Client) PlanServer {
	return PlanServer{
		plans:planrepo.NewPlanRepository(client),
		valid:validator.NewPlanServerValidator(),
	}
}