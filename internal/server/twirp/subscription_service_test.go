package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/pkg/planrpc"
	"github.com/pepeunlimited/prices/pkg/pricerpc"
	"github.com/pepeunlimited/prices/pkg/productrpc"
	"github.com/pepeunlimited/prices/pkg/subscriptionrpc"
	"testing"
	"time"
)

func TestSubscriptionServer_StartSubscription(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	server := NewSubscriptionServer(client)
	server.plans.Wipe(ctx)
	planServer := NewPlanServer(client)
	length := 5
	plan, err := planServer.CreatePlan(ctx, &planrpc.CreatePlanParams{
		TitleI18NId: 1,
		Length:      int32(length),
		Unit:        "days",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	productServer := NewProductServer(client)
	product, err := productServer.CreateProduct(ctx, &productrpc.CreateProductParams{
		Sku: "sku-product",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	priceServer 	 := NewPriceServer(client)
	_, err 			  = priceServer.CreatePrice(ctx, &pricerpc.CreatePriceParams{
		StartAtDay:   0,
		StartAtMonth: 0,
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        3,
		Discount:     3,
		ProductId:    product.Id,
		PlanId:       plan.Id,
		ThirdPartyId: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	subscription, err := server.StartSubscription(ctx, &subscriptionrpc.StartSubscriptionParams{
		UserId: 1,
		PlanId: plan.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	now := time.Now()
	startAt, err := time.Parse(time.RFC3339, subscription.StartAt)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt, err := time.Parse(time.RFC3339, subscription.EndAt)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if now.Month() != startAt.Month() {
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
	fromServer, err := server.GetSubscription(ctx, &subscriptionrpc.GetSubscriptionParams{
		SubscriptionId: subscription.Id,
		UserId:         1,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if fromServer.Id != subscription.Id {
		t.FailNow()
	}
	subscriptions, err := server.GetSubscriptions(ctx, &subscriptionrpc.GetSubscriptionsParams{
		UserId:    2,
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