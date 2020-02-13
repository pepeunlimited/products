package productrepo

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"testing"
)

func TestProductMySQL_CreateProduct(t *testing.T) {
	ctx := context.TODO()
	repository := NewProductRepository(ent.NewEntClient())
	repository.Wipe(ctx)
	sku := "sku"
	product, err := repository.CreateProduct(ctx, sku)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if product.Sku != sku {
		t.FailNow()
	}
	_, err = repository.GetProductBySku(ctx, product.Sku)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = repository.GetProductByID(ctx, product.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}