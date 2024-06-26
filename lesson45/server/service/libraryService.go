package service

import (
	"context"
	pb "lesson45/protos"
	"lesson45/server/storage"
)

type service struct {
	pb.UnimplementedLibraryServiceServer
	b *storage.Book
}

func NewService(b *storage.Book) *service {
	return &service{b: b}
}

func (s *service) AddBook(ctx context.Context, in *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	res, err := s.b.AddBook(ctx, &pb.AddBookRequest{
		Title:         in.Title,
		Author:        in.Author,
		YearPublished: in.YearPublished,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) SearchBookByTitle(ctx context.Context, in *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	res, err := s.b.SearchBookByTitle(ctx, in.SearchText)
	if err != nil {
		return nil, err
	}
	return &pb.SearchBookResponse{
		Book: res,
	}, nil
}

func (s *service) SearchBookByID(ctx context.Context, in *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	res, err := s.b.SearchBookByID(ctx, in.SearchText)
	if err != nil {
		return nil, err
	}

	return &pb.SearchBookResponse{
		Book: res,
	}, nil
}

func (s *service) BorrowBook(ctx context.Context, in *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	res, err := s.b.BorrowBook(ctx, &pb.BorrowBookRequest{
		BookId: in.BookId,
		UserId: in.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.BorrowBookResponse{Message: res}, nil
}
