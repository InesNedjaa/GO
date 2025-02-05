PROTO_DIRS = api/proxy_service

compile:
	@for dir in $(PROTO_DIRS); do \
		protoc $$dir/*.proto \
			--go_out=. \
			--go-grpc_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_opt=paths=source_relative \
			--proto_path=.; \
	done

