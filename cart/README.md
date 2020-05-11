# Cart
Cart parsing service. Verifies cart items against database.
## HTTP
```
POST /cart
```
## GRPC
Use the client in `clients/grpc/` exposes the `/cart` endpoint
## Environment Variables
```
DB_ADDRESS
CART_HTTP_PORT
CART_GRPC_PORT
```
## TODO
* DB connection or no row error handling
* Max cart size
