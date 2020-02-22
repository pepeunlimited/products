package product

import (
	"context"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"testing"
)

func TestProductMySQL_CreateProduct(t *testing.T) {
	ctx := context.TODO()
	repository := New(ent.NewEntClient())
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
	_, err = repository.GetProductBySku(ctx, false, false, product.Sku)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = repository.GetProductByID(ctx, false, false, product.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	products, nextPageToken, err := repository.GetProducts(ctx, 0, 20)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(products) != 1 {
		t.FailNow()
	}
	if nextPageToken == 0 {
		t.FailNow()
	}
}