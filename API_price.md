##### CreatePrice
```
$ curl -H "Content-Type: application/json" \
  -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PriceService/CreatePrice" \
  -d '{"price": 100, "discount": 100, "product_id":173}'
```
```
$ curl -H "Content-Type: application/json" \
  -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PriceService/CreatePrice" \
  -d '{"price": 100, "discount": 100, "product_id":173, "third_party_id": 1}'
```
```
$ curl -H "Content-Type: application/json" \
  -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PriceService/CreatePrice" \
  -d '{"price": 100, "discount": 100, "product_id":173, "third_party_id": 1, "start_at_day":1, "start_at_month":1}'
```
##### EndPriceAt
```
$ curl -H "Content-Type: application/json" \
  -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PriceService/EndPriceAt" \
  -d '{"params": { "product_id": 0, "price_id": 0, "product_sku": "" }, "end_at_day": 1, "end_at_month":1}'
```
##### GetPrice
```
$ curl -H "Content-Type: application/json" \
  -X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.PriceService/GetPrice" \
  -d '{"product_id": 0, "price_id": 0, "product_sku": "" }'
```