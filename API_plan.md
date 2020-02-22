##### CreatePlan (for subscriptions)
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PlanService/CreatePlan" \
-d '{"title_i18n_id": 1, "length": 1, "unit": "days,weeks,months"}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PlanService/CreatePlan" \
-d '{"length": 1, "unit": "days"}'
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
-d '{"show": false}'
```