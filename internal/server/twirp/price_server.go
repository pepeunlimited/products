package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/server/validator"
	"github.com/pepeunlimited/prices/pkg/pricerpc"
)

type PriceServer struct {
	client *ent.Client
	valid validator.PriceServerValidator
}

func (server PriceServer) CreatePrice(ctx context.Context, params *pricerpc.CreatePriceParams) (*pricerpc.Price, error) {
	err := server.valid.CreatePrice(params)
}

func (server PriceServer) GetPrice(ctx context.Context, params *pricerpc.GetPriceParams) (*pricerpc.Price, error) {
	panic("implement me")
}

func NewPriceServer(client *ent.Client) PriceServer {
	return PriceServer{
		client:client,
		valid:validator.NewPriceServerValidator(),
	}
}