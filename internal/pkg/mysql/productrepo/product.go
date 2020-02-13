package productrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/ent/product"
)

var (
	ErrProductSkuExist = errors.New("products: product sku exist")
	ErrProductNotExist = errors.New("products: product not exist")
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, sku string) 		(*ent.Product, error)

	GetProductBySku(ctx context.Context, sku string) 	(*ent.Product, error)
	GetProductByID(ctx context.Context, id int) 		(*ent.Product, error)

	Wipe(ctx context.Context)
}

type productMySQL struct {
	client *ent.Client
}

func (mysql productMySQL) GetProductByID(ctx context.Context, id int) (*ent.Product, error) {
	prod, err := mysql.client.Product.Get(ctx, id)
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
	mysql.client.Price.Delete().ExecX(ctx)
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