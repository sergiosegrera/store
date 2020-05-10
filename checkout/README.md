# Checkout
Checkout logic service.  
Requires cart validation with the Cart service.  
Uses the stripe api to handle payments.  
## HTTP
```
POST /checkout -- Start checkout process, returns stripe checkout session
POST /confirm/{id} -- Confirms payment
```
## Environment Variables
```
DB_ADDRESS
CART_GRPC_ADDRESS
STRIPE_SECRET -- Stripe API private key
CHECKOUT_HTTP_PORT
```

## TODO
* Finish confirm endpoint
* Create cancel endpoint
