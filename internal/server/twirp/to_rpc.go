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
	p := &pricerpc.Price{
		Id:       int64(price.ID),
		Price:    uint32(price.Price),
		Discount: uint32(price.Discount),
		StartAt:  price.StartAt.Format(time.RFC3339),
		EndAt:    price.EndAt.Format(time.RFC3339),
	}
	if price.Edges.ThirdParties != nil {
		p.ThirdPartyId = int32(price.Edges.ThirdParties.ID)
	}
	if price.Edges.Plans != nil {
		p.PlanId = int64(price.Edges.Plans.ID)
	}
	if price.Edges.Products != nil {
		p.ProductId = int64(price.Edges.Products.ID)
	}
	return p
}

func ToPrices(prices []*ent.Price) []*pricerpc.Price {
	list := make([]*pricerpc.Price, 0)
	for _, price := range prices {
		list = append(list, ToPrice(price))
	}
	return list
}