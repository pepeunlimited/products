package twirp

import (
	"context"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/pkg/rpc/price"
	"github.com/pepeunlimited/products/pkg/rpc/thirdpartyprice"
	"github.com/twitchtv/twirp"
	"testing"
	"time"
)


func TestPriceServer_GetPrice(t *testing.T) {
	ctx := context.TODO()
	server := NewPriceServer(ent.NewEntClient())
	server.products.Wipe(ctx)
	product, err := server.products.CreateProduct(ctx, "sku")
	price, err := server.CreatePrice(ctx, &price.CreatePriceParams{
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
	createdPrice, err := server.CreatePrice(ctx, &price.CreatePriceParams{
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
	fromServer, err := server.GetPrice(ctx, &price.GetPriceParams{
		ProductSku: "sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if fromServer.Id != createdPrice.Id {
		t.FailNow()
	}
	_, err = server.EndPriceAt(ctx, &price.EndPriceAtParams{
		EndAtDay:   int32(now.Day()),
		EndAtMonth: int32(now.Month()),
		Params: &price.GetPriceParams{
			ProductId: int64(product.ID),
		},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	thirdpartypriceServer := NewThirdPartyPriceServer(ent.NewEntClient())
	party,_ := thirdpartypriceServer.CreateThirdPartyPrice(ctx, &thirdpartyprice.CreateThirdPartyPriceParams{
		InAppPurchaseSku: "in-app-purchase-sku",
	})
	createdPrice, err = server.CreatePrice(ctx, &price.CreatePriceParams{
		StartAtDay:   int32(now.Day()),
		StartAtMonth: int32(now.Month()),
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        5,
		Discount:     5,
		ProductId:    int64(product.ID),
		ThirdPartyId: party.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err = server.GetPrice(ctx, &price.GetPriceParams{
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
	if fromServer.Id != createdPrice.Id {
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
	fromServer, err := server.CreatePrice(ctx, &price.CreatePriceParams{
		StartAtDay:   0,
		StartAtMonth: 0,
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        3,
		Discount:     0,
		ProductId:    int64(product.ID),
		ThirdPartyId: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.GetPrice(ctx, &price.GetPriceParams{
		ProductId: int64(product.ID),
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Code() != twirp.InvalidArgument {
		t.FailNow()
	}
	_, err = server.GetPrice(ctx, &price.GetPriceParams{
		ProductSku: "sku",
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Code() != twirp.InvalidArgument {
		t.FailNow()
	}
	price, err := server.GetPrice(ctx, &price.GetPriceParams{
		PriceId: fromServer.Id,
	})
	if err != nil {
		t.Error(err)
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
	_, err = server.CreatePrice(ctx, &price.CreatePriceParams{
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
	fromServer, err := server.GetPrice(ctx, &price.GetPriceParams{
		ProductId: int64(product.ID),
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromDB, err := server.prices.GetPriceByID(ctx, int(fromServer.Id), true, true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if fromDB.Price != uint16(fromServer.Price) {
		t.FailNow()
	}
	fromServer, err = server.GetPrice(ctx, &price.GetPriceParams{
		ProductSku: "sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromDB, err = server.prices.GetPriceByID(ctx, int(fromServer.Id), true, true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if fromDB.Price != uint16(fromServer.Price) {
		t.FailNow()
	}
}

func  TestPriceServer_GetPriceWithThirdPartyPrices(t *testing.T) {
	ctx := context.TODO()
	server := NewPriceServer(ent.NewEntClient())
	server.products.Wipe(ctx)
	product, err := server.products.CreateProduct(ctx, "sku")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	thirdpartypriceServer := NewThirdPartyPriceServer(ent.NewEntClient())
	thirdparty, err := thirdpartypriceServer.CreateThirdPartyPrice(ctx, &thirdpartyprice.CreateThirdPartyPriceParams{
		InAppPurchaseSku: "in-app-sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	prices, err := server.CreatePrice(ctx, &price.CreatePriceParams{
		StartAtDay:   0,
		StartAtMonth: 0,
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        1,
		Discount:     1,
		ProductId:    int64(product.ID),
		ThirdPartyId: thirdparty.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	now := time.Now().UTC()
	server.EndPriceAt(ctx, &price.EndPriceAtParams{
		Params: &price.GetPriceParams{
			PriceId: prices.Id,
		},
		EndAtDay:   int32(now.Day()),
		EndAtMonth: int32(now.Month()),
	})
	_, err = server.CreatePrice(ctx, &price.CreatePriceParams{
		StartAtDay:   int32(now.Day()),
		StartAtMonth: int32(now.Month()),
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        3,
		Discount:     3,
		ProductId:    int64(product.ID),
		ThirdPartyId: thirdparty.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	//price, err := server.GetPrice(ctx, &price.GetPriceParams{ProductId: int64(plan.ID)})
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
	//if price.Price != 3 {
	//	t.FailNow()
	//}
	//if price.ProductId == 0 {
	//	t.FailNow()
	//}
	//if price.ThirdPartyId == 0 {
	//	t.FailNow()
	//}
	//if price.PlanId == 0 {
	//	t.FailNow()
	//}
}