##### StartSubscription
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.SubscriptionService/StartSubscription" \
-d '{"user_id": 1, "plan_id": 1}'
```
##### GetSubscription
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.SubscriptionService/GetSubscription" \
-d '{"user_id": 1, "subscription_id": 1}'
```
##### GetSubscriptions
```
$ curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.products.SubscriptionService/GetSubscriptions" \
-d '{"user_id": 1, "page_size": 20, "page_token":0}'
```