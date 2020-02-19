package twirp

import (
	"context"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/planrepo"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/pricerepo"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/productrepo"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/thirdpartyrepo"
	"github.com/pepeunlimited/prices/internal/server/errorz"
	"github.com/pepeunlimited/prices/internal/server/validator"
	"github.com/pepeunlimited/prices/pkg/pricerpc"
	"github.com/twitchtv/twirp"
	"time"
)

type PriceServer struct {
	prices 				pricerepo.PriceRepository
	valid 				validator.PriceServerValidator
	products			productrepo.ProductRepository
	thirdparties    	thirdpartyrepo.ThirdPartyRepository
	plans				planrepo.PlanRepository
	productErrorz 		errorz.ProductErrorz
	priceErrorz 		errorz.PriceErrorz
	thirdPartyErrorz	errorz.ThirdPartyErrorz
	planErrorz			errorz.PlanErrorz
}

func (server PriceServer) EndPrice(ctx context.Context, params *pricerpc.EndPriceParams) (*pricerpc.Price, error) {
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

func (server PriceServer) GetSubscriptionPrices(ctx context.Context, params *pricerpc.GetSubscriptionPricesParams) (*pricerpc.GetSubscriptionPricesResponse, error) {
	panic("implement me")
}

func (server PriceServer) CreatePrice(ctx context.Context, params *pricerpc.CreatePriceParams) (*pricerpc.Price, error) {
	err := server.valid.CreatePrice(params)
	if err != nil {
		return nil, err
	}
	_, err = server.products.GetProductByID(ctx, false, int(params.ProductId))
	if err != nil {
		return nil, server.productErrorz.IsProductError(err)
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
		_, err := server.plans.GetPlansByID(ctx, plansId)
		if err != nil {
			return nil, server.planErrorz.IsPlanError(err)
		}
	}
	if thirdPartyId != 0 {
		_, err := server.thirdparties.GetByID(ctx, thirdPartyId)
		if err != nil {
			return nil, server.thirdPartyErrorz.IsThirdPartyError(err)
		}
	}
	price, err := server.prices.CreatePrice(ctx, uint16(params.Price), uint16(params.Discount), int(params.ProductId), startAt, endAt, &thirdPartyId, &plansId)
	if err != nil {
		return nil, server.priceErrorz.IsPriceError(err)
	}
	return ToPrice(price), nil
}

func (server PriceServer) GetPrice(ctx context.Context, params *pricerpc.GetPriceParams) (*pricerpc.Price, error) {
	err := server.valid.GetPrice(params)
	if err != nil {
		return nil, err
	}
	var price *ent.Price
	if params.ProductId != 0 {
		err = server.isSubscribableByProductId(ctx, int(params.ProductId))
		if err != nil {
			return nil, err
		}
		price, err = server.prices.GetPriceByProductID(ctx, int(params.ProductId), true, true, true)
	}
	if !validator2.IsEmpty(params.ProductSku) {
		product, err := server.isSubscribableByProductSku(ctx, params.ProductSku)
		if err != nil {
			return nil, err
		}
		price, err = server.prices.GetPriceByProductID(ctx, product.ID, true, true,true)
	}
	if params.PriceId != 0 {
		price, err = server.prices.GetPriceByID(ctx, int(params.PriceId), true, true,true)
	}
	if nil != err {
		return nil, server.priceErrorz.IsPriceError(err)
	}
	return ToPrice(price), nil
}

func (server PriceServer) isSubscribableByProductSku(ctx context.Context, productSku string) (*ent.Product, error) {
	sku, err := server.products.GetProductBySku(ctx, productSku)
	if err != nil {
		return nil, server.productErrorz.IsProductError(err)
	}
	return sku, server.isSubscribableByProductId(ctx, sku.ID)
}

func (server PriceServer) isSubscribableByProductId(ctx context.Context, productId int) error {
	isSubscribable, err := server.products.IsSubscribableByID(ctx, productId)
	if err != nil {
		return server.productErrorz.IsProductError(err)
	}
	if *isSubscribable {
		return twirp.InvalidArgumentError("product_id", "is_subscribable")
	}
	return nil
}

func NewPriceServer(client *ent.Client) PriceServer {
	return PriceServer{
		prices:pricerepo.NewPriceRepository(client),
		valid:validator.NewPriceServerValidator(),
		products:productrepo.NewProductRepository(client),
		priceErrorz: errorz.NewPriceErrorz(),
		thirdparties:thirdpartyrepo.NewThirdPartyRepository(client),
		productErrorz:errorz.NewProductErrorz(),
		plans:planrepo.NewPlanRepository(client),
		thirdPartyErrorz:errorz.NewThirdPartyErrorz(),
		planErrorz:errorz.NewPlanErrorz(),
	}
}