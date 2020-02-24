package twirp

import (
	"context"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/products/internal/pkg/clock"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/internal/pkg/mysql/plan"
	pricerepo "github.com/pepeunlimited/products/internal/pkg/mysql/price"
	"github.com/pepeunlimited/products/internal/pkg/mysql/product"
	"github.com/pepeunlimited/products/internal/pkg/mysql/thirdpartyprice"
	"github.com/pepeunlimited/products/internal/server/errorz"
	"github.com/pepeunlimited/products/internal/server/validator"
	"github.com/pepeunlimited/products/pkg/rpc/price"
	"time"
)

type PriceServer struct {
	prices       pricerepo.PriceRepository
	valid        validator.PriceServerValidator
	products     product.ProductRepository
	thirdparties thirdpartyprice.ThirdPartyPriceRepository
	plans        plan.PlanRepository
}

func (server PriceServer) EndPriceAt(ctx context.Context, params *price.EndPriceAtParams) (*price.Price, error) {
	err := server.valid.EndPrice(params)
	if err != nil {
		return nil, err
	}
	price, err := server.GetPrice(ctx, params.Params)
	if err != nil {
		return nil, err
	}
	at, err := server.prices.EndAt(ctx, time.Month(params.EndAtMonth), int(params.EndAtDay), int(price.Id))
	if err != nil {
		return nil, errorz.Price(err)
	}
	return ToPrice(at), nil
}

func (server PriceServer) CreatePrice(ctx context.Context, params *price.CreatePriceParams) (*price.Price, error) {
	err := server.valid.CreatePrice(params)
	if err != nil {
		return nil, err
	}
	_, err = server.products.GetProductByID(ctx, false, false, int(params.ProductId))
	if err != nil {
		return nil, errorz.Product(err)
	}
	var startAt time.Time
	var endAt 	time.Time
	if params.Discount == 0 {
		params.Discount = params.Price
	}
	if params.StartAtMonth != 0 && params.StartAtDay != 0 { //StartAt
		startAt, err = clock.ToMonthDate(time.Month(params.StartAtMonth), int(params.StartAtDay))
		if err != nil {
			return nil, err
		}
	} else {
		startAt = clock.ZeroAt()
	}
	if params.EndAtMonth != 0 && params.EndAtDay != 0 {//EndAt
		endAt, err = clock.ToMonthDate(time.Month(params.EndAtMonth), int(params.EndAtDay))
		if err != nil {
			return nil, err
		}
	} else {
		endAt = clock.InfinityAt()
	}
	thirdPartyId  := int(params.ThirdPartyId)
	if thirdPartyId != 0 {
		_, err := server.thirdparties.GetByID(ctx, thirdPartyId)
		if err != nil {
			return nil, errorz.ThirdParty(err)
		}
	}
	price, err := server.prices.CreatePrice(ctx, params.Price, params.Discount, int(params.ProductId), startAt, endAt, &thirdPartyId)
	if err != nil {
		return nil, errorz.Price(err)
	}
	return ToPrice(price), nil
}

func (server PriceServer) GetPrice(ctx context.Context, params *price.GetPriceParams) (*price.Price, error) {
	err := server.valid.GetPrice(params)
	if err != nil {
		return nil, err
	}
	var price *ent.Price
	if params.ProductId != 0 { // ByProductID
		price, err = server.prices.GetPriceByProductID(ctx, int(params.ProductId), true, true)
	}
	if !validator2.IsEmpty(params.ProductSku) { // ByProductSku
		price, err = server.prices.GetPriceByProductSku(ctx, params.ProductSku, true, true)
	}
	if params.PriceId != 0 { // ByPriceID
		price, err = server.prices.GetPriceByID(ctx, int(params.PriceId), true, true)
	}
	if nil != err {
		return nil, errorz.Price(err)
	}
	return ToPrice(price), nil
}

func NewPriceServer(client *ent.Client) PriceServer {
	return PriceServer{
		prices:       pricerepo.New(client),
		valid:        validator.NewPriceServerValidator(),
		products:     product.New(client),
		thirdparties: thirdpartyprice.New(client),
		plans:        plan.New(client),
	}
}