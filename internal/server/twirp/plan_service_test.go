package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/pkg/planrpc"
	"testing"
)

func TestPlanServer_CreatePlanAndGetPlan(t *testing.T) {
	ctx := context.TODO()
	server := NewPlanServer(ent.NewEntClient())
	server.plans.Wipe(ctx)
	plan, err := server.CreatePlan(ctx, &planrpc.CreatePlanParams{
		TitleI18NId: 1,
		Length:      12,
		Unit:        "days",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	plan, err = server.GetPlan(ctx, &planrpc.GetPlanParams{PlanId: plan.Id})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}