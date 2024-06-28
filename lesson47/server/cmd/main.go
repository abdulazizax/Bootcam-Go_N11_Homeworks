package main

import (
	"fmt"
	"log"
	"net"
	"server/genproto/TransportService/transport"
	"server/genproto/WeatherService/weather"
	serveces "server/services"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func GetDB() (*sqlx.DB, error) {

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"pass",
		"n11",
	)

	db, err := sqlx.Connect("postgres", psqlUrl)
	return db, err
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}

	transportService := serveces.NewTransportRepo(db)
	weatherService := serveces.NewWeatherRepo(db)

	transport.RegisterTransportServiceServer(grpcServer, transportService)
	weather.RegisterWeatherServiceServer(grpcServer, weatherService)

	log.Println("gRPC server is running on port :8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
