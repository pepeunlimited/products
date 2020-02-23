package twirp

import (
	"context"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/pkg/rpc/plan"
	"github.com/pepeunlimited/products/pkg/rpc/price"
	"github.com/pepeunlimited/products/pkg/rpc/product"
	"github.com/pepeunlimited/products/pkg/rpc/subscription"
	"testing"
	"time"
)

func TestSubscriptionServer_StartSubscription(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	server := NewSubscriptionServer(client)
	server.plans.Wipe(ctx)
	planServer := NewPlanServer(client)

	productServer := NewProductServer(client)
	product, err := productServer.CreateProduct(ctx, &product.CreateProductParams{
		Sku: "sku-product",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	length := 5
	plan, err := planServer.CreatePlan(ctx, &plan.CreatePlanParams{
		TitleI18NId: 1,
		Length:      int32(length),
		Unit:        "days",
		ProductId:	product.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	priceServer 	 := NewPriceServer(client)
	_, err 			  = priceServer.CreatePrice(ctx, &price.CreatePriceParams{
		StartAtDay:   0,
		StartAtMonth: 0,
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        3,
		Discount:     3,
		ProductId:    product.Id,
		ThirdPartyId: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err := server.StartSubscription(ctx, &subscription.StartSubscriptionParams{
		UserId: 1,
		PlanId: plan.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	now := time.Now().UTC()
	startAt, err := time.Parse(time.RFC3339, fromServer.StartAt)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt, err := time.Parse(time.RFC3339, fromServer.EndAt)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if now.Month() != startAt.UTC().Month() {
		t.FailNow()
	}
	if now.Day() != startAt.Day() {
		t.FailNow()
	}
	if now.Add((24 * time.Hour) * time.Duration(length)).Month() != endAt.Month() {
		t.FailNow()
	}
	if now.Add((24 * time.Hour) * time.Duration(length)).Day() != endAt.Day() {
		t.FailNow()
	}
	fromServer, err = server.GetSubscription(ctx, &subscription.GetSubscriptionParams{
		SubscriptionId: fromServer.Id,
		UserId:         1,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if fromServer.Id != fromServer.Id {
		t.FailNow()
	}
	subscriptions, err := server.GetSubscriptions(ctx, &subscription.GetSubscriptionsParams{
		UserId:    1,
		PageSize:  20,
		PageToken: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(subscriptions.Subscriptions) != 1 {
		t.FailNow()
	}
	if subscriptions.NextPageToken == 0 {
		t.FailNow()
	}
}