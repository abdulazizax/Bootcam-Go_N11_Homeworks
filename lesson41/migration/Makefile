DB_URL=postgres://postgres:abdulaziz1221@localhost:5432/n11?sslmode=disable

run:
	go run main.go
  
migrate_down:
	migrate -database ${DB_URL} -path migrations down

migrate_force:
	migrate -database ${DB_URL} -path migrations force 1

migrate_up:
	migrate -database ${DB_URL} -path migrations up

migrate_file:
	migrate create -ext sql -dir migrations -seq postgres
