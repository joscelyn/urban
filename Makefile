# vi: ft=make
.PHONY: proto test

proto:
	go get google.golang.org/protobuf
	protoc -I . urban.proto --lile-server_out=. --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative

test: proto
	go test -p 1 -v ./...
