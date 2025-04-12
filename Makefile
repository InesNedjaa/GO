compile:
	protoc api/monitoring_mgmt/*.proto \
		--go_out=. --go-grpc_out=. --grpc-gateway_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_opt=paths=source_relative \
		--proto_path=.

	protoc  api/power_mgmt/*.proto \
		--go_out=. --go-grpc_out=. --grpc-gateway_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_opt=paths=source_relative \
		--proto_path=.

	protoc api/script_mgmt/*.proto \
		--go_out=. --go-grpc_out=. --grpc-gateway_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_opt=paths=source_relative \
		--proto_path=.

	protoc api/proxy_service/*.proto \
		--go_out=. --go-grpc_out=. --grpc-gateway_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_opt=paths=source_relative \
		--proto_path=.	