##### CreateThirdParty
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyService/CreateThirdParty" \
-d '{"in_app_purchase_sku": "apple-sku", "start_at_month": 1, "start_at_day": 1}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyService/CreateThirdParty" \
-d '{"in_app_purchase_sku": "apple-sku2"}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyService/CreateThirdParty" \
-d '{"in_app_purchase_sku": "apple-sku3", "google_billing_service_sku":"google-billing-sku"}'
```
##### GetThirdParty
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyService/GetThirdParty" \
-d '{"third_party_id": 1}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyService/GetThirdParty" \
-d '{"in_app_purchase_sku": "sku"}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyService/GetThirdParty" \
-d '{"google_billing_service_sku": "sku"}'
```
##### GetThirdParties
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyService/GetThirdParties" \
-d '{"show": false}'
```
##### EndThirdParty
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyService/EndThirdParty" \
-d '{"params": {"third_party_id": 91}, "end_at_day": 1, "end_at_month":1}'
```