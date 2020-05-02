#! /usr/bin/bash

protoc cart.proto --go_out=plugins=grpc:.
