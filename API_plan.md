##### CreatePlan (for subscriptions)
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PlanService/CreatePlan" \
-d '{"length": 1, "unit": "days", "price": 1, "discount": 1, "product_id": 1, "third_party_price_id": 0}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PlanService/CreatePlan" \
-d '{"length": 1, "unit": "days", "price": 1, "discount": 1, "product_id": 1, "third_party_price_id": 0, "start_at_day": 1,"start_at_month": 1, "end_at_day": 1,"end_at_month": 1}'
```
##### GetPlan
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PlanService/GetPlan" \
-d '{"plan_id": 1}'
```
##### GetPlans
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PlanService/GetPlans" \
-d '{"product_id": 10}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PlanService/GetPlans" \
-d '{"product_sku": "sku"}'
```
##### EndPlanAt
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PlanService/EndPlanAt" \
-d '{"plan_id": 10, "end_at_day": 1, "end_at_month": 1}'
```