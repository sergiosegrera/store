#! /bin/bash
export AUTH_HTTP_PORT=8085
export AUTH_GRPC_PORT=8001
export JWT_KEY=verysecretmuchwow
export ADMIN_PASSWORD='$2a$04$OsCCq919prxkMC1I5zInYODnx4Y3Cc2sGv.HmYysC/Mc.qa1Ffxii'
echo "Building $(basename $PWD)"
go build -o service.o ./main.go && ./service.o
