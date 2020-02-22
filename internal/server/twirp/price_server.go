package twirp

import (
	"context"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/plan"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/price"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/product"
	"github.com/pepeunlimited/prices/internal/server/errorz"
	"github.com/pepeunlimited/prices/internal/server/validator"
	"github.com/pepeunlimited/prices/pkg/rpc/price"
	"github.com/twitchtv/twirp"
	"time"
)

type PriceServer struct {
	prices       price.PriceRepository
	valid        validator.PriceServerValidator
	products     product.ProductRepository
	thirdparties thirdpartyrepo.ThirdPartyRepository
	plans        plan.PlanRepository
}

func (server PriceServer) EndPrice(ctx context.Context, params *price.EndPriceParams) (*price.Price, error) {
	err := server.valid.EndPrice(params)
	if err != nil {
		return nil, err
	}
	price, err := server.GetPrice(ctx, params.Params)
	if err != nil {
		return nil, err
	}
	at, err := server.prices.EndAt(ctx, time.Month(params.EndAtMonth), int(params.EndAtDay), int(price.Id))
	return ToPrice(at), nil
}

func (server PriceServer) GetSubscriptionPrices(ctx context.Context, params *price.GetSubscriptionPricesParams) (*price.GetSubscriptionPricesResponse, error) {
	err := server.valid.GetSubscriptionPrices(params)
	if err != nil {
		return nil, err
	}
	var prices []*ent.Price
	if params.ProductId != 0 { // ByProductId
		_, isSubscription, err2 := server.isProductSubscribableById(ctx, int(params.ProductId))
		if err2 != nil {
			return nil, errorz.Product(err)
		}
		if !*isSubscription {
			return nil, twirp.InvalidArgumentError("product_id", "price_is_not_subscription")
		}
		prices, err = server.prices.GetPricesByProductID(ctx, int(params.ProductId), true)
	}
	if !validator2.IsEmpty(params.ProductSku) { // ByProductSku
		product, isSubscription, err := server.isProductSubscribableBySku(ctx, params.ProductSku)
		if err != nil {
			return nil, errorz.Product(err)
		}
		if !*isSubscription {
			return nil, twirp.InvalidArgumentError("product_sku", "price_is_not_subscription")
		}
		prices, err = server.prices.GetPricesByProductID(ctx, product.ID, true)
	}
	if err != nil {
		return nil, errorz.Price(err)
	}
	return &price.GetSubscriptionPricesResponse{Prices: ToPrices(prices)}, nil
}

func (server PriceServer) CreatePrice(ctx context.Context, params *price.CreatePriceParams) (*price.Price, error) {
	err := server.valid.CreatePrice(params)
	if err != nil {
		return nil, err
	}
	_, err = server.products.GetProductByID(ctx, false, int(params.ProductId))
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
	plansId 	  := int(params.PlanId)
	thirdPartyId  := int(params.ThirdPartyId)
	if plansId != 0 {
		_, err := server.plans.GetPlanByID(ctx, plansId)
		if err != nil {
			return nil, errorz.Plan(err)
		}
	}
	if thirdPartyId != 0 {
		_, err := server.thirdparties.GetByID(ctx, thirdPartyId)
		if err != nil {
			return nil, errorz.ThirdParty(err)
		}
	}
	price, err := server.prices.CreatePrice(ctx, uint16(params.Price), uint16(params.Discount), int(params.ProductId), startAt, endAt, &thirdPartyId, &plansId)
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
		_, isSubscription, err := server.isProductSubscribableById(ctx, int(params.ProductId))
		if err != nil {
			return nil, errorz.Product(err)
		}
		if *isSubscription {
			return nil, twirp.InvalidArgumentError("product_id", "price_is_subscription")
		}
		price, err = server.prices.GetPriceByProductID(ctx, int(params.ProductId), true, true, true)
	}
	if !validator2.IsEmpty(params.ProductSku) { // ByProductSku
		product, isSubscription, err := server.isProductSubscribableBySku(ctx, params.ProductSku)
		if err != nil {
			return nil, errorz.Product(err)
		}
		if *isSubscription {
			return nil, twirp.InvalidArgumentError("product_sku", "price_is_subscription")
		}
		price, err = server.prices.GetPriceByProductID(ctx, product.ID, true, true,true)
	}
	if params.PriceId != 0 { // ByPriceID
		price, err = server.prices.GetPriceByID(ctx, int(params.PriceId), true, true,true)
	}
	if params.PlanId != 0 { // ByPlanId
		price, err = server.prices.GetPriceByPlanId(ctx, int(params.PlanId), true, true, true)
	}
	if nil != err {
		return nil, errorz.Price(err)
	}
	return ToPrice(price), nil
}

func (server PriceServer) isProductSubscribableBySku(ctx context.Context, productSku string) (*ent.Product, *bool, error) {
	sku, err := server.products.GetProductBySku(ctx, productSku)
	if err != nil {
		return nil, nil, errorz.Product(err)
	}
	return server.isProductSubscribableById(ctx, sku.ID)
}

func (server PriceServer) isProductSubscribableById(ctx context.Context, productId int) (*ent.Product, *bool, error) {
	product, isSubscribable, err := server.products.IsSubscribableByID(ctx, productId)
	if err != nil {
		return nil, nil, errorz.Product(err)
	}
	return product, isSubscribable, nil
}

func NewPriceServer(client *ent.Client) PriceServer {
	return PriceServer{
		prices:       price.New(client),
		valid:        validator.NewPriceServerValidator(),
		products:     product.New(client),
		thirdparties: thirdpartyrepo.NewThirdPartyRepository(client),
		plans:        plan.NewPlanRepository(client),
	}
}