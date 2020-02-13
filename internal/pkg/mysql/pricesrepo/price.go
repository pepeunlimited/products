package pricesrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/ent/price"
	"math"
	"time"
)

var (
	ErrInvalidDay 		= errors.New("prices: invalid day")
	ErrInvalidMonth 	= errors.New("prices: invalid month")
	ErrPriceSkuExist   	= errors.New("prices: sku exist")
	ErrPriceNotExist  	= errors.New("prices: not exist")
)

type PriceRepository interface {
	GetPriceByID(ctx context.Context, id int) 							  (*ent.Price, error)
	GetPriceBySku(ctx context.Context, sku string) 						  (*ent.Price, error)

	CreateInitialPrice(ctx context.Context, price uint16, sku *string)    (*ent.Price, error)
	EndPrice(ctx context.Context, month time.Month, day int, priceId int) (*ent.Price, error)
	Wipe(ctx context.Context)
}

type priceMySQL struct {
	client *ent.Client
}

func (mysql priceMySQL) GetPriceByID(ctx context.Context, id int) (*ent.Price, error) {
	selected, err := mysql.client.Price.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrPriceNotExist
		}
		return nil, err
	}
	return selected, nil
}

func (mysql priceMySQL) GetPriceBySku(ctx context.Context, sku string) (*ent.Price, error) {
	selected, err := mysql.client.Price.Query().Where(price.Sku(sku)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrPriceNotExist
		}
		return nil, err
	}
	return selected, nil
}

func (mysql priceMySQL) EndPrice(ctx context.Context, month time.Month, day int, priceId int) (*ent.Price, error) {
	prices, err := mysql.GetPriceByID(ctx, priceId)
	if err != nil {
		return nil, err
	}
	entAt, err := mysql.toMonthDate(month, day)
	if err != nil {
		return nil, err
	}
	updated, err := prices.Update().SetEndAt(entAt).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (mysql priceMySQL) Wipe(ctx context.Context) {
	mysql.client.Price.Delete().ExecX(ctx)
}

func (mysql priceMySQL) CreateInitialPrice(ctx context.Context, price uint16, sku *string) (*ent.Price, error) {
	prices, err := mysql.
		client.
		Price.
		Create().
		SetPrice(price).
		SetDiscount(price).
		SetNillableSku(sku).
		SetStartAt(mysql.zeroAt()).
		SetEndAt(mysql.infinityAt()).Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, ErrPriceSkuExist
		}
		return nil, err
	}
	return prices, nil
}

func NewPriceRepository(client *ent.Client) PriceRepository {
	return &priceMySQL{client: client}
}

// 2106-02-07 06:28:15.000
func (mysql priceMySQL) infinityAt() time.Time {
	return time.Unix(math.MaxUint32, 0)
}

// 1970-01-01 00:00:00.000
func (mysql priceMySQL) zeroAt() time.Time {
	return time.Unix(0, 0)
}

func (mysql priceMySQL) toMonthDate(month time.Month, day int) (time.Time, error) {
	if day <= 0 || day > 31 {
		return time.Time{}, ErrInvalidDay
	}
	if month <= 0 || month > 12 {
		return time.Time{}, ErrInvalidMonth
	}
	current := time.Now().UTC()
	return time.Date(current.Year(), month, day, 0, 0,0,0, time.UTC), nil
}