package twirp

import (
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/pkg/planrpc"
	"github.com/pepeunlimited/prices/pkg/pricerpc"
	"github.com/pepeunlimited/prices/pkg/productrpc"
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

func ToPlan(plan *ent.Plan) *planrpc.Plan {
	return &planrpc.Plan{
		Id:          int64(plan.ID),
		TitleI18NId: plan.TitleI18nID,
		Unit:        plan.Unit,
		Length: 	 int32(plan.Length),
	}
}

func ToProduct(product *ent.Product) *productrpc.Product {
	return &productrpc.Product{
		Sku: product.Sku,
		Id:  int64(product.ID),
	}
}