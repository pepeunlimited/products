package twirp

import (
	"context"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/pkg/rpc/thirdpartyprice"
	"github.com/twitchtv/twirp"
	"testing"
	"time"
)

func TestThirdPartyServer_CreateThirdParty(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	server := NewThirdPartyPriceServer(client)
	server.thirdpartyprice.Wipe(ctx)
	party, err := server.CreateThirdPartyPrice(ctx, &thirdpartyprice.CreateThirdPartyPriceParams{
		InAppPurchaseSku: "sku",
		StartAtMonth:2,
		StartAtDay:12,
		Type:			  "NON_RENEWING_SUBSCRIPTIONS",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	thirdParty, err := server.GetThirdPartyPrice(ctx, &thirdpartyprice.GetThirdPartyPriceParams{
		ThirdPartyPriceId: party.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if thirdParty.Id != party.Id {
		t.FailNow()
	}
	query,err := server.thirdpartyprice.GetByID(ctx, int(party.Id))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if query.StartAt.Month() != 2 {
		t.FailNow()
	}
	if query.StartAt.Day() != 12 {
		t.FailNow()
	}
	thirdParty, err = server.GetThirdPartyPrice(ctx, &thirdpartyprice.GetThirdPartyPriceParams{
		InAppPurchaseSku: "sku",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if thirdParty.Id != party.Id {
		t.FailNow()
	}
	thirdParty, err = server.GetThirdPartyPrice(ctx, &thirdpartyprice.GetThirdPartyPriceParams{
		GoogleBillingServiceSku:"sku",
	})
	if err == nil {
		t.Error(err)
		t.FailNow()
	}
	if err.(twirp.Error).Code() != twirp.NotFound {
		t.FailNow()
	}
	party, err = server.CreateThirdPartyPrice(ctx, &thirdpartyprice.CreateThirdPartyPriceParams{
		InAppPurchaseSku: "sku2",
		Type:			  "NON_RENEWING_SUBSCRIPTIONS",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	query,_ = server.thirdpartyprice.GetByID(ctx, int(party.Id))
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
	server := NewThirdPartyPriceServer(ent.NewEntClient())
	server.thirdpartyprice.Wipe(ctx)
	_, err := server.CreateThirdPartyPrice(ctx, &thirdpartyprice.CreateThirdPartyPriceParams{
		InAppPurchaseSku: "sku0",
		Type:			  "NON_RENEWING_SUBSCRIPTIONS",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	sku1, err := server.CreateThirdPartyPrice(ctx, &thirdpartyprice.CreateThirdPartyPriceParams{
		InAppPurchaseSku: "sku1",
		Type:			  "NON_RENEWING_SUBSCRIPTIONS",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.CreateThirdPartyPrice(ctx, &thirdpartyprice.CreateThirdPartyPriceParams{
		InAppPurchaseSku: "sku2",
		Type:			  "NON_RENEWING_SUBSCRIPTIONS",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.CreateThirdPartyPrice(ctx, &thirdpartyprice.CreateThirdPartyPriceParams{
		InAppPurchaseSku: "sku3",
		Type:			  "NON_RENEWING_SUBSCRIPTIONS",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	now := time.Now().Add(24 * time.Hour)
	_, err = server.CreateThirdPartyPrice(ctx, &thirdpartyprice.CreateThirdPartyPriceParams{
		InAppPurchaseSku: "sku4",
		StartAtMonth:     int32(now.Month()),
		StartAtDay:       int32(now.Day()),
		Type:			  "NON_RENEWING_SUBSCRIPTIONS",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	now = now.Add(-48 * time.Hour)
	_, err = server.EndThirdPartyPrice(ctx, &thirdpartyprice.EndThirdPartyPriceParams{
		Params: &thirdpartyprice.GetThirdPartyPriceParams{
			ThirdPartyPriceId: sku1.Id,
		},
		EndAtMonth: int32(now.Month()),
		EndAtDay:   int32(now.Day()),
	})

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.thirdpartyprice.GetThirdPartyPrices(ctx, "NON_RENEWING_SUBSCRIPTIONS")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fromServer, err := server.GetThirdPartyPrices(ctx, &thirdpartyprice.GetThirdPartyPricesParams{
		Type: "NON_RENEWING_SUBSCRIPTIONS",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(fromServer.ThirdPartyPrices) != 3 {
		t.Log(len(fromServer.ThirdPartyPrices))
		t.FailNow()
	}
}