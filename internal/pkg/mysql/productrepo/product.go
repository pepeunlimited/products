package productrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/ent/plan"
	"github.com/pepeunlimited/prices/internal/pkg/ent/product"
)

var (
	ErrProductSkuExist = errors.New("products: product sku exist")
	ErrProductNotExist = errors.New("products: product not exist")

	ErrCantFigureIsSubscription = errors.New("products: cant figure is subscription")
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, sku string) 						  (*ent.Product, error)
	CreateProducts(ctx context.Context, pageToken int64, pageSize int32)  ([]*ent.Product, int64, error)
	GetProductBySku(ctx context.Context, sku string) 					  (*ent.Product, error)
	GetProductByID(ctx context.Context, withPrices bool, id int) 		  (*ent.Product, error)
	// note: not includes subscribable products
	GetProducts(ctx context.Context, pageToken int64, pageSize int32)     ([]*ent.Product, int64, error)
	Wipe(ctx context.Context)

	IsSubscribableByID(ctx context.Context, id int)						  (*ent.Product, *bool, error)
}

type productMySQL struct {
	client *ent.Client
}

func (mysql productMySQL) CreateProducts(ctx context.Context, pageToken int64, pageSize int32) ([]*ent.Product, int64, error) {
	products, err := mysql.client.Product.Query().Where(
		product.IDGT(int(pageToken))).
		Order(ent.Asc(product.FieldID)).
		Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(products) == 0 {
		return []*ent.Product{}, pageToken, nil
	}
	return products, int64(products[len(products) - 1].ID), nil
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

func (mysql productMySQL) IsSubscribableByID(ctx context.Context, id int) (*ent.Product, *bool, error) {
	selected, err := mysql.GetProductByID(ctx, true, id)
	if err != nil {
		return nil, nil, err
	}
	prices := selected.Edges.Prices
	if len(prices) == 0 {
		return nil, nil, ErrCantFigureIsSubscription
	}
	_, err = prices[0].QueryPlans().Where(plan.HasPrices()).Only(ctx)
	isSubscribable := false
	if err != nil {
		if ent.IsNotFound(err) {
			return selected, &isSubscribable, nil
		}
		return nil, nil, err
	}
	isSubscribable = true
	return selected, &isSubscribable, nil
}

//			https://dbdiagram.io/d
//			**  look up the price columns | plan_prices | iap_source_prices **
//
// subscribable 	=> true : false => [plans] => query prices from external service
// is_iap_source 	=> true : false => [iapsource] => query prices from external service
//

func (mysql productMySQL) GetProductByID(ctx context.Context, withPrices bool, id int) (*ent.Product, error) {
	query := mysql.client.Product.Query().Where(product.ID(id))
	if withPrices {
		query.WithPrices()
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

func (mysql productMySQL) GetProductBySku(ctx context.Context, sku string) (*ent.Product, error) {
	prod, err := mysql.client.Product.Query().Where(product.Sku(sku)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrProductNotExist
		}
		return nil, err
	}
	return prod, nil
}

func (mysql productMySQL) Wipe(ctx context.Context) {
	mysql.client.Subscription.Delete().ExecX(ctx)
	mysql.client.Price.Delete().ExecX(ctx)
	mysql.client.ThirdParty.Delete().ExecX(ctx)
	mysql.client.Plan.Delete().ExecX(ctx)
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

func NewProductRepository(client *ent.Client) ProductRepository {
	return &productMySQL{client:client}
}