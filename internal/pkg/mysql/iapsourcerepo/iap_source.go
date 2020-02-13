package iapsourcerepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/ent/iapsource"
	"time"
)

var (
	ErrIapSourceNotExist  			= errors.New("iap-source: not exist")
	ErrInAppPurchaseSkuExist 		= errors.New("iap-source: in-app-purchase sku exist")
	ErrGoogleBillingServiceExist 	= errors.New("iap-source: google-billing-service sku exist")
)

type IapSourceRepository interface {
	GetByID(ctx context.Context, id int)					  									(*ent.IapSource, error)
	GetByApple(ctx context.Context, apple string)					  							(*ent.IapSource, error)
	GetByGoogle(ctx context.Context, google string)					  							(*ent.IapSource, error)
	Create(ctx context.Context, inAppPurchaseSku *string, googleBillingServiceSku *string)  	(*ent.IapSource, error)
	EndAt(ctx context.Context, month time.Month, day int, id int)								(*ent.IapSource, error)
	GetSources(ctx context.Context) 										  					([]*ent.IapSource, error)
}

type iapSourceMySQL struct {
	client *ent.Client
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

func (mysql iapSourceMySQL) GetByApple(ctx context.Context, apple string) (*ent.IapSource, error) {
	only, err := mysql.client.IapSource.Query().Where(iapsource.InAppPurchaseSku(apple)).Only(ctx)
}

func (mysql iapSourceMySQL) GetByGoogle(ctx context.Context, google string) (*ent.IapSource, error) {
	panic("implement me")
}

func (mysql iapSourceMySQL) Create(ctx context.Context, inAppPurchaseSku *string, googleBillingServiceSku *string) (*ent.IapSource, error) {
	panic("implement me")
}

func (mysql iapSourceMySQL) EndAt(ctx context.Context, month time.Month, day int, id int) (*ent.IapSource, error) {
	panic("implement me")
}

func (mysql iapSourceMySQL) GetSources(ctx context.Context) ([]*ent.IapSource, error) {
	panic("implement me")
}

func NewIapSourceRepository(client *ent.Client) IapSourceRepository {
	return iapSourceMySQL{client:client}
}
