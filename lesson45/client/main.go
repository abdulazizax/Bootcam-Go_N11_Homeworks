package main

import (
	"context"
	pb "lesson45/protos"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	gen := pb.NewLibraryServiceClient(conn)

	id := AddBook(gen)
	SearchBookByTitle(gen)
	SearchBookByID(gen, id)
	BorrowBook(gen, id)

}

func AddBook(gen pb.LibraryServiceClient) string {
	req := &pb.AddBookRequest{
		Title:         "Sariq devni minib",
		Author:        "Xudoyberdi To'xtaboyev",
		YearPublished: 2009,
	}

	res, err := gen.AddBook(context.Background(), req)
	if err != nil {
		log.Println(err)
		return ""
	}

	log.Println(res)
	return res.BookId
}

func SearchBookByID(gen pb.LibraryServiceClient, id string) {
	req := &pb.SearchBookRequest{
		SearchText: id,
	}

	res, err := gen.SearchBookByID(context.Background(), req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}

func SearchBookByTitle(gen pb.LibraryServiceClient) {
	req := &pb.SearchBookRequest{
		SearchText: "Sariq devni minib",
	}

	res, err := gen.SearchBookByTitle(context.Background(), req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}

func BorrowBook(gen pb.LibraryServiceClient, id string) {
	req := &pb.BorrowBookRequest{
		BookId: id,
		UserId: "64c92799-7729-490b-aab8-9566b7d0ed47",
	}
	res, err := gen.BorrowBook(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	log.Println("Message: ", res)
}
