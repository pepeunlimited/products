package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/pkg/productrpc"
	"testing"
)

func TestProductServer_CreateProduct(t *testing.T) {
	ctx := context.TODO()
	server := NewProductServer(ent.NewEntClient())
	server.products.Wipe(ctx)
	product, err := server.CreateProduct(ctx, &productrpc.CreateProductParams{
		Sku: "sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	product, err = server.GetProduct(ctx, &productrpc.GetProductParams{
		ProductId: product.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	product, err = server.GetProduct(ctx, &productrpc.GetProductParams{
		Sku:"sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}