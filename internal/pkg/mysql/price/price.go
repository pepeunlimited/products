package price

import (
	"context"
	"errors"
	"github.com/pepeunlimited/products/internal/pkg/clock"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/internal/pkg/ent/price"
	"github.com/pepeunlimited/products/internal/pkg/ent/product"
	"time"
)

var (
	ErrPriceNotExist 		= errors.New("prices: not exist")
	ErrInvalidStartAt 		= errors.New("prices: invalid startAt")
	ErrInvalidEndAt 		= errors.New("prices: invalid endAt")
	ErrInvalidStartAtEndAt 	= errors.New("price: startAt and endAt are equal")
	ErrInvalidProduct 		= errors.New("prices: invalid product")
)

type PriceRepository interface {
	GetPriceByProductSku(ctx context.Context, productSku string, withProduct bool, withThirdParty bool) (*ent.Price, error)
	GetPriceByProductSkuAndTime(ctx context.Context, productSku string, now time.Time, withProduct bool, withThirdParty bool) (*ent.Price, error)

	GetPriceByProductID(ctx context.Context, productId int, withProduct bool, withThirdParty bool) (*ent.Price, error)
	GetPriceByProductIDAndTime(ctx context.Context, productId int, now time.Time, withProduct bool, withThirdParty bool) (*ent.Price, error)
	GetPricesByProductID(ctx context.Context, productId int, isSequence bool) ([]*ent.Price, error)
	GetPriceByID(ctx context.Context, id int, withProduct bool, withThirdParty bool) (*ent.Price, error)
	CreateNewPrice(ctx context.Context, price uint16, discount uint16, productId int, thirdPartyID *int) (*ent.Price, error)
	CreatePrice(ctx context.Context, price uint16, discount uint16, productId int, startAt time.Time, endAt time.Time, thirdPartyID *int) (*ent.Price, error)

	EndAt(ctx context.Context, month time.Month, day int, priceId int) (*ent.Price, error)
	Wipe(ctx context.Context)
}

type priceMySQL struct {
	client *ent.Client
}

func (mysql priceMySQL) GetPriceByProductSku(ctx context.Context, productSku string, withProduct bool, withThirdParty bool) (*ent.Price, error) {
	now := time.Now().UTC()
	return mysql.GetPriceByProductSkuAndTime(ctx, productSku, now, withProduct, withThirdParty)
}

func (mysql priceMySQL) GetPriceByProductSkuAndTime(ctx context.Context, productSku string, now time.Time, withProduct bool, withThirdParty bool) (*ent.Price, error) {
	query := mysql.client.Price.Query().Where(
		price.And(
			price.StartAtLTE(now),
			price.EndAtGTE(now),
			price.HasProductsWith(product.Sku(productSku)),
		))
	if withProduct {
		query.WithProducts()
	}
	if withThirdParty {
		query.WithThirdPartyPrices()
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

func (mysql priceMySQL) GetPriceByProductID(ctx context.Context, productId int, withProduct bool, withThirdParty bool) (*ent.Price, error) {
	now := time.Now().UTC()
	return mysql.GetPriceByProductIDAndTime(ctx, productId, now, withProduct, withThirdParty)
}

func (mysql priceMySQL) GetPriceByProductIDAndTime(ctx context.Context, productId int, now time.Time, withProduct bool, withThirdParty bool) (*ent.Price, error) {
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
		query.WithThirdPartyPrices()
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
func (mysql priceMySQL) CreatePrice(ctx context.Context, price uint16, discount uint16, productId int, startAt time.Time, endAt time.Time, thirdPartyID *int) (*ent.Price, error) {
	if startAt.Equal(endAt) {
		return nil, ErrInvalidStartAtEndAt
	}
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
		SetNillableThirdPartyPricesID(thirdPartyID)
	// check collision with month and day with existing price
	if len(prices) > 0 {
		// other products..
		index := len(prices) - 1
		latestEndAt := prices[index].EndAt
		if startAt.Unix() < latestEndAt.Unix() {
			return nil, ErrInvalidStartAt
		}
		build.SetStartAt(startAt.Add(1 * time.Second).UTC()).SetEndAt(endAt.UTC())
	} else {
		if startAt.Year() != 1970 {
			startAt = startAt.Add(1*time.Second)
		}
		// initial..
		build.SetStartAt(startAt).SetEndAt(endAt)
	}
	created, err := build.Save(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.GetPriceByID(ctx, created.ID, true, true)
}

func (mysql priceMySQL) GetPricesByProductID(ctx context.Context, productId int, isSequence bool) ([]*ent.Price, error) {
	now := time.Now().UTC()
	query := mysql.client.Price.Query().Order(ent.Asc(price.FieldID))
	if isSequence {
		query.Where(price.And(
			price.StartAtLTE(now),
			price.EndAtGTE(now),
			price.HasProductsWith(product.ID(productId)),
		)).
			WithThirdPartyPrices().
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

func (mysql priceMySQL) GetPriceByID(ctx context.Context, id int, withProduct bool, withThirdParty bool) (*ent.Price, error) {
	query := mysql.client.Price.Query().Where(price.ID(id))
	if withProduct {
		query.WithProducts()
	}
	if withThirdParty {
		query.WithThirdPartyPrices()
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
	prices, err := mysql.GetPriceByID(ctx, priceId, false, false)
	if err != nil {
		return nil, err
	}
	entAt, err := clock.ToMonthDate(month, day)
	if prices.StartAt.Add(-1 * time.Second).Equal(entAt) { // do not allow same date
		return nil, ErrInvalidEndAt
	}
	if err != nil {
		return nil, err
	}
	updated, err := prices.Update().SetEndAt(entAt.UTC()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.GetPriceByID(ctx,updated.ID, true, true)
}

func (mysql priceMySQL) Wipe(ctx context.Context) {
	mysql.client.Subscription.Delete().ExecX(ctx)
	mysql.client.Price.Delete().ExecX(ctx)
	mysql.client.Plan.Delete().ExecX(ctx)
	mysql.client.ThirdPartyPrice.Delete().ExecX(ctx)
	mysql.client.Product.Delete().ExecX(ctx)
}

func (mysql priceMySQL) CreateNewPrice(ctx context.Context, price uint16, discount uint16, productId int, iapSourceId *int) (*ent.Price, error) {
	return mysql.CreatePrice(ctx, price, discount, productId, clock.ZeroAt().UTC(), clock.InfinityAt().UTC(), iapSourceId)
}

func New(client *ent.Client) PriceRepository {
	return &priceMySQL{client: client}
}
