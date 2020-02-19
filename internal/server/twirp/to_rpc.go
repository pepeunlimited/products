package twirp

import (
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/pkg/pricerpc"
	"github.com/pepeunlimited/prices/pkg/thirdpartyrpc"
	"time"
)

func ToThirdParty(thirdparty *ent.ThirdParty) *thirdpartyrpc.ThirdParty {
	return &thirdpartyrpc.ThirdParty{
		Id:                      int32(thirdparty.ID),
		InAppPurchaseSku:        thirdparty.InAppPurchaseSku,
		GoogleBillingServiceSku: thirdparty.GoogleBillingServiceSku,
		StartAt:                 thirdparty.StartAt.Format(time.RFC3339),
		EndAt:                   thirdparty.EndAt.Format(time.RFC3339),
	}
}

func ToThirdParties(parties []*ent.ThirdParty) []*thirdpartyrpc.ThirdParty {
	list := make([]*thirdpartyrpc.ThirdParty, 0)
	for _, party := range parties {
		list = append(list, ToThirdParty(party))
	}
	return list
}

func ToPrice(price *ent.Price) *pricerpc.Price {
	return &pricerpc.Price{
		Id:       int64(price.ID),
		Price:    uint32(price.Price),
		Discount: uint32(price.Discount),
		StartAt:  price.StartAt.Format(time.RFC3339),
		EndAt:    price.EndAt.Format(time.RFC3339),
	}
}