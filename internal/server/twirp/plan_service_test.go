package twirp

import (
	"context"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/pkg/rpc/plan"
	"github.com/pepeunlimited/products/pkg/rpc/product"
	"github.com/pepeunlimited/products/pkg/rpc/thirdpartyprice"
	"testing"
	"time"
)

func TestPlanServer_CreatePlanAndGetPlan(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()

	server := NewPlanServer(client)
	server.plans.Wipe(ctx)

	productServer := NewProductServer(client)
	fromServerProduct, err := productServer.CreateProduct(ctx, &product.CreateProductParams{
		Sku: "skuu",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err := server.CreatePlan(ctx, &plan.CreatePlanParams{
		TitleI18NId:       0,
		Length:            12,
		Unit:              "days",
		Price:             1,
		Discount:          1,
		ProductId:         fromServerProduct.Id,
		ThirdPartyPriceId: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err = server.GetPlan(ctx, &plan.GetPlanParams{PlanId: fromServer.Id})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestPlanServer_GetPlans(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()

	server := NewPlanServer(client)
	server.plans.Wipe(ctx)
	now := time.Now()
	thirdpartypriceServer := NewThirdPartyPriceServer(client)
	thirdpartyprice, err := thirdpartypriceServer.CreateThirdPartyPrice(ctx, &thirdpartyprice.CreateThirdPartyPriceParams{
		InAppPurchaseSku: "sku",
		StartAtMonth:     int32(now.Month()),
		StartAtDay:       int32(now.Day()),
		Type:             "CONSUMABLE",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	productServer := NewProductServer(client)
	fromServerProduct, err := productServer.CreateProduct(ctx, &product.CreateProductParams{
		Sku: "skuu",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err := server.CreatePlan(ctx, &plan.CreatePlanParams{
		TitleI18NId:       0,
		Length:            12,
		Unit:              "days",
		Price:             1,
		Discount:          1,
		StartAtDay:		   int32(now.Day()),
		StartAtMonth:	   int32(now.Month()),
		ProductId:         fromServerProduct.Id,
		ThirdPartyPriceId: thirdpartyprice.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err = server.GetPlan(ctx, &plan.GetPlanParams{PlanId: fromServer.Id})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}