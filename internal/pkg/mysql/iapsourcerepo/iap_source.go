package iapsourcerepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/prices/internal/pkg/clock"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/ent/iapsource"
	"time"
)

var (
	ErrIapSourceNotExist  				= errors.New("iap-source: not exist")
	ErrInAppPurchaseSkuExist 			= errors.New("iap-source: in-app-purchase sku exist")
	ErrGoogleBillingServiceSkuExist 	= errors.New("iap-source: google-billing-service sku exist")
)

type IapSourceRepository interface {
	GetByID(ctx context.Context, id int)					  									(*ent.IapSource, error)
	GetByIAPSku(ctx context.Context, apple string)					  							(*ent.IapSource, error)
	GetByBillingSku(ctx context.Context, google string)					  						(*ent.IapSource, error)
	Create(ctx context.Context, iapsku string, billingsku *string)  							(*ent.IapSource, error)
	EndAt(ctx context.Context, month time.Month, day int, id int)								(*ent.IapSource, error)
	GetSources(ctx context.Context) 										  					([]*ent.IapSource, error)
	GetSourcesByTime(ctx context.Context, now time.Time) 										([]*ent.IapSource, error)
	Wipe(ctx context.Context)
}

type iapSourceMySQL struct {
	client *ent.Client
}

func (mysql iapSourceMySQL) GetSourcesByTime(ctx context.Context, now time.Time) ([]*ent.IapSource, error) {
	all, err := mysql.client.IapSource.Query().Where(iapsource.StartAtLTE(now), iapsource.EndAtGTE(now)).All(ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (mysql iapSourceMySQL) Wipe(ctx context.Context) {
	mysql.client.Subscription.Delete().ExecX(ctx)
	mysql.client.Price.Delete().ExecX(ctx)
	mysql.client.IapSource.Delete().ExecX(ctx)
	mysql.client.Plan.Delete().ExecX(ctx)
	mysql.client.Product.Delete().ExecX(ctx)
}

func (mysql iapSourceMySQL) GetByID(ctx context.Context, id int) (*ent.IapSource, error) {
	sources, err := mysql.client.IapSource.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrIapSourceNotExist
		}
		return nil, err
	}
	return sources, nil
}

func (mysql iapSourceMySQL) GetByIAPSku(ctx context.Context, iapsku string) (*ent.IapSource, error) {
	sources, err := mysql.client.IapSource.Query().Where(iapsource.InAppPurchaseSku(iapsku)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrIapSourceNotExist
		}
		return nil, err
	}
	return sources, err
}

func (mysql iapSourceMySQL) GetByBillingSku(ctx context.Context, billingsku string) (*ent.IapSource, error) {
	sources, err := mysql.client.IapSource.Query().Where(iapsource.GoogleBillingServiceSku(billingsku)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrIapSourceNotExist
		}
		return nil, err
	}
	return sources, err
}

func (mysql iapSourceMySQL) Create(ctx context.Context, iapsku string, billingsku *string) (*ent.IapSource, error) {
	if _, err := mysql.GetByIAPSku(ctx, iapsku); err == nil {
		return nil, ErrInAppPurchaseSkuExist
	}
	if billingsku != nil {
		if _, err := mysql.GetByBillingSku(ctx, *billingsku); err == nil {
			return nil, ErrGoogleBillingServiceSkuExist
		}
	}
	save, err := mysql.
		client.
		IapSource.
		Create().
		SetEndAt(clock.InfinityAt()).
		SetStartAt(clock.ZeroAt()).
		SetNillableGoogleBillingServiceSku(billingsku).
		SetInAppPurchaseSku(iapsku).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (mysql iapSourceMySQL) EndAt(ctx context.Context, month time.Month, day int, id int) (*ent.IapSource, error) {
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

func (mysql iapSourceMySQL) GetSources(ctx context.Context) ([]*ent.IapSource, error) {
	now := time.Now().UTC()
	return mysql.GetSourcesByTime(ctx, now)
}

func NewIapSourceRepository(client *ent.Client) IapSourceRepository {
	return iapSourceMySQL{client:client}
}
