package twirp

import (
	"context"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/products/internal/pkg/clock"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	planrepo "github.com/pepeunlimited/products/internal/pkg/mysql/plan"
	"github.com/pepeunlimited/products/internal/pkg/mysql/product"
	"github.com/pepeunlimited/products/internal/pkg/mysql/thirdpartyprice"
	"github.com/pepeunlimited/products/internal/server/errorz"
	"github.com/pepeunlimited/products/internal/server/validator"
	"github.com/pepeunlimited/products/pkg/rpc/plan"
	"time"
)

type PlanServer struct {
	plans planrepo.PlanRepository
	valid validator.PlanServerValidator
	products product.ProductRepository
	thirdpartyprices thirdpartyprice.ThirdPartyPriceRepository
}

func (server PlanServer) EndPlanAt(ctx context.Context, params *plan.EndPlanAtParams) (*plan.Plan, error) {
	if err := server.valid.EndPlanAt(params); err != nil {
		return nil, err
	}
	endAt, err := clock.ToMonthDate(time.Month(params.EndAtMonth), int(params.EndAtDay))
	if err != nil {
		return nil, err
	}
	_, err = server.GetPlan(ctx, &plan.GetPlanParams{
		PlanId: params.PlanId,
	})
	if err != nil {
		return nil, err
	}
	plan, err := server.plans.EndPlanAt(ctx, endAt.Month(), endAt.Day(), int(params.PlanId))
	if err != nil {
		return nil, errorz.Plan(err)
	}
	return ToPlan(plan), nil
}

func (server PlanServer) CreatePlan(ctx context.Context, params *plan.CreatePlanParams) (*plan.Plan, error) {
	if err := server.valid.CreatePlan(params); err != nil {
		return nil, err
	}
	_, err := server.products.GetProductByID(ctx, false, false, int(params.ProductId))
	if err != nil {
		return nil, errorz.Product(err)
	}
	var startAt time.Time
	var endAt 	time.Time
	if params.StartAtMonth != 0 && params.StartAtDay != 0 {
		startAt, err = clock.ToMonthDate(time.Month(params.StartAtMonth), int(params.StartAtDay))
	} else {
		startAt = clock.ZeroAt()
	}
	if params.EndAtMonth != 0 && params.EndAtDay != 0 {
		endAt, err = clock.ToMonthDate(time.Month(params.EndAtMonth), int(params.EndAtDay))
	} else {
		endAt = clock.InfinityAt()
	}
	if params.ThirdPartyPriceId != 0 {
		_, err := server.thirdpartyprices.GetByID(ctx, int(params.ThirdPartyPriceId))
		if err != nil {
			return nil, errorz.ThirdParty(err)
		}
	}
	thirdpartypriceId := int(params.ThirdPartyPriceId)
	plan, err := server.plans.CreatePlan(ctx, params.TitleI18NId, uint8(params.Length), planrepo.PlanUnitFromString(params.Unit), params.Price, params.Discount, int(params.ProductId), startAt, endAt, &thirdpartypriceId)
	if err != nil {
		return nil, errorz.Plan(err)
	}
	return ToPlan(plan), nil
}

func (server PlanServer) GetPlans(ctx context.Context, params *plan.GetPlansParams) (*plan.GetPlansResponse, error) {
	if err := server.valid.GetPlans(params); err != nil {
		return nil, err
	}
	var err error
	var plans []*ent.Plan
	if params.ProductId != 0 {
		plans, err = server.plans.GetPlansByProductId(ctx, int(params.ProductId))
	}
	if !validator2.IsEmpty(params.ProductSku) {
		plans, err = server.plans.GetPlansByProductSku(ctx, params.ProductSku)
	}
	if err != nil {
		return nil, errorz.Plan(err)
	}
	return &plan.GetPlansResponse{Plans: ToPlans(plans)}, nil
}

func (server PlanServer) GetPlan(ctx context.Context, params *plan.GetPlanParams) (*plan.Plan, error) {
	if err := server.valid.GetPlan(params); err != nil {
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
		plans: planrepo.New(client),
		valid: validator.NewPlanServerValidator(),
		products:product.New(client),
		thirdpartyprices: thirdpartyprice.New(client),
	}
}