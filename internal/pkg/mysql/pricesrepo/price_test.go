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
	repo.CreatePrice(ctx)
}
