package price

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/product"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/thirdpartyprice"
	"testing"
)

func TestPriceMySQL_CreatePrice(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	products := product.New(client)
	products.Wipe(ctx)
	product,_ := products.CreateProduct(ctx, "STCKR-4")
	priceRepo := New(client)
	price := uint16(0)
	price3rdrepo := thirdpartyprice.New(client)
	iap, err := price3rdrepo.Create(ctx, "apple-sku", nil, thirdpartyprice.Consumable)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	created, err := priceRepo.CreateNewPrice(ctx, price,price, product.ID, &iap.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !created.EndAt.Equal(clock.InfinityAt()) {
		t.FailNow()
	}
	if !created.StartAt.Equal(clock.ZeroAt()) {
		t.FailNow()
	}
	ended, err := priceRepo.EndAt(ctx, 2, 12, created.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt,_ := clock.ToMonthDate(2, 12)
	if !ended.EndAt.Equal(endAt) {
		t.FailNow()
	}
}

func TestPriceMySQL_GetPricesByProductID(t *testing.T) {
	ctx 		:= context.TODO()
	client 		:= ent.NewEntClient()
	productrepo := product.New(client)
	pricerepo   := New(client)
	productrepo.Wipe(ctx)
	product, err := productrepo.CreateProduct(ctx, "STCKR-4")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	initial, err := pricerepo.CreateNewPrice(ctx, uint16(3), uint16(3), product.ID, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	pricerepo.EndAt(ctx, 2, 14, initial.ID)
	startAt,_ := clock.ToMonthDate(2, 14)
	createCurrent, err := pricerepo.CreatePrice(ctx, uint16(5), uint16(5), product.ID, startAt, clock.InfinityAt(), nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	startAt,_ = clock.ToMonthDate(2, 15)
	_, err = pricerepo.CreatePrice(ctx, uint16(3),uint16(3), product.ID, startAt, clock.InfinityAt(), nil)
	if err == nil {
		t.FailNow()
	}
	if err != ErrInvalidStartAt {
		t.FailNow()
	}
	selectedCurrent, err := pricerepo.GetPriceByProductID(ctx, product.ID, false, false)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if selectedCurrent.ID != createCurrent.ID {
		t.FailNow()
	}

	id, err := pricerepo.GetPriceByProductID(ctx, product.ID, false,false)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if id == nil {
		t.FailNow()
	}
}