package pricerepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/ent/plan"
	"github.com/pepeunlimited/prices/internal/pkg/ent/price"
	"github.com/pepeunlimited/prices/internal/pkg/ent/product"
	"time"
)

var (
	ErrPriceNotExist 		= errors.New("prices: not exist")
	ErrInvalidStartAt 		= errors.New("prices: invalid startAt")
	ErrInvalidEndAt 		= errors.New("prices: invalid endAt")
)

type PriceRepository interface {
	GetPriceByProductID(ctx context.Context, productId int, withProduct bool, withPlan bool, withThirdParty bool) 						  					  (*ent.Price, error)
	GetPriceByProductIDAndTime(ctx context.Context, productId int, now time.Time, withProduct bool, withPlan bool, withThirdParty bool) 						  (*ent.Price, error)
	GetPricesByProductID(ctx context.Context, productId int, isSequence bool) 							  ([]*ent.Price, error)
	GetPriceByID(ctx context.Context, id int, withProduct bool, withPlan bool, withThirdParty bool) 	  (*ent.Price, error)
	GetPricesByPlanId(ctx context.Context, planId int) 							  					  	  ([]*ent.Price, error)
	GetPriceByPlanId(ctx context.Context, planId int, withProduct bool, withPlan bool, withThirdParty bool) 							  					  	  (*ent.Price, error)
	GetPriceByPlanIdAndTime(ctx context.Context, planId int, now time.Time, withProduct bool, withPlan bool, withThirdParty bool) 							  (*ent.Price, error)

	CreateNewPrice(ctx context.Context, price uint16, discount uint16, productId int, thirdPartyID *int, plansId *int)  	  (*ent.Price, error)
	CreatePrice(ctx context.Context, price uint16, discount uint16, productId int, startAt time.Time, endAt time.Time, thirdPartyID *int, plansId *int) (*ent.Price, error)

	EndAt(ctx context.Context, month time.Month, day int, priceId int) 					  			  (*ent.Price, error)
	Wipe(ctx context.Context)
}

type priceMySQL struct {
	client *ent.Client
}

func (mysql priceMySQL) GetPriceByPlanIdAndTime(ctx context.Context, planId int, now time.Time, withProduct bool, withPlan bool, withThirdParty bool) (*ent.Price, error) {
	query := mysql.client.Price.Query().Where(
		price.And(
			price.StartAtLTE(now),
			price.EndAtGTE(now),
			price.HasPlansWith(plan.ID(planId))),
	)
	if withThirdParty {
		query.WithThirdParties()
	}
	if withPlan {
		query.WithPlans()
	}
	if withProduct {
		query.WithProducts()
	}
	plan, err := query.Only(ctx)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (mysql priceMySQL) GetPriceByPlanId(ctx context.Context, planId int, withProduct bool, withPlan bool, withThirdParty bool) (*ent.Price, error) {
	return mysql.GetPriceByPlanIdAndTime(ctx, planId, time.Now().UTC(), withProduct, withPlan, withThirdParty)
}

func (mysql priceMySQL) GetPricesByPlanId(ctx context.Context, planId int) ([]*ent.Price, error) {
	plans, err := mysql.client.Price.Query().Where(price.HasPlansWith(plan.ID(planId))).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func (mysql priceMySQL) GetPriceByProductID(ctx context.Context, productId int, withProduct bool, withPlan bool, withThirdParty bool) (*ent.Price, error) {
	now := time.Now().UTC()
	return mysql.GetPriceByProductIDAndTime(ctx, productId, now, withProduct, withPlan, withThirdParty)
}

func (mysql priceMySQL) GetPriceByProductIDAndTime(ctx context.Context, productId int, now time.Time, withProduct bool, withPlan bool, withThirdParty bool) (*ent.Price, error) {
	query := mysql.client.Price.Query().Where(
		price.And(
			price.StartAtLTE(now),
			price.EndAtGTE(now),
			price.HasProductsWith(product.ID(productId)),
		))
	if withProduct {
		query.WithProducts()
	}
	if withThirdParty {
		query.WithThirdParties()
	}
	if withPlan {
		query.WithPlans()
	}
	only, err := query.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrPriceNotExist
		}
		return nil, err
	}
	return only, nil
}

//1   |    Jan 1, 1970, 00:00:00 |  Dec 20, 2011, 00:00:00  |   10$ |   10$
//1   |   Dec 20, 2011, 00:00:01 |  Dec 26, 2011, 00:00:00  |  	10$ |   10$
func (mysql priceMySQL) CreatePrice(ctx context.Context, price uint16, discount uint16, productId int, startAt time.Time, endAt time.Time, thirdPartyID *int, plansId *int) (*ent.Price, error) {
	if endAt.Before(startAt) {
		return nil, ErrInvalidEndAt
	}
	if startAt.After(endAt) {
		return nil, ErrInvalidStartAt
	}
	prices, err := mysql.GetPricesByProductID(ctx, productId, false)
	if err != nil {
		return nil, err
	}

	if plansId != nil {
		if *plansId == 0 {
			plansId = nil
		}
	}
	if thirdPartyID != nil {
		if *thirdPartyID == 0 {
			thirdPartyID = nil
		}
	}
	build := mysql.
		client.
		Price.
		Create().
		SetPrice(price).
		SetDiscount(discount).
		SetProductsID(productId).
		SetNillablePlansID(plansId).
		SetNillableThirdPartiesID(thirdPartyID)
	// check collision with month and day with existing price
	if plansId != nil {
		// subscription plans..
		subscriptionPlans, err := mysql.GetPricesByPlanId(ctx, *plansId)
		if err != nil {
			return nil, err
		}
		if len(subscriptionPlans) > 0 {
			index := len(subscriptionPlans) - 1
			latestEndAt := subscriptionPlans[index].EndAt
			if startAt.Unix() < latestEndAt.Unix() {
				return nil, ErrInvalidStartAt
			}
			build.SetStartAt(startAt.Add(1 * time.Second).UTC()).SetEndAt(endAt.UTC())
		} else {
			build.SetStartAt(startAt.Add(1 * time.Second).UTC()).SetEndAt(endAt.UTC())
		}
	} else if len(prices) > 0 {
		// other products..
		index := len(prices) - 1
		latestEndAt := prices[index].EndAt
		if startAt.Unix() < latestEndAt.Unix() {
			return nil, ErrInvalidStartAt
		}
		build.SetStartAt(startAt.Add(1 * time.Second).UTC()).SetEndAt(endAt.UTC())
	} else {
		// initial..
		build.SetStartAt(startAt).SetEndAt(endAt)
	}
	created, err := build.Save(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.GetPriceByID(ctx, created.ID, true, true, true)
}


func (mysql priceMySQL) GetPricesByProductID(ctx context.Context, productId int, isSequence bool) ([]*ent.Price, error) {
	now := time.Now().UTC()
	query := mysql.client.Price.Query().Order(ent.Asc(price.FieldID))
	if isSequence {
		query.Where(price.And(
			price.StartAtLTE(now),
			price.EndAtGTE(now),
			price.HasProductsWith(product.ID(productId)),
		)).WithPlans().
			WithThirdParties().
			WithProducts()
	} else {
		query.Where(price.HasProductsWith(product.ID(productId)))
	}
	all, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return all, err
}

func (mysql priceMySQL) GetPriceByID(ctx context.Context, id int, withProduct bool, withPlan bool, withThirdParty bool) (*ent.Price, error) {
	query := mysql.client.Price.Query().Where(price.ID(id))
	if withProduct {
		query.WithProducts()
	}
	if withPlan {
		query.WithPlans()
	}
	if withThirdParty {
		query.WithThirdParties()
	}
	only, err := query.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrPriceNotExist
		}
		return nil, err
	}
	return only, nil
}

func (mysql priceMySQL) EndAt(ctx context.Context, month time.Month, day int, priceId int) (*ent.Price, error) {
	prices, err := mysql.GetPriceByID(ctx, priceId, false, false, false)
	if err != nil {
		return nil, err
	}
	entAt, err := clock.ToMonthDate(month, day)
	if err != nil {
		return nil, err
	}
	updated, err := prices.Update().SetEndAt(entAt.UTC()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.GetPriceByID(ctx,updated.ID, true, true,true)
}

func (mysql priceMySQL) Wipe(ctx context.Context) {
	mysql.client.Subscription.Delete().ExecX(ctx)
	mysql.client.ThirdParty.Delete().ExecX(ctx)
	mysql.client.Price.Delete().ExecX(ctx)
	mysql.client.Plan.Delete().ExecX(ctx)
	mysql.client.Product.Delete().ExecX(ctx)
}

func (mysql priceMySQL) CreateNewPrice(ctx context.Context, price uint16, discount uint16, productId int, iapSourceId *int, plansId *int) (*ent.Price, error) {
	return mysql.CreatePrice(ctx, price, discount, productId, clock.ZeroAt().UTC(), clock.InfinityAt().UTC(), iapSourceId, plansId)
}

func NewPriceRepository(client *ent.Client) PriceRepository {
	return &priceMySQL{client: client}
}
