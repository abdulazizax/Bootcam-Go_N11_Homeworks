DB_URL=postgres://postgres:abdulaziz1221@localhost:5432/imtihon2?sslmode=disable

run:
	go run cmd/main.go
  
migrate_up:
	migrate -path migrations -database ${DB_URL} -verbose up

migrate_down:
	migrate -path migrations -database ${DB_URL} -verbose down

migrate_force:
	migrate -path migrations -database ${DB_URL} -verbose force 3

migrate_file:
	migrate create -ext sql -dir migrations -seq insert_to_book

