##### CreateThirdPartyPrice
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyPriceService/CreateThirdPartyPrice" \
-d '{"in_app_purchase_sku": "apple-sku", "start_at_month": 1, "start_at_day": 1, "type": "CONSUMABLE"}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyPriceService/CreateThirdPartyPrice" \
-d '{"in_app_purchase_sku": "apple-sku2", "type": "CONSUMABLE"}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyPriceService/CreateThirdPartyPrice" \
-d '{"in_app_purchase_sku": "apple-sku3", "google_billing_service_sku":"google-billing-sku" "type": "CONSUMABLE"}'
```
##### GetThirdPartyPrice
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyPriceService/GetThirdPartyPrice" \
-d '{"third_party_id": 1}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyPriceService/GetThirdPartyPrice" \
-d '{"in_app_purchase_sku": "sku"}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyPriceService/GetThirdPartyPrice" \
-d '{"google_billing_service_sku": "sku"}'
```
##### GetThirdPartyPrices
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyPriceService/GetThirdPartyPrices" \
-d '{"type": "CONSUMABLE"}'
```
##### EndThirdPartyPrice
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ThirdPartyPriceService/EndThirdPartyPrice" \
-d '{"params": {"third_party_price_id": 91}, "end_at_day": 1, "end_at_month":1}'
```