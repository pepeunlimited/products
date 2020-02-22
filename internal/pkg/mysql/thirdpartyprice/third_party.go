package thirdpartyprice

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/products/internal/pkg/clock"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/internal/pkg/ent/thirdpartyprice"
	"time"
)

var (
	ErrThirdPartyPriceNotExist      = errors.New("third-party-price: not exist")
	ErrInAppPurchaseSkuExist        = errors.New("third-party-price: in-app-purchase sku exist")
	ErrGoogleBillingServiceSkuExist = errors.New("third-party-price: google-billing-service sku exist")
)

type ThirdPartyPriceRepository interface {
	GetByID(ctx context.Context, id int)					  									(*ent.ThirdPartyPrice, error)
	GetInAppPurchaseBySku(ctx context.Context, sku string)										(*ent.ThirdPartyPrice, error)
	GetBillingBySku(ctx context.Context, sku string)					  						(*ent.ThirdPartyPrice, error)
	Create(ctx context.Context, apple string, google *string, types Types)  					(*ent.ThirdPartyPrice, error)
	CreateStartAt(ctx context.Context, apple string, google *string, startAt time.Time, types Types) (*ent.ThirdPartyPrice, error)
	EndAt(ctx context.Context, month time.Month, day int, id int)								(*ent.ThirdPartyPrice, error)
	GetThirdPartyPrices(ctx context.Context) 										  			([]*ent.ThirdPartyPrice, error)
	GetThirdPartyPricesByTime(ctx context.Context, now time.Time) 								([]*ent.ThirdPartyPrice, error)
	Wipe(ctx context.Context)
}

type price3rd struct {
	client *ent.Client
}

func (mysql price3rd) CreateStartAt(ctx context.Context, apple string, google *string, startAt time.Time, types Types) (*ent.ThirdPartyPrice, error) {
	if _, err := mysql.GetInAppPurchaseBySku(ctx, apple); err == nil {
		return nil, ErrInAppPurchaseSkuExist
	}
	if google != nil {
		if validator.IsEmpty(*google) {
			google = nil
		} else {
			if _, err := mysql.GetBillingBySku(ctx, *google); err == nil {
				return nil, ErrGoogleBillingServiceSkuExist
			}
		}
	}
	save, err := mysql.
		client.
		ThirdPartyPrice.
		Create().
		SetEndAt(clock.InfinityAt().UTC()).
		SetStartAt(startAt.UTC()).
		SetNillableGoogleBillingServiceSku(google).
		SetInAppPurchaseSku(apple).
		SetType(types.String()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (mysql price3rd) GetThirdPartyPricesByTime(ctx context.Context, now time.Time) ([]*ent.ThirdPartyPrice, error) {
	query := mysql.client.ThirdPartyPrice.Query().Where(thirdpartyprice.StartAtLTE(now), thirdpartyprice.EndAtGTE(now))
	return query.All(ctx)
}

func (mysql price3rd) Wipe(ctx context.Context) {
	mysql.client.Subscription.Delete().ExecX(ctx)
	mysql.client.Price.Delete().ExecX(ctx)
	mysql.client.Plan.Delete().ExecX(ctx)
	mysql.client.ThirdPartyPrice.Delete().ExecX(ctx)
	mysql.client.Product.Delete().ExecX(ctx)
}

func (mysql price3rd) GetByID(ctx context.Context, id int) (*ent.ThirdPartyPrice, error) {
	sources, err := mysql.client.ThirdPartyPrice.Query().Where(thirdpartyprice.ID(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrThirdPartyPriceNotExist
		}
		return nil, err
	}
	return sources, nil
}

func (mysql price3rd) GetInAppPurchaseBySku(ctx context.Context, sku string) (*ent.ThirdPartyPrice, error) {
	sources, err := mysql.client.ThirdPartyPrice.Query().Where(thirdpartyprice.InAppPurchaseSku(sku)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrThirdPartyPriceNotExist
		}
		return nil, err
	}
	return sources, err
}

func (mysql price3rd) GetBillingBySku(ctx context.Context, sku string) (*ent.ThirdPartyPrice, error) {
	sources, err := mysql.client.ThirdPartyPrice.Query().Where(thirdpartyprice.GoogleBillingServiceSku(sku)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrThirdPartyPriceNotExist
		}
		return nil, err
	}
	return sources, err
}

func (mysql price3rd) Create(ctx context.Context, apple string, google *string, types Types) (*ent.ThirdPartyPrice, error) {
	return mysql.CreateStartAt(ctx, apple, google, clock.ZeroAt(), types)
}

func (mysql price3rd) EndAt(ctx context.Context, month time.Month, day int, id int) (*ent.ThirdPartyPrice, error) {
	sources, err := mysql.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	endAt, err := clock.ToMonthDate(month, day)
	if err != nil {
		return nil, err
	}
	updated, err := sources.Update().SetEndAt(endAt).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (mysql price3rd) GetThirdPartyPrices(ctx context.Context) ([]*ent.ThirdPartyPrice, error) {
	now := time.Now().UTC()
	return mysql.GetThirdPartyPricesByTime(ctx, now)
}

func New(client *ent.Client) ThirdPartyPriceRepository {
	return price3rd{client: client}
}
