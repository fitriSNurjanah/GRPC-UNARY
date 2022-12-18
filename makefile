GO_WORKSPACE := $(GOPATH)/protoc

protoc:
	protoc --proto_path=protos protos/*.Protoc --go_out=$(GO_WORKSPACE) --go-grpc_out=$(GO_WORKSPACE)
	@echo "Protoc compile selesai"

	//betul
protoc --proto_path=protos protos/*.proto --go_out=. --go-grpc_out=.