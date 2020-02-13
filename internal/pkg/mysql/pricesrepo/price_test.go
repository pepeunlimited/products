package pricesrepo

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"testing"
)

func TestPriceMySQL_CreatePrice(t *testing.T) {
	ctx := context.TODO()
	repo := NewPriceRepository(ent.NewEntClient())
	repo.Wipe(ctx)
	price := uint16(0)
	sku := "sku"
	created, err := repo.CreateInitialPrice(ctx, price, &sku)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !created.EndAt.Equal(repo.(*priceMySQL).infinityAt()) {
		t.FailNow()
	}
	if !created.StartAt.Equal(repo.(*priceMySQL).zeroAt()) {
		t.FailNow()
	}
	ended, err := repo.EndPrice(ctx, 2, 12, created.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	endAt,_ := repo.(*priceMySQL).toMonthDate(2, 12)
	if !ended.EndAt.Equal(endAt) {
		t.FailNow()
	}
}