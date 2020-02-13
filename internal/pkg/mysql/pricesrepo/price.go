package pricesrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"log"
	"time"
)

var (
	ErrPriceNotExist  					= errors.New("prices: not exist")
)

type PriceRepository interface {
	GetPriceByID(ctx context.Context, id int) 							  					  			  (*ent.Price, error)
	CreateInitialPrice(ctx context.Context, price uint16, productId int, iapSourceId *int, plansId *int)  (*ent.Price, error)
	EndPrice(ctx context.Context, month time.Month, day int, priceId int) 					  			  (*ent.Price, error)
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

func (mysql priceMySQL) EndPrice(ctx context.Context, month time.Month, day int, priceId int) (*ent.Price, error) {
	prices, err := mysql.GetPriceByID(ctx, priceId)
	if err != nil {
		return nil, err
	}
	entAt, err := clock.ToMonthDate(month, day)
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
	mysql.client.Product.Delete().ExecX(ctx)
}

func (mysql priceMySQL) CreateInitialPrice(ctx context.Context, price uint16, productId int, iapSourceId *int, plansId *int) (*ent.Price, error) {
	//TODO: check if already exist
	prices, err := mysql.
		client.
		Price.
		Create().
		SetPrice(price).
		SetDiscount(price).
		SetProductsID(productId).
		SetStartAt(clock.ZeroAt()).
		SetEndAt(clock.InfinityAt()).
		SetNillablePlansID(plansId).
		SetNillableIapSourceID(iapSourceId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return prices, nil
}

func NewPriceRepository(client *ent.Client) PriceRepository {
	return &priceMySQL{client: client}
}

func (mysql priceMySQL) rollback(tx *ent.Tx){
	if err := tx.Rollback(); err != nil {
		log.Print("prices: failed execute rollback.."+err.Error())
	}
}

func (mysql priceMySQL) commit(tx *ent.Tx) error {
	if err := tx.Commit(); err != nil {
		log.Print("prices: failed execute commit.."+err.Error())
		return err
	}
	return nil
}