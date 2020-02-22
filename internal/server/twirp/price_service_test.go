package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/plan"
	"github.com/pepeunlimited/prices/pkg/pricerpc"
	"github.com/pepeunlimited/prices/pkg/thirdpartypricerpc"
	"github.com/twitchtv/twirp"
	"testing"
	"time"
)


func TestPriceServer_GetPrice(t *testing.T) {
	ctx := context.TODO()
	server := NewPriceServer(ent.NewEntClient())
	server.products.Wipe(ctx)
	product, err := server.products.CreateProduct(ctx, "sku")
	price, err := server.CreatePrice(ctx, &pricerpc.CreatePriceParams{
		StartAtDay:   0,
		StartAtMonth: 0,
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        3,
		Discount:     3,
		ProductId:    int64(product.ID),
		PlanId:       0,
		ThirdPartyId: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if price.ProductId == 0 {
		t.FailNow()
	}
}

func TestPriceServer_CreatePrice(t *testing.T) {
	ctx := context.TODO()
	server := NewPriceServer(ent.NewEntClient())
	server.products.Wipe(ctx)
	product, err := server.products.CreateProduct(ctx, "sku")
	now := time.Now()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	price, err := server.CreatePrice(ctx, &pricerpc.CreatePriceParams{
		StartAtDay:   0,
		StartAtMonth: 0,
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        3,
		Discount:     3,
		ProductId:    int64(product.ID),
		PlanId:       0,
		ThirdPartyId: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err := server.GetPrice(ctx, &pricerpc.GetPriceParams{
		ProductSku: "sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if fromServer.Id != price.Id {
		t.FailNow()
	}
	_, err = server.EndPrice(ctx, &pricerpc.EndPriceParams{
		EndAtDay:   int32(now.Day()),
		EndAtMonth: int32(now.Month()),
		Params: &pricerpc.GetPriceParams{
			ProductId: int64(product.ID),
		},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	thirdPartyServer := NewThirdPartyServer(ent.NewEntClient())
	party,_ := thirdPartyServer.CreateThirdParty(ctx, &thirdpartypricerpc.CreateThirdPartyParams{
		InAppPurchaseSku: "in-app-purchase-sku",
	})
	price, err = server.CreatePrice(ctx, &pricerpc.CreatePriceParams{
		StartAtDay:   int32(now.Day()),
		StartAtMonth: int32(now.Month()),
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        5,
		Discount:     5,
		ProductId:    int64(product.ID),
		PlanId:       0,
		ThirdPartyId: party.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err = server.GetPrice(ctx, &pricerpc.GetPriceParams{
		ProductSku: "sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if fromServer.Price != 5 {
		t.FailNow()
	}
	if fromServer.Discount != 5 {
		t.FailNow()
	}
	if fromServer.Id != price.Id {
		t.FailNow()
	}
	if fromServer.ThirdPartyId == 0 {
		t.FailNow()
	}
	if fromServer.ProductId == 0 {
		t.FailNow()
	}

}

func TestPriceServer_GetPriceByProductIdIsSubscribableTrue(t *testing.T) {
	ctx := context.TODO()
	server := NewPriceServer(ent.NewEntClient())
	server.products.Wipe(ctx)
	product, err := server.products.CreateProduct(ctx, "sku")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	planrepos := plan.NewPlanRepository(ent.NewEntClient())
	plan, err := planrepos.Create(ctx, 1, 1, plan.Days)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.CreatePrice(ctx, &pricerpc.CreatePriceParams{
		StartAtDay:   0,
		StartAtMonth: 0,
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        3,
		Discount:     0,
		ProductId:    int64(product.ID),
		PlanId:       int64(plan.ID),
		ThirdPartyId: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.GetPrice(ctx, &pricerpc.GetPriceParams{
		ProductId: int64(product.ID),
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Code() != twirp.InvalidArgument {
		t.FailNow()
	}
	_, err = server.GetPrice(ctx, &pricerpc.GetPriceParams{
		ProductSku: "sku",
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Code() != twirp.InvalidArgument {
		t.FailNow()
	}
	price, err := server.GetPrice(ctx, &pricerpc.GetPriceParams{
		PlanId: int64(plan.ID),
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if price.PlanId != int64(plan.ID) {
		t.FailNow()
	}
	if price.ProductId != int64(product.ID) {
		t.FailNow()
	}
}

func TestPriceServer_GetPriceByProductIdIsSubscribableFalse(t *testing.T) {
	ctx := context.TODO()
	server := NewPriceServer(ent.NewEntClient())
	server.products.Wipe(ctx)
	product, err := server.products.CreateProduct(ctx, "sku")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.CreatePrice(ctx, &pricerpc.CreatePriceParams{
		StartAtDay:   0,
		StartAtMonth: 0,
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        3,
		Discount:     3,
		ProductId:    int64(product.ID),
		ThirdPartyId: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err := server.GetPrice(ctx, &pricerpc.GetPriceParams{
		ProductId: int64(product.ID),
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromDB, err := server.prices.GetPriceByID(ctx, int(fromServer.Id), true, true,true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if fromDB.Price != uint16(fromServer.Price) {
		t.FailNow()
	}
	fromServer, err = server.GetPrice(ctx, &pricerpc.GetPriceParams{
		ProductSku: "sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromDB, err = server.prices.GetPriceByID(ctx, int(fromServer.Id), true, true,true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if fromDB.Price != uint16(fromServer.Price) {
		t.FailNow()
	}
}

func  TestPriceServer_GetSubscriptionPrices(t *testing.T) {
	ctx := context.TODO()
	server := NewPriceServer(ent.NewEntClient())
	server.products.Wipe(ctx)
	product, err := server.products.CreateProduct(ctx, "sku")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	plan, err := server.plans.Create(ctx, 1, uint8(12), plan.Days)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	thirdpartyServer := NewThirdPartyServer(ent.NewEntClient())
	thirdparty, err := thirdpartyServer.CreateThirdParty(ctx, &thirdpartypricerpc.CreateThirdPartyParams{
		InAppPurchaseSku: "in-app-sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	prices, err := server.CreatePrice(ctx, &pricerpc.CreatePriceParams{
		StartAtDay:   0,
		StartAtMonth: 0,
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        1,
		Discount:     1,
		ProductId:    int64(product.ID),
		PlanId:       int64(plan.ID),
		ThirdPartyId: thirdparty.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	now := time.Now().UTC()
	server.EndPrice(ctx, &pricerpc.EndPriceParams{
		Params: &pricerpc.GetPriceParams{
			PriceId: prices.Id,
		},
		EndAtDay:   int32(now.Day()),
		EndAtMonth: int32(now.Month()),
	})
	_, err = server.CreatePrice(ctx, &pricerpc.CreatePriceParams{
		StartAtDay:   int32(now.Day()),
		StartAtMonth: int32(now.Month()),
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        3,
		Discount:     3,
		ProductId:    int64(product.ID),
		PlanId:       int64(plan.ID),
		ThirdPartyId: thirdparty.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	subscriptions, err := server.GetSubscriptionPrices(ctx, &pricerpc.GetSubscriptionPricesParams{
		ProductId: int64(product.ID),
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(subscriptions.Prices) != 1 {
		t.FailNow()
	}
	if subscriptions.Prices[0].Price != 3 {
		t.FailNow()
	}
	price, err := server.GetPrice(ctx, &pricerpc.GetPriceParams{PlanId: int64(plan.ID)})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if price.Price != 3 {
		t.FailNow()
	}
	if price.ProductId == 0 {
		t.FailNow()
	}
	if price.ThirdPartyId == 0 {
		t.FailNow()
	}
	if price.PlanId == 0 {
		t.FailNow()
	}
}