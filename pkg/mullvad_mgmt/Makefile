PROTO_URL_BASE := https://raw.githubusercontent.com/mullvad/mullvadvpn-app/main
PROTO_URL := $(PROTO_URL_BASE)/mullvad-management-interface/proto/management_interface.proto
GET_PROTOC_GO_RPC = $(shell which protoc-gen-go-grpc)
GET_PROTOC_GO_RPC = $(eval PROTOC_GO_RPC=$(GET_PROTOC_GO_RPC))

all :: dep gen prune

dep ::
	if ! protoc --version 2>&1 > /dev/null; then echo -e \
		"\e[1;31m\n\nmissing protoc, can't generate management gRPC client\e[0m\n"; \
		return 1; \
	fi
	if ! command -v protoc-gen-go 2>&1 > /dev/null; then \
		go install -x google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 || return 1; \
	fi
	if ! command -v protoc-gen-go-grpc 2>&1 > /dev/null; then \
		go install -x google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 || return 1; \
	fi

fetch: proto/management_interface.proto

proto/management_interface.proto:
	mkdir -p proto
	curl "$(PROTO_URL)" -o proto/management_interface.proto
	sed -i '/package mullvad_daemon.management_interface;/a option go_package="git.tcp.direct/kayos/mully/pkg/mullvad_mgmt";' proto/management_interface.proto

management_interface.pb.go: proto/management_interface.proto
	protoc \
		--plugin="$(PROTOC_GO_RPC)" \
		-I ./proto \
		-I /usr/include \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--go_out=. --go_opt=paths=source_relative \
		management_interface.proto

gen: dep fetch management_interface.pb.go

prune ::
	rm -r proto 2>/dev/null || true

clean ::
	rm -r proto 2>/dev/null || true
	rm *.go 2>/dev/null || true
