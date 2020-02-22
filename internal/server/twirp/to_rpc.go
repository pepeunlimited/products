package twirp

import (
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/pkg/rpc/plan"
	"github.com/pepeunlimited/prices/pkg/rpc/price"
	"github.com/pepeunlimited/prices/pkg/rpc/product"
	"github.com/pepeunlimited/prices/pkg/rpc/subscription"
	"github.com/pepeunlimited/prices/pkg/rpc/thirdpartyprice"
	"time"
)

func ToThirdParty(thirdparty *ent.ThirdParty) *thirdpartyprice.ThirdParty {
	return &thirdpartyprice.ThirdParty{
		Id:                      int32(thirdparty.ID),
		InAppPurchaseSku:        thirdparty.InAppPurchaseSku,
		GoogleBillingServiceSku: thirdparty.GoogleBillingServiceSku,
		StartAt:                 thirdparty.StartAt.Format(time.RFC3339),
		EndAt:                   thirdparty.EndAt.Format(time.RFC3339),
	}
}

func ToThirdParties(parties []*ent.ThirdParty) []*thirdpartyprice.ThirdParty {
	list := make([]*thirdpartyprice.ThirdParty, 0)
	for _, party := range parties {
		list = append(list, ToThirdParty(party))
	}
	return list
}

func ToPrice(price *ent.Price) *price.Price {
	p := &price.Price{
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

func ToPrices(prices []*ent.Price) []*price.Price {
	list := make([]*price.Price, 0)
	for _, price := range prices {
		list = append(list, ToPrice(price))
	}
	return list
}

func ToPlan(plan *ent.Plan) *plan.Plan {
	return &plan.Plan{
		Id:          int64(plan.ID),
		TitleI18NId: plan.TitleI18nID,
		Unit:        plan.Unit,
		Length: 	 int32(plan.Length),
	}
}

func ToProduct(product *ent.Product) *product.Product {
	return &product.Product{
		Sku: product.Sku,
		Id:  int64(product.ID),
	}
}

func ToSubscription(subscription *ent.Subscription) *subscription.Subscription {
	s := &subscription.Subscription{
		Id:     	int64(subscription.ID),
		UserId: 	subscription.UserID,
		StartAt:	subscription.StartAt.Format(time.RFC3339),
		EndAt:		subscription.EndAt.Format(time.RFC3339),
	}
	if subscription.Edges.Plans != nil {
		s.PlanId = int64(subscription.Edges.Plans.ID)
	}
	return s
}

func ToSubscriptions(subscriptions []*ent.Subscription) []*subscription.Subscription {
	list := make([]*subscription.Subscription, 0)
	for _, subscription := range subscriptions {
		list = append(list, ToSubscription(subscription))
	}
	return list
}

func ToProducts(products []*ent.Product) []*product.Product {
	list := make([]*product.Product, 0)
	for _,product := range products {
		list = append(list, ToProduct(product))
	}
	return list
}

func ToPlans(plans []*ent.Plan) []*plan.Plan {
	list := make([]*plan.Plan, 0)
	for _, plan := range plans {
		list = append(list, ToPlan(plan))
	}
	return list
}