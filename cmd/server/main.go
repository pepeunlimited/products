package main

import (
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/prices/internal/server/twirp"
	"github.com/pepeunlimited/prices/pkg/todorpc"
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
	log.Printf("Starting the PricesServer... version=[%v]", Version)

	ts := todorpc.NewTodoServiceServer(twirp.NewTodoServer(), nil)

	mux := http.NewServeMux()
	mux.Handle(ts.PathPrefix(), middleware.Adapt(ts, headers.Username()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}