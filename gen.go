package main

//go:generate protoc --go_out=. --go_opt=module=github.com/Yifangmo/micro-shop-services ./common/proto/common.proto
//go:generate protoc --go_out=./user/proto --go-grpc_out=./user/proto ./user/proto/user.proto
//go:generate protoc --go_out=./goods/proto --go-grpc_out=./goods/proto ./goods/proto/goods.proto
//go:generate protoc --go_out=./inventory/proto --go-grpc_out=./inventory/proto ./inventory/proto/inventory.proto
//go:generate protoc -I=. -I=./order/proto/ --go_out=./order/proto --go-grpc_out=./order/proto goods.proto inventory.proto order.proto
