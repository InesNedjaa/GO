compile:
	protoc api/service1/*.proto \
		--go_out=. --go-grpc_out=. --grpc-gateway_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_opt=paths=source_relative \
		--proto_path=.

	protoc api/service2/*.proto \
		--go_out=. --go-grpc_out=. --grpc-gateway_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_opt=paths=source_relative \
		--proto_path=.

	protoc api/service3/*.proto \
		--go_out=. --go-grpc_out=. --grpc-gateway_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_opt=paths=source_relative \
		--proto_path=.
