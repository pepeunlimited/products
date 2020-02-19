package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/pkg/pricerpc"
	"log"
	"testing"
)

func TestPriceServer_CreatePrice(t *testing.T) {
	ctx := context.TODO()
	server := NewPriceServer(ent.NewEntClient())
	server.products.Wipe(ctx)
	product, err := server.products.CreateProduct(ctx, "sku")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	price, err := server.CreatePrice(ctx, &pricerpc.CreatePriceParams{
		StartAtDay:   0,
		StartAtMonth: 0,
		EndAtDay:     0,
		EndAtMonth:   0,
		Price:        3,
		Discount:     3,
		ProductId:    int64(product.ID),
		PlanId:       0,
		ThirdPartyId: 0,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(price)
}
