package thirdpartyrepo

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"testing"
)

func TestIapSourceMySQL_Create(t *testing.T) {
	ctx := context.TODO()
	repo := NewThirdPartyRepository(ent.NewEntClient())
	repo.Wipe(ctx)
	iapsku := "IPOD2008PINK"
	iapSource, err := repo.Create(ctx, iapsku, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !iapSource.EndAt.Equal(clock.InfinityAt()) {
		t.FailNow()
	}
	if !iapSource.StartAt.Equal(clock.ZeroAt()) {
		t.FailNow()
	}
	selected, err := repo.GetInAppPurchaseBySku(ctx, iapsku)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if selected.ID != iapSource.ID {
		t.FailNow()
	}
	id, err := repo.GetByID(ctx, selected.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if id.ID != selected.ID {
		t.Error(err)
		t.FailNow()
	}
	sources, err := repo.GetThirdParties(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(sources) != 1 {
		t.FailNow()
	}
	_, err = repo.EndAt(ctx, 1, 1, iapSource.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	sources, err = repo.GetThirdParties(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(sources) != 0 {
		t.FailNow()
	}
}