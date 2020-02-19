package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/planrepo"
	"github.com/pepeunlimited/prices/pkg/pricerpc"
	"github.com/pepeunlimited/prices/pkg/thirdpartyrpc"
	"github.com/twitchtv/twirp"
	"log"
	"testing"
	"time"
)

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
	party,_ := thirdPartyServer.CreateThirdParty(ctx, &thirdpartyrpc.CreateThirdPartyParams{
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
	plan := planrepo.NewPlanRepository(ent.NewEntClient())
	create, err := plan.Create(ctx, 1, 1, planrepo.Days)
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
		PlanId:       int64(create.ID),
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
	log.Print(fromServer)
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