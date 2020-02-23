package main

import (
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/internal/server/twirp"
	"github.com/pepeunlimited/products/pkg/rpc/plan"
	"github.com/pepeunlimited/products/pkg/rpc/price"
	"github.com/pepeunlimited/products/pkg/rpc/product"
	"github.com/pepeunlimited/products/pkg/rpc/subscription"
	"github.com/pepeunlimited/products/pkg/rpc/thirdpartyprice"
	"log"
	"net/http"
)

const (
	Version = "0.0.1"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Printf("Starting the ProductsServer... version=[%v]", Version)

	client := ent.NewEntClient()
	ps := product.NewProductServiceServer(twirp.NewProductServer(client), nil)
	pls := plan.NewPlanServiceServer(twirp.NewPlanServer(client), nil)
	ss := subscription.NewSubscriptionServiceServer(twirp.NewSubscriptionServer(client), nil)
	tp := thirdpartyprice.NewThirdPartyPriceServiceServer(twirp.NewThirdPartyPriceServer(client), nil)
	prs := price.NewPriceServiceServer(twirp.NewPriceServer(client), nil)

	mux := http.NewServeMux()
	mux.Handle(ps.PathPrefix(), middleware.Adapt(ps))
	mux.Handle(pls.PathPrefix(), middleware.Adapt(pls))
	mux.Handle(ss.PathPrefix(), middleware.Adapt(ss))
	mux.Handle(tp.PathPrefix(), middleware.Adapt(tp))
	mux.Handle(prs.PathPrefix(), middleware.Adapt(prs))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}