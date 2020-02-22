package twirp

import (
	"context"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/products/internal/pkg/clock"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	thirdpartypricerepo "github.com/pepeunlimited/products/internal/pkg/mysql/thirdpartyprice"
	"github.com/pepeunlimited/products/internal/server/errorz"
	"github.com/pepeunlimited/products/internal/server/validator"
	"github.com/pepeunlimited/products/pkg/rpc/thirdpartyprice"

	"time"
)

type ThirdPartyPriceServer struct {
	thirdpartyprice thirdpartypricerepo.ThirdPartyPriceRepository
	valid validator.ThirdPartyServerValidator
}

func (server ThirdPartyPriceServer) CreateThirdPartyPrice(ctx context.Context, params *thirdpartyprice.CreateThirdPartyPriceParams) (*thirdpartyprice.ThirdPartyPrice, error) {
	err := server.valid.CreateThirdParty(params)
	if err != nil {
		return nil, err
	}
	types := thirdpartypricerepo.ThirdPartyPriceTypeFromString(params.Type)
	if params.StartAtDay != 0 && params.StartAtMonth != 0 { // without start day
		startAt, err := clock.ToMonthDate(time.Month(params.StartAtMonth), int(params.StartAtDay))
		if err != nil {
			return nil, err
		}
		startAt = startAt.Add(1 * time.Second)
		thirdparty, err := server.thirdpartyprice.CreateStartAt(ctx, params.InAppPurchaseSku, &params.GoogleBillingServiceSku, startAt, types)
		if err != nil {
			return nil, errorz.ThirdParty(err)
		}
		return ToThirdParty(thirdparty), nil
	}
	thirdparty, err := server.thirdpartyprice.Create(ctx, params.InAppPurchaseSku, &params.GoogleBillingServiceSku, types)
	if err != nil {
		return nil, errorz.ThirdParty(err)
	}
	return ToThirdParty(thirdparty), nil
}

func (server ThirdPartyPriceServer) GetThirdPartyPrices(ctx context.Context, params *thirdpartyprice.GetThirdPartiesParams) (*thirdpartyprice.GetThirdPartyPricesResponse, error) {
	err := server.valid.GetThirdParties(params)
	if err != nil {
		return nil, err
	}
	thirdParties, err := server.thirdpartyprice.GetThirdPartyPrices(ctx)
	if err != nil {
		return nil, errorz.ThirdParty(err)
	}
	return &thirdpartyprice.GetThirdPartyPricesResponse{ThirdPartyPrices: ToThirdParties(thirdParties)}, nil
}

func (server ThirdPartyPriceServer) GetThirdPartyPrice(ctx context.Context, params *thirdpartyprice.GetThirdPartyPriceParams) (*thirdpartyprice.ThirdPartyPrice, error) {
	err := server.valid.GetThirdParty(params)
	if err != nil {
		return nil, err
	}
	var price *ent.ThirdPartyPrice
	if !validator2.IsEmpty(params.InAppPurchaseSku) {
		price, err = server.thirdpartyprice.GetInAppPurchaseBySku(ctx, params.InAppPurchaseSku)
	}
	if !validator2.IsEmpty(params.GoogleBillingServiceSku) {
		price, err = server.thirdpartyprice.GetBillingBySku(ctx, params.GoogleBillingServiceSku)
	}
	if params.ThirdPartyPriceId != 0 {
		price, err = server.thirdpartyprice.GetByID(ctx, int(params.ThirdPartyPriceId))
	}
	if err != nil {
		return nil, errorz.ThirdParty(err)
	}
	return ToThirdParty(price), nil
}

func (server ThirdPartyPriceServer) EndThirdParty(ctx context.Context, params *thirdpartyprice.EndThirdPartyPriceParams) (*thirdpartyprice.ThirdPartyPrice, error) {
	err := server.valid.EndThirdParty(params)
	if err != nil {
		return nil, err
	}
	party, err := server.GetThirdPartyPrice(ctx, params.Params)
	if err != nil {
		return nil, err
	}
	ended, err := server.thirdpartyprice.EndAt(ctx, time.Month(params.EndAtMonth), int(params.EndAtDay), int(party.Id))
	if err != nil {
		return nil, errorz.ThirdParty(err)
	}
	return ToThirdParty(ended), nil
}

func NewThirdPartyPriceServer(client *ent.Client) ThirdPartyPriceServer {
	return ThirdPartyPriceServer{
		thirdpartyprice:thirdpartypricerepo.New(client),
		valid:validator.NewThirdPartyServerValidator(),
	}
}