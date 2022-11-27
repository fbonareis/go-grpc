PROTO_OUT ?= pkg/pb

# Current dir
DIR ?= $(shell pwd)

# Protoc command dockerized
PROTOC_CMD ?= docker run --rm -v $(DIR):$(DIR) -w $(DIR) thethingsindustries/protoc

.PHONY: compile-proto-go
compile-proto-go: ## Compile the protobuf contracts to generated Go code
	@rm -rf $(PROTO_OUT)
	@mkdir -p $(PROTO_OUT)
	@find proto -name "*.proto" -type f -exec $(PROTOC_CMD) -Iproto/ \
	     --go_opt=paths=source_relative \
	     --go_out=plugins=grpc:$(PROTO_OUT) {} \;