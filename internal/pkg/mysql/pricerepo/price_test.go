package pricerepo

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/iapsourcerepo"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/productrepo"
	"log"
	"testing"
)

func TestPriceMySQL_CreatePrice(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	products := productrepo.NewProductRepository(client)
	products.Wipe(ctx)
	product,_ := products.CreateProduct(ctx, "STCKR-4")
	repo := NewPriceRepository(client)
	price := uint16(0)

	iapsource := iapsourcerepo.NewIapSourceRepository(client)
	iap, err := iapsource.Create(ctx, "apple-sku", nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	created, err := repo.CreateNewPrice(ctx, price, product.ID, &iap.ID, nil)
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
	ended, err := repo.EndAt(ctx, 2, 12, created.ID)
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
	productrepo := productrepo.NewProductRepository(client)
	pricerepo   := NewPriceRepository(client)
	productrepo.Wipe(ctx)
	product, err := productrepo.CreateProduct(ctx, "STCKR-4")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	initial, err := pricerepo.CreateNewPrice(ctx, uint16(3), product.ID, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	pricerepo.EndAt(ctx, 2, 14, initial.ID)
	startAt,_ := clock.ToMonthDate(2, 14)
	createCurrent, err := pricerepo.CreatePrice(ctx, uint16(5), product.ID, startAt, clock.InfinityAt(), nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	startAt,_ = clock.ToMonthDate(2, 15)
	_, err = pricerepo.CreatePrice(ctx, uint16(3), product.ID, startAt, clock.InfinityAt(), nil, nil)
	if err == nil {
		t.FailNow()
	}
	if err != ErrInvalidStartAt {
		t.FailNow()
	}
	selectedCurrent, err := pricerepo.GetPriceByProductID(ctx, product.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if selectedCurrent.ID != createCurrent.ID {
		t.FailNow()
	}

	id, err := pricerepo.GetPriceByProductID(ctx, product.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(id)

}