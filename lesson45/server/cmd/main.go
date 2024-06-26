package main

import (
	pb "lesson45/protos"
	"lesson45/server/service"
	"lesson45/server/storage"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	db, err := storage.GetDB()
	if err != nil {
		panic(err)
	}

	book := storage.NewBook(db)

	server := service.NewService(book)

	grpc := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpc, server)

	err = grpc.Serve(listener)
	if err != nil {
		panic(err)
	}
}
