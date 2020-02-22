package twirp

import (
	"context"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/pkg/rpc/product"
	"testing"
)

func TestProductServer_CreateProduct(t *testing.T) {
	ctx := context.TODO()
	server := NewProductServer(ent.NewEntClient())
	server.products.Wipe(ctx)
	fromServer, err := server.CreateProduct(ctx, &product.CreateProductParams{
		Sku: "sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.GetProduct(ctx, &product.GetProductParams{
		ProductId: fromServer.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.GetProduct(ctx, &product.GetProductParams{
		Sku:"sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	products, err := server.GetProducts(ctx, &product.GetProductsParams{
		PageSize:  20,
		PageToken: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(products.Products) != 1 {
		t.FailNow()
	}
	if products.NextPageToken == 0 {
		t.FailNow()
	}
}
