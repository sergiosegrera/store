#! /bin/bash
export DB_ADDRESS="localhost:5432"
export PRODUCT_MANAGER_HTTP_PORT=8080
echo "Building $(basename $PWD)"
go build -o service.o ./main.go && ./service.o
