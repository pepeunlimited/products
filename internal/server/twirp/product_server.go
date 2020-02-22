package twirp

import (
	"context"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/product"
	"github.com/pepeunlimited/prices/internal/server/errorz"
	"github.com/pepeunlimited/prices/internal/server/validator"
	"github.com/pepeunlimited/prices/pkg/rpc/product"
)

type ProductServer struct {
	products product.ProductRepository
	valid    validator.ProductValidator
}

func (server ProductServer) GetProducts(ctx context.Context, params *product.GetProductsParams) (*product.GetProductsResponse, error) {
	err := server.valid.GetProducts(params)
	if err != nil {
		return nil, err
	}
	products, nextPageToken, err := server.products.GetProducts(ctx, params.PageToken, params.PageSize)
	if err != nil {
		return nil, errorz.Product(err)
	}
	return &product.GetProductsResponse{
		Products:      ToProducts(products),
		NextPageToken: nextPageToken,
	}, nil
}

func (server ProductServer) CreateProduct(ctx context.Context, params *product.CreateProductParams) (*product.Product, error) {
	err := server.valid.CreateProduct(params)
	if err != nil {
		return nil, err
	}
	product, err := server.products.CreateProduct(ctx, params.Sku)
	if err != nil {
		return nil, errorz.Product(err)
	}
	return ToProduct(product), nil
}

func (server ProductServer) GetProduct(ctx context.Context, params *product.GetProductParams) (*product.Product, error) {
	err := server.valid.GetProduct(params)
	if err != nil {
		return nil, err
	}
	var product *ent.Product
	if !validator2.IsEmpty(params.Sku) {
		product, err = server.products.GetProductBySku(ctx, params.Sku)
	}
	if params.ProductId != 0 {
		product, err = server.products.GetProductByID(ctx, false, int(params.ProductId))
	}
	if err != nil {
		return nil, errorz.Product(err)
	}
	return ToProduct(product), nil
}

func NewProductServer(client *ent.Client) ProductServer {
	return ProductServer{
		products: product.New(client),
		valid:    validator.NewProductValidator(),
	}
}