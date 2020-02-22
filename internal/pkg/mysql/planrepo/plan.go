package planrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/ent/plan"
	"github.com/pepeunlimited/prices/internal/pkg/ent/price"
	"time"
)

var (
	ErrPlanNotExist 			= errors.New("plans: plan not exist")
	ErrUnknownPlanUnit 			= errors.New("plans: unknown plan unit")
)

type PlanRepository interface {
	Create(ctx context.Context, titleI18nId int64, length uint8, unit Unit) (*ent.Plan, error)

	EndAtByPlanID(ctx context.Context, startAt time.Time, planId int) (time.Time, error)

	GetPlanByID(ctx context.Context, plansID int) 						 (*ent.Plan, error)
	GetPlans(ctx context.Context, show bool) 							 ([]*ent.Plan, error)
	GetPlansByTime(ctx context.Context, time time.Time, show bool) 	     ([]*ent.Plan, error)

	Wipe(ctx context.Context)
}

type planMySQL struct {
	client *ent.Client
}

func (mysql planMySQL) GetPlans(ctx context.Context, show bool) ([]*ent.Plan, error) {
	now := time.Now().UTC()
	return mysql.GetPlansByTime(ctx, now, show)
}

func (mysql planMySQL) GetPlansByTime(ctx context.Context, now time.Time, show bool) ([]*ent.Plan, error) {
	query := mysql.client.Plan.Query()
	if !show {
		query.Where(plan.HasPricesWith(price.And(price.StartAtLTE(now), price.EndAtGTE(now), price.HasPlans())))
	}
	return query.All(ctx)
}

func (mysql planMySQL) GetPlanByID(ctx context.Context, plansID int) (*ent.Plan, error) {
	plans, err := mysql.client.Plan.Get(ctx, plansID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrPlanNotExist
		}
		return nil, err
	}
	return plans, nil
}

func (mysql planMySQL) EndAtByPlanID(ctx context.Context, startAt time.Time, planId int) (time.Time, error) {
	plans, err := mysql.GetPlanByID(ctx, planId)
	if err != nil {
		return time.Time{}, err
	}
	unit := PlanUnitFromString(plans.Unit)
	endAt := startAt
	switch unit {
	case Days:
		return endAt.AddDate(0, 0, int(plans.Length)).UTC(), nil
	case Weeks:
		return endAt.AddDate(0, 0, int(plans.Length) * 7).UTC(), nil
	case Months:
		return endAt.AddDate(0, int(plans.Length), 0).UTC(),nil
	case Years:
		return endAt.AddDate(int(plans.Length), 0, 0).UTC(),nil
	}
	return time.Time{}, ErrUnknownPlanUnit
}

func (mysql planMySQL) Wipe(ctx context.Context) {
	mysql.client.Subscription.Delete().ExecX(ctx)
	mysql.client.Price.Delete().ExecX(ctx)
	mysql.client.ThirdParty.Delete().ExecX(ctx)
	mysql.client.Plan.Delete().ExecX(ctx)
	mysql.client.Product.Delete().ExecX(ctx)
}

func (mysql planMySQL) Create(ctx context.Context, titleI18nId int64, length uint8, unit Unit) (*ent.Plan, error) {
	plans, err := mysql.
		client.
		Plan.
		Create().
		SetLength(length).
		SetTitleI18nID(titleI18nId).
		SetUnit(unit.String()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func NewPlanRepository(client *ent.Client) PlanRepository {
	return planMySQL{client:client}
}