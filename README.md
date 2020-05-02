# Store
Store is an online e-commerce web application designed for small online stores.

## Installation
Install Docker compose first, then:
```
$ git clone https://github.com/sergiosegrera/store
$ cd store
$ docker-compose up
```

## Structure

```
product/ -- Product service, store front shows available products
product-manager/ -- Product Manager service, administrator interface (add products...)
cart/ -- Cart service, verifies cart contents
checkout/ -- TODO: Checkout service, interacts with Stripe
frontend/ -- TODO: Front end client
```

## TODO
* Finish implementing endpoints for product-manager
* Request input verification
* Resolve database models
* Auth server (gRPC, redis, JWT?)
* Easier way to convert models to proto and back?
