DB_URL=postgres://postgres:pass@localhost:5432/n11?sslmode=disable

run :
	go run cmd/main.go
  
migrate_up:
	migrate -path server/migrations -database ${DB_URL}  -verbose up

migrate_down:
	migrate -path server/migrations -database ${DB_URL}  -verbose down

migrate_force:
	migrate -path server/migrations -database ${DB_URL}  -verbose force 1

migrate_file:
	migrate create -ext sql -dir server/migrations -seq create_tables

proto-gen:
	protoc --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    protos/libraryService.proto