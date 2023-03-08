OS := linux
LOCAL_BIN:=$(CURDIR)/bin

install-proto-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/envoyproxy/protoc-gen-validate@v0.9.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.0

get-vendor-proto:
		mkdir -p vendor-proto
		@if [ ! -d vendor-proto/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor-proto/googleapis &&\
			mkdir -p  vendor-proto/google/ &&\
			mv vendor-proto/googleapis/google/api vendor-proto/google &&\
			rm -rf vendor-proto/googleapis ;\
		fi
		@if [ ! -d vendor-proto/google/protobuf ]; then\
			git clone https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
			mkdir -p  vendor-proto/google/protobuf &&\
			mv vendor-proto/protobuf/src/google/protobuf/*.proto vendor-proto/google/protobuf &&\
			rm -rf vendor-proto/protobuf ;\
		fi

generate-proto:
	mkdir -p loms/pkg/loms_v1

	protoc -I loms/api -I vendor-proto \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
		--go_out=loms/pkg/loms_v1 --go_opt=paths=source_relative \
		--go-grpc_out=loms/pkg/loms_v1 --go-grpc_opt=paths=source_relative \
		loms/api/loms_service.proto

	mkdir -p products/pkg/products_v1

	protoc -I products/api -I vendor-proto \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
		--go_out=products/pkg/products_v1 --go_opt=paths=source_relative \
		--go-grpc_out=products/pkg/products_v1 --go-grpc_opt=paths=source_relative \
		products/api/product_service.proto

	mkdir -p checkout/pkg/checkout_v1

	protoc -I checkout/api -I vendor-proto \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
		--go_out=checkout/pkg/checkout_v1 --go_opt=paths=source_relative \
		--go-grpc_out=checkout/pkg/checkout_v1 --go-grpc_opt=paths=source_relative \
		checkout/api/checkout_service.proto

prepare-proto: install-proto-deps get-vendor-proto generate-proto

build-all:
	cd checkout && GOOS=$(OS) make build
	cd loms && GOOS=$(OS) make build
	cd notifications && GOOS=$(OS) make build

run-all: build-all
	sudo docker compose up --force-recreate --build

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit
