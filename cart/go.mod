module github.com/sergiosegrera/store/cart

go 1.14

replace github.com/sergiosegrera/store/product => ../product

require (
	github.com/go-chi/chi v4.1.1+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/go-pg/pg/v9 v9.1.6
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.0
	github.com/sergiosegrera/store/product v0.0.0-20200502133723-acca03cee988
	go.uber.org/zap v1.13.0
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.21.0
)
