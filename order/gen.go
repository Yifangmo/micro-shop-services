package main

//go:generate protoc --go_out=./proto --go-grpc_out=./proto ./proto/order.proto
//go:generate protoc --go_out=./proto --go-grpc_out=./proto ./proto/inventory.proto
//go:generate protoc --go_out=./proto --go-grpc_out=./proto ./proto/goods.proto
