run:
	go run .

build:
	go build -o bin/library ./main.go
	./bin/library

generate_protoc:
	protoc --go_out=. ./proto/*.proto

gen_protoc_grpc:
	protoc --go-grpc_out=. ./proto/*.proto

protoc_grpc:
	protoc --go_out=. --go-grpc_out=. ./proto/*.proto

gen_protoc:
	protoc -I=./proto --go_out=proto --go_opt=paths=source_relative \
    	--go-grpc_out=proto --go-grpc_opt=paths=source_relative \
    	./proto/*.proto

test:
	go test ./... -v coverprofile cover.out
	
result:
	go tool cover -html=cover.out

exec_grpc_server:
	go run ./cmd/main.go

exec_grpc_client:
	go run ./cmd/clients/main.go