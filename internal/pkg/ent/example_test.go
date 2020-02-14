// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleIapSource() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the iapsource's edges.
	pr0 := client.Price.
		Create().
		SetStartAt(time.Now()).
		SetEndAt(time.Now()).
		SetPrice(1).
		SetDiscount(1).
		SaveX(ctx)
	log.Println("price created:", pr0)

	// create iapsource vertex with its edges.
	is := client.IapSource.
		Create().
		SetInAppPurchaseSku("string").
		SetGoogleBillingServiceSku("string").
		SetStartAt(time.Now()).
		SetEndAt(time.Now()).
		AddPrices(pr0).
		SaveX(ctx)
	log.Println("iapsource created:", is)

	// query edges.
	pr0, err = is.QueryPrices().First(ctx)
	if err != nil {
		log.Fatalf("failed querying prices: %v", err)
	}
	log.Println("prices found:", pr0)

	// Output:
}
func ExamplePlan() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the plan's edges.
	s0 := client.Subscription.
		Create().
		SetUserID(1).
		SetStartAt(time.Now()).
		SetEndAt(time.Now()).
		SaveX(ctx)
	log.Println("subscription created:", s0)
	pr1 := client.Price.
		Create().
		SetStartAt(time.Now()).
		SetEndAt(time.Now()).
		SetPrice(1).
		SetDiscount(1).
		SaveX(ctx)
	log.Println("price created:", pr1)

	// create plan vertex with its edges.
	pl := client.Plan.
		Create().
		SetTitleI18nID(1).
		SetLength(1).
		SetUnit("string").
		AddSubscriptions(s0).
		AddPrices(pr1).
		SaveX(ctx)
	log.Println("plan created:", pl)

	// query edges.
	s0, err = pl.QuerySubscriptions().First(ctx)
	if err != nil {
		log.Fatalf("failed querying subscriptions: %v", err)
	}
	log.Println("subscriptions found:", s0)

	pr1, err = pl.QueryPrices().First(ctx)
	if err != nil {
		log.Fatalf("failed querying prices: %v", err)
	}
	log.Println("prices found:", pr1)

	// Output:
}
func ExamplePrice() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the price's edges.

	// create price vertex with its edges.
	pr := client.Price.
		Create().
		SetStartAt(time.Now()).
		SetEndAt(time.Now()).
		SetPrice(1).
		SetDiscount(1).
		SaveX(ctx)
	log.Println("price created:", pr)

	// query edges.

	// Output:
}
func ExampleProduct() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the product's edges.
	pr0 := client.Price.
		Create().
		SetStartAt(time.Now()).
		SetEndAt(time.Now()).
		SetPrice(1).
		SetDiscount(1).
		SaveX(ctx)
	log.Println("price created:", pr0)

	// create product vertex with its edges.
	pr := client.Product.
		Create().
		SetSku("string").
		AddPrices(pr0).
		SaveX(ctx)
	log.Println("product created:", pr)

	// query edges.
	pr0, err = pr.QueryPrices().First(ctx)
	if err != nil {
		log.Fatalf("failed querying prices: %v", err)
	}
	log.Println("prices found:", pr0)

	// Output:
}
func ExampleSubscription() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the subscription's edges.

	// create subscription vertex with its edges.
	s := client.Subscription.
		Create().
		SetUserID(1).
		SetStartAt(time.Now()).
		SetEndAt(time.Now()).
		SaveX(ctx)
	log.Println("subscription created:", s)

	// query edges.

	// Output:
}
