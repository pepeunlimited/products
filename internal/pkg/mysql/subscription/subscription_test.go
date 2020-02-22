package subscription

import (
	"context"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/internal/pkg/mysql/plan"
	"github.com/pepeunlimited/products/internal/pkg/mysql/product"
	"testing"
	"time"
)

func TestSubscriptionMySQL_Create(t *testing.T) {
	client := ent.NewEntClient()
	ctx := context.TODO()
	subscriptionrepo := New(client)
	plans 		 	 := plan.New(client)
	plans.Wipe(ctx)
	startAt 	:= time.Now().UTC()
	userID  	:= int64(1)
	unit  		:= plan.PlanUnitFromString("days")

	productRepo := product.New(client)
	product, err := productRepo.CreateProduct(ctx, "sku")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	plan, err 	:= plans.CreateNewPlan(ctx, 1, 1, unit, 0, 0, product.ID, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt := startAt.Add(24 * time.Hour)
	subscription, err := subscriptionrepo.Create(ctx, userID, startAt, endAt, plan.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if subscription.UserID != userID {
		t.FailNow()
	}

}