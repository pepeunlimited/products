package planrepo

import (
	"context"
	"errors"
	"time"
)

var (
	ErrPlanNotExist 			= errors.New("plans: plan not exist")
	ErrUnknownPlanUnit 			= errors.New("plans: unknown plan unit")
	ErrPriceIDAlreadyExist   	= errors.New("plans: priceId already exist")
)

type PlanRepository interface {
	Create(ctx context.Context, startAt time.Time, endAt time.Time, titleI18nId int64, priceId int64, length uint8, unit Unit) (*ent.Plan, error)

	LengthByPlansID(ctx context.Context, startAt time.Time, plansID int) (time.Time, error)

	GetPlansByID(ctx context.Context, plansID int) 						 (*ent.Plan, error)
	GetPlansByPriceID(ctx context.Context, priceID int64)				 (*ent.Plan, error)
	GetPlans(ctx context.Context, active bool) 							 ([]*ent.Plan, error)

	Wipe(ctx context.Context)
}

type planMySQL struct {
	client *ent.Client
}

func (mysql planMySQL) GetPlans(ctx context.Context, active bool) ([]*ent.Plan, error) {
	query := mysql.client.Plan.Query()
	now := time.Now().UTC()
	if active {
		query.Where(plan.And(
			plan.StartAtLTE(now),
			plan.EndAtGTE(now)))
	}
	return query.All(ctx)
}

func (mysql planMySQL) GetPlansByPriceID(ctx context.Context, plansID int64) (*ent.Plan, error) {
	plan, err := mysql.client.Plan.Query().Where(plan.PriceID(plansID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrPlanNotExist
		}
		return nil, err
	}
	return plan, nil
}

func (mysql planMySQL) GetPlansByID(ctx context.Context, plansID int) (*ent.Plan, error) {
	plans, err := mysql.client.Plan.Get(ctx, plansID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrPlanNotExist
		}
		return nil, err
	}
	return plans, nil
}

func (mysql planMySQL) LengthByPlansID(ctx context.Context, startAt time.Time, plansID int) (time.Time, error) {
	plans, err := mysql.GetPlansByID(ctx, plansID)
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
	mysql.client.Plan.Delete().ExecX(ctx)
}

func (mysql planMySQL) Create(ctx context.Context, startAt time.Time, endAt time.Time, titleI18nId int64, priceId int64, length uint8, unit Unit) (*ent.Plan, error) {
	plans, err := mysql.
		client.
		Plan.
		Create().
		SetStartAt(startAt.UTC()).
		SetEndAt(endAt.UTC()).
		SetLength(length).
		SetPriceID(priceId).
		SetTitleI18nID(titleI18nId).
		SetUnit(unit.String()).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, ErrPriceIDAlreadyExist
		}
		return nil, err
	}
	return plans, nil
}

func NewPlanRepository(client *ent.Client) PlanRepository {
	return planMySQL{client:client}
}