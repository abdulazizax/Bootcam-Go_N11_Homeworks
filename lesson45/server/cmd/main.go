package main

import (
	"context"
	pb "lesson45/protos"
	m "lesson45/server/models"
	"lesson45/server/storage"
	"net"

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
	req := m.BookRequest{
		Title:         in.Title,
		Author:        in.Author,
		Yearpublished: in.YearPublished,
	}

	res, err := s.b.AddBook(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AddBookResponse{BookId: res}, nil
}

func (s *server) SearchBook(ctx context.Context, in *pb.SeachBookRequest) (*pb.SeachBookResponse, error) {
	res, err := s.b.SearchBookByTitle(ctx, in.SerchText)
	if err != nil {
		return nil, err
	}
	return &pb.SeachBookResponse{
		Books: res,
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

	return &pb.BorrowBookResponse{
		Success: res,
	}, nil
}
