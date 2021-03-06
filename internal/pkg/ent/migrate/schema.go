// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// PlansColumns holds the columns for the "plans" table.
	PlansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title_i18n_id", Type: field.TypeInt64},
		{Name: "length", Type: field.TypeUint8},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "price", Type: field.TypeInt64},
		{Name: "discount", Type: field.TypeInt64},
		{Name: "unit", Type: field.TypeString, Size: 7},
		{Name: "product_plans", Type: field.TypeInt, Nullable: true},
		{Name: "third_party_price_plans", Type: field.TypeInt, Nullable: true},
	}
	// PlansTable holds the schema information for the "plans" table.
	PlansTable = &schema.Table{
		Name:       "plans",
		Columns:    PlansColumns,
		PrimaryKey: []*schema.Column{PlansColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "plans_products_plans",
				Columns: []*schema.Column{PlansColumns[8]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "plans_third_party_prices_plans",
				Columns: []*schema.Column{PlansColumns[9]},

				RefColumns: []*schema.Column{ThirdPartyPricesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PricesColumns holds the columns for the "prices" table.
	PricesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "price", Type: field.TypeInt64},
		{Name: "discount", Type: field.TypeInt64},
		{Name: "product_prices", Type: field.TypeInt, Nullable: true},
		{Name: "third_party_price_prices", Type: field.TypeInt, Nullable: true},
	}
	// PricesTable holds the schema information for the "prices" table.
	PricesTable = &schema.Table{
		Name:       "prices",
		Columns:    PricesColumns,
		PrimaryKey: []*schema.Column{PricesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "prices_products_prices",
				Columns: []*schema.Column{PricesColumns[5]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "prices_third_party_prices_prices",
				Columns: []*schema.Column{PricesColumns[6]},

				RefColumns: []*schema.Column{ThirdPartyPricesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "sku", Type: field.TypeString, Unique: true, Size: 32},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:        "products",
		Columns:     ProductsColumns,
		PrimaryKey:  []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "plan_subscriptions", Type: field.TypeInt, Nullable: true},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "subscriptions_plans_subscriptions",
				Columns: []*schema.Column{SubscriptionsColumns[4]},

				RefColumns: []*schema.Column{PlansColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ThirdPartyPricesColumns holds the columns for the "third_party_prices" table.
	ThirdPartyPricesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "in_app_purchase_sku", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "google_billing_service_sku", Type: field.TypeString, Unique: true, Nullable: true, Size: 32},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString, Unique: true, Size: 28},
	}
	// ThirdPartyPricesTable holds the schema information for the "third_party_prices" table.
	ThirdPartyPricesTable = &schema.Table{
		Name:        "third_party_prices",
		Columns:     ThirdPartyPricesColumns,
		PrimaryKey:  []*schema.Column{ThirdPartyPricesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PlansTable,
		PricesTable,
		ProductsTable,
		SubscriptionsTable,
		ThirdPartyPricesTable,
	}
)

func init() {
	PlansTable.ForeignKeys[0].RefTable = ProductsTable
	PlansTable.ForeignKeys[1].RefTable = ThirdPartyPricesTable
	PricesTable.ForeignKeys[0].RefTable = ProductsTable
	PricesTable.ForeignKeys[1].RefTable = ThirdPartyPricesTable
	SubscriptionsTable.ForeignKeys[0].RefTable = PlansTable
}
