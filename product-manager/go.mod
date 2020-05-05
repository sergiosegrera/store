module github.com/sergiosegrera/store/product-manager

go 1.14

replace github.com/sergiosegrera/store/auth => ../auth

require (
	github.com/go-chi/chi v4.1.1+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/go-pg/pg/v9 v9.1.6
	github.com/sergiosegrera/store/auth v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.29.1
)
