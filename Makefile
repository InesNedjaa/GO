PROTO_DIRS = api/proxy_service api/monitoring_mgmt api/power_mgmt api/script_mgmt 

compile:
	@for dir in $(PROTO_DIRS); do \
		protoc $$dir/*.proto \
			--go_out=. \
			--go-grpc_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_opt=paths=source_relative \
			--proto_path=.; \
	done

