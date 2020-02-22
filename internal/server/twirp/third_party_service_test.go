package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/pkg/thirdpartypricerpc"
	"github.com/twitchtv/twirp"
	"log"
	"testing"
	"time"
)

func TestThirdPartyServer_CreateThirdParty(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	server := NewThirdPartyServer(client)
	server.thirdparty.Wipe(ctx)
	party, err := server.CreateThirdParty(ctx, &thirdpartypricerpc.CreateThirdPartyParams {
		InAppPurchaseSku: "sku",
		StartAtMonth:2,
		StartAtDay:12,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	thirdParty, err := server.GetThirdParty(ctx, &thirdpartypricerpc.GetThirdPartyParams{
		Id: party.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if thirdParty.Id != party.Id {
		t.FailNow()
	}
	query,_ := server.thirdparty.GetByID(ctx, int(party.Id))
	if query.StartAt.Month() != 2 {
		t.FailNow()
	}
	if query.StartAt.Day() != 12 {
		t.FailNow()
	}
	thirdParty, err = server.GetThirdParty(ctx, &thirdpartypricerpc.GetThirdPartyParams{
		InAppPurchaseSku: "sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if thirdParty.Id != party.Id {
		t.FailNow()
	}
	thirdParty, err = server.GetThirdParty(ctx, &thirdpartypricerpc.GetThirdPartyParams{
		GoogleBillingServiceSku:"sku",
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Code() != twirp.NotFound {
		t.FailNow()
	}
	party, err = server.CreateThirdParty(ctx, &thirdpartypricerpc.CreateThirdPartyParams {
		InAppPurchaseSku: "sku2",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	query,_ = server.thirdparty.GetByID(ctx, int(party.Id))
	if query.StartAt.Month() != 1 {
		t.FailNow()
	}
	if query.StartAt.Day() != 1 {
		t.FailNow()
	}
	if query.StartAt.Year() != 1970 {
		t.FailNow()
	}
}

func TestThirdPartyServer_GetThirdParties(t *testing.T) {
	// one is already ended.. sku1
	// three exist sku0,sku2,sku3
	// one in the future.. sku4
	ctx := context.TODO()
	server := NewThirdPartyServer(ent.NewEntClient())
	server.thirdparty.Wipe(ctx)
	_, err := server.CreateThirdParty(ctx, &thirdpartypricerpc.CreateThirdPartyParams{
		InAppPurchaseSku: "sku0",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	sku1, err := server.CreateThirdParty(ctx, &thirdpartypricerpc.CreateThirdPartyParams{
		InAppPurchaseSku: "sku1",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.CreateThirdParty(ctx, &thirdpartypricerpc.CreateThirdPartyParams{
		InAppPurchaseSku: "sku2",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.CreateThirdParty(ctx, &thirdpartypricerpc.CreateThirdPartyParams{
		InAppPurchaseSku: "sku3",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	now := time.Now().Add(24 * time.Hour)
	_, err = server.CreateThirdParty(ctx, &thirdpartypricerpc.CreateThirdPartyParams{
		InAppPurchaseSku: "sku4",
		StartAtMonth:     int32(now.Month()),
		StartAtDay:       int32(now.Day()),
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	now = now.Add(-24 * time.Hour)
	_, err = server.EndThirdParty(ctx, &thirdpartypricerpc.EndThirdPartyParams{
		Params: &thirdpartypricerpc.GetThirdPartyParams{
			Id: sku1.Id,
		},
		EndAtMonth: int32(now.Month()),
		EndAtDay:   int32(now.Day()),
	})

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromDB, err := server.thirdparty.GetThirdParties(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err := server.GetThirdParties(ctx, &thirdpartypricerpc.GetThirdPartiesParams{})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(fromDB)
	log.Print(fromServer)
}