.PHONY: proto openapi gen

PROTO_DIR=docs/protobuff
PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)
PROTO_GEN_DIR=gen/proto

OPENAPI_FILE=docs/echo/openapi.yaml
OPENAPI_GEN_DIR=gen/echo
OPENAPI_PACKAGE=echo

proto:
	protoc --proto_path=$(PROTO_DIR) --go_out=$(PROTO_GEN_DIR) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_GEN_DIR) --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false $(PROTO_FILES)

openapi:
	oapi-codegen -generate types,server -package $(OPENAPI_PACKAGE) $(OPENAPI_FILE) > $(OPENAPI_GEN_DIR)/api.gen.go

gen: proto openapi