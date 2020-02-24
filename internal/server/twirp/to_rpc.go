package twirp

import (
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/pkg/rpc/plan"
	"github.com/pepeunlimited/products/pkg/rpc/price"
	"github.com/pepeunlimited/products/pkg/rpc/product"
	"github.com/pepeunlimited/products/pkg/rpc/subscription"
	"github.com/pepeunlimited/products/pkg/rpc/thirdpartyprice"
	"time"
)

func ToThirdParty(from *ent.ThirdPartyPrice) *thirdpartyprice.ThirdPartyPrice {
	return &thirdpartyprice.ThirdPartyPrice{
		Id:                      int64(from.ID),
		InAppPurchaseSku:        from.InAppPurchaseSku,
		GoogleBillingServiceSku: from.GoogleBillingServiceSku,
		StartAt:                 from.StartAt.Format(time.RFC3339),
		EndAt:                   from.EndAt.Format(time.RFC3339),
	}
}

func ToThirdParties(from []*ent.ThirdPartyPrice) []*thirdpartyprice.ThirdPartyPrice {
	list := make([]*thirdpartyprice.ThirdPartyPrice, 0)
	for _, party := range from {
		list = append(list, ToThirdParty(party))
	}
	return list
}

func ToPrice(from *ent.Price) *price.Price {
	p := &price.Price{
		Id:       int64(from.ID),
		Price:    from.Price,
		Discount: from.Discount,
		StartAt:  from.StartAt.Format(time.RFC3339),
		EndAt:    from.EndAt.Format(time.RFC3339),
	}
	if from.Edges.ThirdPartyPrices != nil {
		p.ThirdPartyId = int64(from.Edges.ThirdPartyPrices.ID)
	}
	if from.Edges.Products != nil {
		p.ProductId = int64(from.Edges.Products.ID)
	}
	return p
}

func ToPrices(from []*ent.Price) []*price.Price {
	list := make([]*price.Price, 0)
	for _, price := range from {
		list = append(list, ToPrice(price))
	}
	return list
}

func ToPlan(from *ent.Plan) *plan.Plan {
	p := &plan.Plan{
		Id:                int64(from.ID),
		TitleI18NId:       from.TitleI18nID,
		Unit:              from.Unit,
		Length:            int32(from.Length),
		Price:             from.Price,
		Discount:          from.Discount,
		StartAt:           from.StartAt.Format(time.RFC3339),
		EndAt:             from.EndAt.Format(time.RFC3339),
	}
	if from.Edges.Products != nil {
		p.ProductId = int64(from.Edges.Products.ID)
	}
	if from.Edges.ThirdPartyPrices != nil {
		p.ThirdPartyPriceId = int64(from.Edges.ThirdPartyPrices.ID)
	}
	return p
}

func ToProduct(from *ent.Product) *product.Product {
	return &product.Product{
		Sku: from.Sku,
		Id:  int64(from.ID),
	}
}

func ToSubscription(from *ent.Subscription) *subscription.Subscription {
	s := &subscription.Subscription{
		Id:     	int64(from.ID),
		UserId: 	from.UserID,
		StartAt:	from.StartAt.Format(time.RFC3339),
		EndAt:		from.EndAt.Format(time.RFC3339),
	}
	if from.Edges.Plans != nil {
		s.PlanId = int64(from.Edges.Plans.ID)
	}
	return s
}

func ToSubscriptions(from []*ent.Subscription) []*subscription.Subscription {
	list := make([]*subscription.Subscription, 0)
	for _, subscription := range from {
		list = append(list, ToSubscription(subscription))
	}
	return list
}

func ToProducts(from []*ent.Product) []*product.Product {
	list := make([]*product.Product, 0)
	for _,product := range from {
		list = append(list, ToProduct(product))
	}
	return list
}

func ToPlans(from []*ent.Plan) []*plan.Plan {
	list := make([]*plan.Plan, 0)
	for _, plan := range from {
		list = append(list, ToPlan(plan))
	}
	return list
}