package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/pricerepo"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/productrepo"
	"github.com/pepeunlimited/prices/internal/server/errorz"
	"github.com/pepeunlimited/prices/internal/server/validator"
	"github.com/pepeunlimited/prices/pkg/pricerpc"
	"time"
)

type PriceServer struct {
	prices 			pricerepo.PriceRepository
	valid 			validator.PriceServerValidator
	products		productrepo.ProductRepository
	productErrorz 	errorz.ProductErrorz
	priceErrorz 	errorz.PriceErrorz
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
	price, err := server.prices.CreatePrice(ctx, uint16(params.Price), uint16(params.Discount), int(params.ProductId), startAt, endAt, nil, nil)
	if err != nil {
		return nil, server.priceErrorz.IsPriceError(err)
	}
	return ToPrice(price), nil
}

func (server PriceServer) GetPrice(ctx context.Context, params *pricerpc.GetPriceParams) (*pricerpc.Price, error) {
	panic("implement me")
}

func NewPriceServer(client *ent.Client) PriceServer {
	return PriceServer{
		prices:pricerepo.NewPriceRepository(client),
		valid:validator.NewPriceServerValidator(),
		products:productrepo.NewProductRepository(client),
		priceErrorz: errorz.NewPriceErrorz(),
		productErrorz:errorz.NewProductErrorz(),
	}
}