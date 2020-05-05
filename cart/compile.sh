#! /bin/bash
export DB_ADDRESS="localhost:5432"
export CART_HTTP_PORT=8082
export CART_GRPC_PORT=8000
echo "Building $(basename $PWD)"
go build -o service.o ./main.go && ./service.o
