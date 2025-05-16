migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-create name=<migration_name>"; \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path ./migrations -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" up

proto-generate:
	@if [ -z "$(path)" ]; then \
		echo "Usage: make proto-generate name=<path_to_proto_file>"; \
		exit 1; \
	fi
    protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        $(path)