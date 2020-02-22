##### CreateProduct
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ProductService/CreateProduct" \
-d '{"sku": "unique-sku"}'
```
##### GetProduct
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ProductService/GetProduct" \
-d '{"sku": "unique-sku"}'
```
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ProductService/GetProduct" \
-d '{"product_id": 1}'
```
##### GetProducts
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.ProductService/GetProducts" \
-d '{"page_size": 3, "page_token":0}'
```