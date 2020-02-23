package plan

import (
	"context"
	"errors"
	"github.com/pepeunlimited/products/internal/pkg/clock"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/internal/pkg/ent/plan"
	"github.com/pepeunlimited/products/internal/pkg/ent/product"
	"time"
)

var (
	ErrPlanNotExist 			= errors.New("plans: plan not exist")
	ErrUnknownPlanUnit 			= errors.New("plans: unknown plan unit")
	ErrInvalidStartAt 			= errors.New("plans: invalid startAt")
	ErrInvalidEndAt 			= errors.New("plans: invalid endAt")
	ErrInvalidStartAtEndAt 		= errors.New("plans: startAt and endAt are equal")
)

type PlanRepository interface {
	CreatePlan(ctx context.Context, i18nId int64, length uint8, unit Unit, price uint16, discount uint16, productId int, startAt time.Time, endAt time.Time, price3rdId *int) (*ent.Plan, error)
	CreateNewPlan(ctx context.Context, i18nId int64, length uint8, unit Unit, price uint16, discount uint16, productId int, price3rdId *int) (*ent.Plan, error)
	EndPlanAt(ctx context.Context, month time.Month, day int, planId int) (*ent.Plan, error)

	GetPlanByID(ctx context.Context, planId int) (*ent.Plan, error)
	GetPlansByProductSkuAndTime(ctx context.Context, productSku string, now time.Time) ([]*ent.Plan, error)
	GetPlansByProductIdAndTime(ctx context.Context, productId int, now time.Time) ([]*ent.Plan, error)
	GetPlansByProductSku(ctx context.Context, productSku string) ([]*ent.Plan, error)
	GetPlansByProductId(ctx context.Context, productId int) ([]*ent.Plan, error)
	GetPlans(ctx context.Context) ([]*ent.Plan, error)
	GetPlansByTime(ctx context.Context, time time.Time) ([]*ent.Plan, error)

	Wipe(ctx context.Context)
}

type planMySQL struct {
	client *ent.Client
}

func (mysql planMySQL) EndPlanAt(ctx context.Context, month time.Month, day int, planId int) (*ent.Plan, error) {
	plan, err := mysql.GetPlanByID(ctx, planId)
	if err != nil {
		return nil, err
	}
	endAt, err := clock.ToMonthDate(month, day)
	if err != nil {
		return nil, err
	}
	save, err := plan.Update().SetEndAt(endAt).Save(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.GetPlanByID(ctx, save.ID)
}

func (mysql planMySQL) GetPlansByProductSkuAndTime(ctx context.Context, productSku string, now time.Time) ([]*ent.Plan, error) {
	return mysql.client.Plan.Query().WithThirdPartyPrices().WithProducts().Where(
		plan.And(
			plan.StartAtLTE(now),
			plan.EndAtGTE(now),
			plan.HasProductsWith(product.Sku(productSku)),
		)).All(ctx)
}

func (mysql planMySQL) GetPlansByProductIdAndTime(ctx context.Context, productId int, now time.Time) ([]*ent.Plan, error) {
	return mysql.client.Plan.Query().WithThirdPartyPrices().WithProducts().Where(
		plan.And(
			plan.StartAtLTE(now),
			plan.EndAtGTE(now),
			plan.HasProductsWith(product.ID(productId)),
		)).All(ctx)
}

func (mysql planMySQL) GetPlansByProductSku(ctx context.Context, productSku string) ([]*ent.Plan, error) {
	now := time.Now().UTC()
	return mysql.GetPlansByProductSkuAndTime(ctx, productSku, now)
}

func (mysql planMySQL) GetPlansByProductId(ctx context.Context, productId int) ([]*ent.Plan, error) {
	now := time.Now().UTC()
	return mysql.GetPlansByProductIdAndTime(ctx, productId, now)
}

func (mysql planMySQL) CreateNewPlan(ctx context.Context, i18nId int64, length uint8, unit Unit, price uint16, discount uint16, productId int, price3rdId *int) (*ent.Plan, error) {
	startAt := clock.ZeroAt()
	endAt   := clock.InfinityAt()
	return mysql.CreatePlan(ctx, i18nId, length, unit, price, discount, productId, startAt, endAt, price3rdId)
}

func (mysql planMySQL) GetPlans(ctx context.Context) ([]*ent.Plan, error) {
	now := time.Now().UTC()
	return mysql.GetPlansByTime(ctx, now)
}

func (mysql planMySQL) GetPlansByTime(ctx context.Context, now time.Time) ([]*ent.Plan, error) {
	query := mysql.client.Plan.Query().Where(plan.And(plan.StartAtLTE(now), plan.EndAtGTE(now)))
	return query.All(ctx)
}

func (mysql planMySQL) GetPlanByID(ctx context.Context, planId int) (*ent.Plan, error) {
	plans, err := mysql.client.Plan.Query().Where(plan.ID(planId)).WithThirdPartyPrices().WithProducts().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrPlanNotExist
		}
		return nil, err
	}
	return plans, nil
}

func (mysql planMySQL) Wipe(ctx context.Context) {
	mysql.client.Subscription.Delete().ExecX(ctx)
	mysql.client.Price.Delete().ExecX(ctx)
	mysql.client.Plan.Delete().ExecX(ctx)
	mysql.client.ThirdPartyPrice.Delete().ExecX(ctx)
	mysql.client.Product.Delete().ExecX(ctx)
}

func (mysql planMySQL) CreatePlan(ctx context.Context, i18nId int64, length uint8, unit Unit, price uint16, discount uint16, productId int, startAt time.Time, endAt time.Time, price3rdId *int) (*ent.Plan, error) {
	if endAt.Equal(startAt) {
		return nil, ErrInvalidStartAtEndAt
	}
	if endAt.Before(startAt) {
		return nil, ErrInvalidEndAt
	}
	if startAt.After(endAt) {
		return nil, ErrInvalidStartAt
	}
	if price3rdId != nil {
		if *price3rdId == 0 {
			price3rdId = nil
		}
	}
	if startAt.Year() != 1970 {
		startAt = startAt.Add(1 * time.Second)
	}
	plans, err := mysql.
		client.
		Plan.
		Create().
		SetLength(length).
		SetTitleI18nID(i18nId).
		SetUnit(unit.String()).
		SetPrice(price).
		SetDiscount(discount).
		SetProductsID(productId).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetNillableThirdPartyPricesID(price3rdId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.GetPlanByID(ctx, plans.ID)
}

func New(client *ent.Client) PlanRepository {
	return planMySQL{client:client}
}