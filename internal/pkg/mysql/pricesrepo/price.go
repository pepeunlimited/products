package pricesrepo

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"math"
	"time"
)

type PriceRepository interface {
	CreatePrice(ctx context.Context)

	Wipe(ctx context.Context)
}

type priceMySQL struct {
	client *ent.Client
}

func (p priceMySQL) Wipe(ctx context.Context) {
	p.client.Price.Delete().ExecX(ctx)
}

func (p priceMySQL) CreatePrice(ctx context.Context) {
	p.client.Price.Create().SetStartAt(p.startAt()).SetEndAt(p.endAt()).SetCost(0).SetDiscount(0).SaveX(ctx)
	p.client.Price.Create().SetStartAt(p.startAt()).SetEndAt(p.toMidnight(time.Now())).SetCost(0).SetDiscount(0).SaveX(ctx)
}

func NewPriceRepository(client *ent.Client) PriceRepository {
	return &priceMySQL{client: client}
}

func (p priceMySQL) endAt() time.Time {
	return time.Unix(math.MaxUint32, 0)
}

func (p priceMySQL) startAt() time.Time {
	return time.Unix(0, 0)
}

func (p priceMySQL) toMidnight(t time.Time) time.Time {
	h := time.Duration(t.UTC().Hour())
	s := time.Duration(t.UTC().Second())
	m := time.Duration(t.UTC().Minute())
	return time.Unix(t.Unix(), 0).Add(-h * time.Hour).Add(-m * time.Minute).Add(-s * time.Second).UTC()
}