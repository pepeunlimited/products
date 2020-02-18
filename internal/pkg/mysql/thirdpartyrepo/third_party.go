package thirdpartyrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/ent/thirdparty"
	"time"
)

var (
	ErrThirdPartyNotExist  				= errors.New("third-party: not exist")
	ErrInAppPurchaseSkuExist 			= errors.New("third-party: in-app-purchase sku exist")
	ErrGoogleBillingServiceSkuExist 	= errors.New("third-party: google-billing-service sku exist")
)

type ThirdPartyRepository interface {
	GetByID(ctx context.Context, id int)					  									(*ent.ThirdParty, error)
	GetInAppPurchaseBySku(ctx context.Context, sku string)					  					(*ent.ThirdParty, error)
	GetBillingBySku(ctx context.Context, sku string)					  						(*ent.ThirdParty, error)
	Create(ctx context.Context, apple string, google *string)  									(*ent.ThirdParty, error)
	CreateStartAt(ctx context.Context, apple string, google *string, startAt time.Time)  		(*ent.ThirdParty, error)
	EndAt(ctx context.Context, month time.Month, day int, id int)								(*ent.ThirdParty, error)
	GetThirdParties(ctx context.Context) 										  				([]*ent.ThirdParty, error)
	GetThirdPartiesByTime(ctx context.Context, now time.Time) 									([]*ent.ThirdParty, error)
	Wipe(ctx context.Context)
}

type thirdpartiesMySQL struct {
	client *ent.Client
}

func (mysql thirdpartiesMySQL) CreateStartAt(ctx context.Context, apple string, google *string, startAt time.Time) (*ent.ThirdParty, error) {
	if _, err := mysql.GetInAppPurchaseBySku(ctx, apple); err == nil {
		return nil, ErrInAppPurchaseSkuExist
	}
	if google != nil {
		if _, err := mysql.GetBillingBySku(ctx, *google); err == nil {
			return nil, ErrGoogleBillingServiceSkuExist
		}
	}
	save, err := mysql.
		client.
		ThirdParty.
		Create().
		SetEndAt(clock.InfinityAt().UTC()).
		SetStartAt(startAt.UTC()).
		SetNillableGoogleBillingServiceSku(google).
		SetInAppPurchaseSku(apple).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (mysql thirdpartiesMySQL) GetThirdPartiesByTime(ctx context.Context, now time.Time) ([]*ent.ThirdParty, error) {
	all, err := mysql.client.ThirdParty.Query().Where(thirdparty.StartAtLTE(now), thirdparty.EndAtGTE(now)).All(ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (mysql thirdpartiesMySQL) Wipe(ctx context.Context) {
	mysql.client.Subscription.Delete().ExecX(ctx)
	mysql.client.Price.Delete().ExecX(ctx)
	mysql.client.ThirdParty.Delete().ExecX(ctx)
	mysql.client.Plan.Delete().ExecX(ctx)
	mysql.client.Product.Delete().ExecX(ctx)
}

func (mysql thirdpartiesMySQL) GetByID(ctx context.Context, id int) (*ent.ThirdParty, error) {
	sources, err := mysql.client.ThirdParty.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrThirdPartyNotExist
		}
		return nil, err
	}
	return sources, nil
}

func (mysql thirdpartiesMySQL) GetInAppPurchaseBySku(ctx context.Context, sku string) (*ent.ThirdParty, error) {
	sources, err := mysql.client.ThirdParty.Query().Where(thirdparty.InAppPurchaseSku(sku)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrThirdPartyNotExist
		}
		return nil, err
	}
	return sources, err
}

func (mysql thirdpartiesMySQL) GetBillingBySku(ctx context.Context, sku string) (*ent.ThirdParty, error) {
	sources, err := mysql.client.ThirdParty.Query().Where(thirdparty.GoogleBillingServiceSku(sku)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrThirdPartyNotExist
		}
		return nil, err
	}
	return sources, err
}

func (mysql thirdpartiesMySQL) Create(ctx context.Context, apple string, google *string) (*ent.ThirdParty, error) {
	now := time.Now()
	return mysql.CreateStartAt(ctx, apple, google, now)
}

func (mysql thirdpartiesMySQL) EndAt(ctx context.Context, month time.Month, day int, id int) (*ent.ThirdParty, error) {
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

func (mysql thirdpartiesMySQL) GetThirdParties(ctx context.Context) ([]*ent.ThirdParty, error) {
	now := time.Now().UTC()
	return mysql.GetThirdPartiesByTime(ctx, now)
}

func NewThirdPartyRepository(client *ent.Client) ThirdPartyRepository {
	return thirdpartiesMySQL{client: client}
}
