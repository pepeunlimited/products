package planrepo

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/pricerepo"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/productrepo"
	"testing"
	"time"
)

func TestPlanMySQL_Create(t *testing.T) {
	ctx := context.TODO()
	repo := NewPlanRepository(ent.NewEntClient())

	titleI18nId := int64(2)
	length 	:= uint8(2)
	unit 	:= PlanUnitFromString("weeks")

	plan0, err := repo.Create(ctx, titleI18nId, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if plan0.Unit == "UNKNOWN" {
		t.FailNow()
	}
}
func TestPlanMySQL_LengthByPlansID(t *testing.T) {
	ctx := context.TODO()
	plans := NewPlanRepository(ent.NewEntClient())
	plans.Wipe(ctx)

	startAt := time.Now()
	endAt 	:= time.Now()
	titleI18nId := int64(2)
	length 	:= uint8(1)

	// one day
	unit 	:= PlanUnitFromString("days")
	planOneDay, err := plans.Create(ctx, titleI18nId, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if planOneDay.Unit == "UNKNOWN" {
		t.FailNow()
	}
	endAt, err = plans.LengthByPlansID(ctx, startAt, planOneDay.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !startAt.AddDate(0,0,1).UTC().Equal(endAt) {
		t.FailNow()
	}

	// one week
	unit 	= PlanUnitFromString("weeks")
	planOneWeek, err := plans.Create(ctx, titleI18nId, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt, err = plans.LengthByPlansID(ctx, startAt, planOneWeek.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !startAt.AddDate(0,0,7).UTC().Equal(endAt) {
		t.FailNow()
	}

	// one month
	unit 	= PlanUnitFromString("months")
	planOneMonth, err := plans.Create(ctx, titleI18nId, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt, err = plans.LengthByPlansID(ctx, startAt, planOneMonth.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !startAt.AddDate(0,1,0).UTC().Equal(endAt) {
		t.FailNow()
	}

	// one year
	unit 	= PlanUnitFromString("years")
	planOneYear, err := plans.Create(ctx, titleI18nId, length, unit)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt, err = plans.LengthByPlansID(ctx, startAt, planOneYear.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !startAt.AddDate(1,0,0).UTC().Equal(endAt) {
		t.FailNow()
	}
}


func TestPlanMySQL_GetPlans(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	productrepo := productrepo.NewProductRepository(client)
	pricerepo 	:= pricerepo.NewPriceRepository(client)
	plansrepo 	:= NewPlanRepository(client)
	plansrepo.Wipe(ctx)

	titleI18nId1 	:= int64(1)
	length 			:= uint8(1)
	unit 			:= PlanUnitFromString("days")
	product1,_ 		:= productrepo.CreateProduct(ctx, "sku-1")

	plan1,_ := plansrepo.Create(ctx, titleI18nId1, length, unit)
	plan2,_ := plansrepo.Create(ctx, titleI18nId1, length, unit)

	pricerepo.CreateNewPrice(ctx, uint16(0), product1.ID, nil, &plan1.ID)
	pricerepo.CreateNewPrice(ctx, uint16(2), product1.ID, nil, &plan2.ID)
	pricerepo.CreateNewPrice(ctx, uint16(2), product1.ID, nil, &plan2.ID)

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
}