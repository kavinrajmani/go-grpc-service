generate: 
	protoc --go_out=product --go_opt=paths=source_relative --go-grpc_out=product --go-grpc_opt=paths=source_relative product.proto