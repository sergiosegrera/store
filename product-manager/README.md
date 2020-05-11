# Product Manager
Store administrator service.  
Requires authentication with the Auth service. Uses gRPC to verify JWT token validity.  
## HTTP
```
GET /products
GET /product/{id}
POST /product
DELETE /product/{id}

POST /option
DELETE /option/{id}
```

## Environment Variables
```
DB_ADDRESS
PRODUCT_MANAGER_HTTP_PORT
AUTH_GRPC_ADDRESS -- Auth service address
```

## TODO
* Auth context timeout?
* Error switch in each handler
* Data validation
