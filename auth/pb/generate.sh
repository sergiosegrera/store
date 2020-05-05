#! /bin/bash

BASE=$(basename $1 .proto)

protoc --gofast_out=plugins=grpc:. $1 && protoc-go-inject-tag -input="./$BASE.pb.go"
