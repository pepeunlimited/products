package product

import (
	"context"
	"errors"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/internal/pkg/ent/product"
)

var (
	ErrProductSkuExist = errors.New("products: product sku exist")
	ErrProductNotExist = errors.New("products: product not exist")
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, sku string) (*ent.Product, error)
	GetProductBySku(ctx context.Context, withPrices bool, withPlans bool, sku string) (*ent.Product, error)
	GetProductByID(ctx context.Context, withPrices bool, withPlans bool, id int) (*ent.Product, error)
	GetProducts(ctx context.Context, pageToken int64, pageSize int32) ([]*ent.Product, int64, error)
	Wipe(ctx context.Context)
}

type productMySQL struct {
	client *ent.Client
}

func (mysql productMySQL) GetProducts(ctx context.Context, pageToken int64, pageSize int32) ([]*ent.Product, int64, error) {
	all, err := mysql.client.Product.Query().Where(product.IDGT(int(pageToken))).
		Order(ent.Asc(product.FieldID)).
		Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(all) == 0 {
		return []*ent.Product{}, 0, nil
	}
	return all, int64(all[len(all) - 1].ID), nil
}

//			https://dbdiagram.io/d
//			**  look up the price columns | plan_prices | iap_source_prices **
//
// subscribable 	=> true : false => [plans] => query prices from external service
// is_iap_source 	=> true : false => [iapsource] => query prices from external service
//

func (mysql productMySQL) GetProductByID(ctx context.Context, withPrices bool, withPlans bool, id int) (*ent.Product, error) {
	query := mysql.client.Product.Query().Where(product.ID(id))
	if withPrices {
		query.WithPrices()
	}
	if withPlans {
		query.WithPlans()
	}
	prod, err := query.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrProductNotExist
		}
		return nil, err
	}
	return prod, nil
}

func (mysql productMySQL) GetProductBySku(ctx context.Context, withPrices bool, withPlans bool, sku string) (*ent.Product, error) {
	query := mysql.client.Product.Query().Where(product.Sku(sku))
	if withPrices {
		query.WithPrices()
	}
	if withPlans {
		query.WithPlans()
	}
	product, err := query.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrProductNotExist
		}
		return nil, err
	}
	return product, nil
}

func (mysql productMySQL) Wipe(ctx context.Context) {
	mysql.client.Subscription.Delete().ExecX(ctx)
	mysql.client.Price.Delete().ExecX(ctx)
	mysql.client.Plan.Delete().ExecX(ctx)
	mysql.client.ThirdPartyPrice.Delete().ExecX(ctx)
	mysql.client.Product.Delete().ExecX(ctx)
}

func (mysql productMySQL) CreateProduct(ctx context.Context, sku string) (*ent.Product, error) {
	products, err := mysql.client.Product.Create().SetSku(sku).Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, ErrProductSkuExist
		}
		return nil, err
	}
	return products, nil
}

func New(client *ent.Client) ProductRepository {
	return &productMySQL{client:client}
}