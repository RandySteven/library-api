run:
	go run .

generate_protoc:
	protoc --go_out=. ./proto/*.proto

gen_protoc_grpc:
	protoc --go-grpc_out=. ./proto/*.proto

gen_protoc:
	protoc -I=./proto --go_out=proto --go_opt=paths=source_relative \
    	--go-grpc_out=proto --go-grpc_opt=paths=source_relative \
    	./proto/*.proto