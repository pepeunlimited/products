##### CreatePrice
```
$ curl -H "Content-Type: application/json" \
  -X POST "localhost:8080/twirp/pepeunlimited.products.PriceService/CreatePrice" \
  -d '{"price": 100, "discount": 100, "product_id":173}'
```
```
$ curl -H "Content-Type: application/json" \
  -X POST "localhost:8080/twirp/pepeunlimited.products.PriceService/CreatePrice" \
  -d '{"price": 100, "discount": 100, "product_id":173, "third_party_id": 1}'
```
```
$ curl -H "Content-Type: application/json" \
  -X POST "localhost:8080/twirp/pepeunlimited.products.PriceService/CreatePrice" \
  -d '{"price": 100, "discount": 100, "product_id":173, "third_party_id": 1, "start_at_day":1, "start_at_month":1}'
```
##### SubscriptionPrice
```
$ curl -H "Content-Type: application/json" \
  -X POST "localhost:8080/twirp/pepeunlimited.products.PriceService/CreatePrice" \
  -d '{"price": 100, "discount": 100, "product_id":173, "start_at_day":1, "start_at_month":1,"end_at_day":2, "end_at_month": 2, "plan_id":1, "third_party_id": 1}'
```

int32   start_at_day     = 1; // optional
    int32   start_at_month   = 2; // optional
    int32   end_at_day       = 3; // optional
    int32   end_at_month     = 4; // optional
    uint32  price           = 5; // => required
    uint32  discount        = 6; // optional
    int64   product_id      = 7; // => required
    int64   plan_id         = 8; // optional
    int32   third_party_id  = 9; // optional
