package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/pkg/thirdpartyrpc"
	"github.com/twitchtv/twirp"
	"testing"
)

func TestThirdPartyServer_CreateThirdParty(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	server := NewThirdPartyServer(client)
	server.thirdparty.Wipe(ctx)
	party, err := server.CreateThirdParty(ctx, &thirdpartyrpc.CreateThirdPartyParams {
		InAppPurchaseSku: "sku",
		StartAtMonth:2,
		StartAtDay:12,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	thirdParty, err := server.GetThirdParty(ctx, &thirdpartyrpc.GetThirdPartyParams{
		Id: party.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if thirdParty.Id != party.Id {
		t.FailNow()
	}
	thirdParty, err = server.GetThirdParty(ctx, &thirdpartyrpc.GetThirdPartyParams{
		InAppPurchaseSku: "sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if thirdParty.Id != party.Id {
		t.FailNow()
	}
	thirdParty, err = server.GetThirdParty(ctx, &thirdpartyrpc.GetThirdPartyParams{
		GoogleBillingServiceSku:"sku",
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Code() != twirp.NotFound {
		t.FailNow()
	}
}

func TestThirdPartyServer_GetThirdParties(t *testing.T) {
	// one is already ended..
	// three exist
	// one in the future..


}