# Auth
JWT authentication service.  
HTTP transport for interacting with the frontend.  
gRPC transport for service authentication.  

## HTTP
```
POST /login {"password": "verysecuremuchwow"}
POST /refresh
```

## GRPC
Use the client in `client/` exposes the `/check` endpoint

## Environment Variables
```
AUTH_HTTP_PORT -- HTTP port
AUTH_GRPC_PORT -- GRPC port
JWT_KEY -- JWT key (very sensitive!!)
ADMIN_PASSWORD -- Bcrypt hashed administrator password
```

## TODO
* `POST /password` password change endpoint
