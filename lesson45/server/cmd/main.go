package main

import (
	"context"
	pb "lesson45/protos"
	m "lesson45/server/models"
	"lesson45/server/storage"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLibraryServiceServer
	b *storage.Book
}

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

	s := server{
		b: book,
	}

	grpc := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpc, &s)

	err = grpc.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func (s *server) AddBook(ctx context.Context, in *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	res, err := s.b.AddBook(ctx, &pb.AddBookRequest{
		Title: in.Title,
		Author: in.Author,
		YearPublished: in.YearPublished,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *server) SearchBookByTitle(ctx context.Context, in *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	res, err := s.b.SearchBookByTitle(ctx, in.SearchText)
	if err != nil {
		return nil, err
	}
	return &pb.SearchBookResponse{
		Book: res,
	}, nil
}

func (s *server) SearchBookByID(ctx context.Context, in *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	res, err := s.b.SearchBookByID(ctx, in.SearchText)
	if err != nil {
		return nil, err
	}

	return &pb.SearchBookResponse{
		Book: res,
	}, nil
}

func (s *server) BorrowBook(ctx context.Context, in *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	req := m.BorrowBookRequest{
		Book_id: in.BookId,
		User_id: in.UserId,
	}
	res, err := s.b.BorrowBook(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.BorrowBookResponse{Message: res}, nil
}
