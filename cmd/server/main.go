package main

import (
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/server/twirp"
	"github.com/pepeunlimited/prices/pkg/planrpc"
	"github.com/pepeunlimited/prices/pkg/pricerpc"
	"github.com/pepeunlimited/prices/pkg/productrpc"
	"github.com/pepeunlimited/prices/pkg/subscriptionrpc"
	"github.com/pepeunlimited/prices/pkg/thirdpartypricerpc"
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
	ps := productrpc.NewProductServiceServer(twirp.NewProductServer(client), nil)
	pls := planrpc.NewPlanServiceServer(twirp.NewPlanServer(client), nil)
	ss := subscriptionrpc.NewSubscriptionServiceServer(twirp.NewSubscriptionServer(client), nil)
	tp := thirdpartypricerpc.NewThirdPartyServiceServer(twirp.NewThirdPartyServer(client), nil)
	prs := pricerpc.NewPriceServiceServer(twirp.NewPriceServer(client), nil)

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