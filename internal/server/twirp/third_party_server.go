package twirp

import (
	"context"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/thirdpartyrepo"
	"github.com/pepeunlimited/prices/internal/server/errorz"
	"github.com/pepeunlimited/prices/internal/server/validator"
	"github.com/pepeunlimited/prices/pkg/thirdpartyrpc"
	"time"
)

type ThirdPartyServer struct {
	thirdparty thirdpartyrepo.ThirdPartyRepository
	valid validator.ThirdPartyServerValidator
	errorz errorz.ThirdPartyErrorz
}

func (server ThirdPartyServer) EndThirdParty(ctx context.Context, params *thirdpartyrpc.EndThirdPartyParams) (*thirdpartyrpc.ThirdParty, error) {
	err := server.valid.EndThirdParty(params)
	if err != nil {
		return nil, err
	}
	party, err := server.GetThirdParty(ctx, params.Params)
	if err != nil {
		return nil, err
	}
	ended, err := server.thirdparty.EndAt(ctx, time.Month(params.EndAtMonth), int(params.EndAtDay), int(party.Id))
	if err != nil {
		return nil, server.errorz.IsThirdPartyError(err)
	}
	return ToThirdParty(ended), nil
}

func (server ThirdPartyServer) CreateThirdParty(ctx context.Context, params *thirdpartyrpc.CreateThirdPartyParams) (*thirdpartyrpc.ThirdParty, error) {
	err := server.valid.CreateThirdParty(params)

	if err != nil {
		return nil, err
	}
	if params.StartAtDay != 0 && params.StartAtMonth != 0 { // without start day
		startAt, err := clock.ToMonthDate(time.Month(params.StartAtMonth), int(params.StartAtDay))
		if err != nil {
			return nil, err
		}
		startAt = startAt.Add(1 * time.Second)
		thirdparty, err := server.thirdparty.CreateStartAt(ctx, params.InAppPurchaseSku, &params.GoogleBillingServiceSku, startAt)
		if err != nil {
			return nil, server.errorz.IsThirdPartyError(err)
		}
		return ToThirdParty(thirdparty), nil
	}
	thirdparty, err := server.thirdparty.Create(ctx, params.InAppPurchaseSku, &params.GoogleBillingServiceSku)
	if err != nil {
		return nil, server.errorz.IsThirdPartyError(err)
	}
	return ToThirdParty(thirdparty), nil
}

func (server ThirdPartyServer) GetThirdParties(ctx context.Context, params *thirdpartyrpc.GetThirdPartiesParams) (*thirdpartyrpc.GetThirdPartiesResponse, error) {
	err := server.valid.GetThirdParties(params)
	if err != nil {
		return nil, err
	}
	parties, err := server.thirdparty.GetThirdParties(ctx)
	if err != nil {
		return nil, server.errorz.IsThirdPartyError(err)
	}
	return &thirdpartyrpc.GetThirdPartiesResponse{ThirdParties: ToThirdParties(parties)}, nil
}

func (server ThirdPartyServer) GetThirdParty(ctx context.Context, params *thirdpartyrpc.GetThirdPartyParams) (*thirdpartyrpc.ThirdParty, error) {
	err := server.valid.GetThirdParty(params)
	if err != nil {
		return nil, err
	}
	var thirdparty *ent.ThirdParty
	if !validator2.IsEmpty(params.InAppPurchaseSku) {
		thirdparty, err = server.thirdparty.GetInAppPurchaseBySku(ctx, params.InAppPurchaseSku)
	}
	if !validator2.IsEmpty(params.GoogleBillingServiceSku) {
		thirdparty, err = server.thirdparty.GetBillingBySku(ctx, params.GoogleBillingServiceSku)
	}
	if params.Id != 0 {
		thirdparty, err = server.thirdparty.GetByID(ctx, int(params.Id))
	}
	if err != nil {
		return nil, server.errorz.IsThirdPartyError(err)
	}
	return ToThirdParty(thirdparty), nil
}

func NewThirdPartyServer(client *ent.Client) ThirdPartyServer {
	return ThirdPartyServer{
		thirdparty:thirdpartyrepo.NewThirdPartyRepository(client),
		valid:validator.NewThirdPartyServerValidator(),
		errorz:errorz.NewThirdPartyErrorz(),
	}
}