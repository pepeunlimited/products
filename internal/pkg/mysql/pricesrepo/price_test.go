package pricesrepo

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/productrepo"
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
	created, err := repo.CreateInitialPrice(ctx, price, product.ID, nil, nil)
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
	ended, err := repo.EndPrice(ctx, 2, 12, created.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt,_ := clock.ToMonthDate(2, 12)
	if !ended.EndAt.Equal(endAt) {
		t.FailNow()
	}
}