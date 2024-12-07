PROTO_PATH ?= pkg/proto
OUTPUT_PATH ?= ./pkg/gen/go
PROTO_FILE ?= organization/organization.proto

proto:
	@protoc -I $(PROTO_PATH) $(PROTO_PATH)/$(PROTO_FILE) --go_out=$(OUTPUT_PATH) --go_opt=paths=source_relative --go-grpc_out=$(OUTPUT_PATH) --go-grpc_opt=paths=source_relative
