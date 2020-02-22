package plan

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/product"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/thirdpartyprice"
	"testing"
)

func TestPlanMySQL_Create(t *testing.T) {
	ctx := context.TODO()
	repo := NewPlanRepository(ent.NewEntClient())

	titleI18nId := int64(2)
	length 	:= uint8(2)
	unit 	:= PlanUnitFromString("weeks")

	productRepo := product.New(ent.NewEntClient())
	product, err := productRepo.CreateProduct(ctx, "sku")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	plan0, err := repo.CreateNewPlan(ctx, titleI18nId, length, unit, 1, 1, product.ID, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if plan0.Unit == "UNKNOWN" {
		t.FailNow()
	}
}

//func TestPlanMySQL_LengthByPlansID(t *testing.T) {
//	ctx := context.TODO()
//	plans := NewPlanRepository(ent.NewEntClient())
//	plans.Wipe(ctx)
//	startAt := time.Now()
//	endAt 	:= time.Now()
//	titleI18nId := int64(2)
//	length 	:= uint8(1)
//	one day
	//unit 	:= PlanUnitFromString("days")
	//planOneDay, err := plans.Create(ctx, titleI18nId, length, unit)
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
	//if planOneDay.Unit == "UNKNOWN" {
	//	t.FailNow()
	//}
	//endAt, err = plans.EndAtByPlanID(ctx, startAt, planOneDay.ID)
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
	//if !startAt.AddDate(0,0,1).UTC().Equal(endAt) {
	//	t.FailNow()
	//}
	//one week
	//unit 	= PlanUnitFromString("weeks")
	//planOneWeek, err := plans.Create(ctx, titleI18nId, length, unit)
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
	//endAt, err = plans.EndAtByPlanID(ctx, startAt, planOneWeek.ID)
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
	//if !startAt.AddDate(0,0,7).UTC().Equal(endAt) {
	//	t.FailNow()
	//}
	//one month
	//unit 	= PlanUnitFromString("months")
	//planOneMonth, err := plans.Create(ctx, titleI18nId, length, unit)
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
	//endAt, err = plans.EndAtByPlanID(ctx, startAt, planOneMonth.ID)
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
	//if !startAt.AddDate(0,1,0).UTC().Equal(endAt) {
	//	t.FailNow()
	//}
	//one year
	//unit 	= PlanUnitFromString("years")
	//planOneYear, err := plans.Create(ctx, titleI18nId, length, unit)
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
	//endAt, err = plans.EndAtByPlanID(ctx, startAt, planOneYear.ID)
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
	//if !startAt.AddDate(1,0,0).UTC().Equal(endAt) {
	//	t.FailNow()
	//}
//}

func TestPlanMySQL_GetPlans(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	productrepo := product.New(client)
	plansrepo 	:= NewPlanRepository(client)
	plansrepo.Wipe(ctx)

	thirdparty := thirdpartyprice.New(client)
	thirdpartyprice, err := thirdparty.Create(ctx, "apple", nil, thirdpartyprice.Consumable)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	product,_ 		:= productrepo.CreateProduct(ctx, "sku-1")
	plansrepo.CreateNewPlan(ctx, 1, 1, Days, 1,1, product.ID, &thirdpartyprice.ID)
	plansrepo.CreateNewPlan(ctx, 1, 2, Days, 2,2, product.ID, nil)
	plans, err := plansrepo.GetPlans(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(plans) != 2 {
		t.FailNow()
	}
	if plans[0].TitleI18nID != 1 {
		t.FailNow()
	}
	plan := plans[0]
	at, err := plansrepo.EndPlanAt(ctx, 2, 1, plan.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if at.EndAt.Month() != 2 {
		t.FailNow()
	}
	if at.EndAt.Day() != 1 {
		t.FailNow()
	}
	plans2, err := plansrepo.GetPlans(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(plans2) != 1 {
		t.FailNow()
	}
}