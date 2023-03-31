all:
	protoc --proto_path=. -I.\
			-I./type \
			-I${GOPATH}/src \
			--go_out=./generated \
			--go-grpc_out=./generated \
			${ARGS}.proto



