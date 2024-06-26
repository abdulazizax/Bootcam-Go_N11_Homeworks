package storage

import (
	"context"
	pb "lesson45/protos"
	m "lesson45/server/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	db *sqlx.DB
}

func NewBook(db *sqlx.DB) *Book {
	return &Book{db: db}
}

func (b *Book) AddBook(ctx context.Context, book *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	query := `
		INSERT INTO books
    		(title, author, year_published)
		VALUES
    		($1, $2, $3)
		RETURNING
    		book_id`

	var res string

	row := b.db.QueryRowContext(ctx, query, book.Title, book.Author, book.YearPublished)
	err := row.Scan(&res)
	if err != nil {
		log.Println(err)
		return &pb.AddBookResponse{}, err
	}

	return &pb.AddBookResponse{BookId: res}, nil
}

func (b *Book) SearchBookByID(ctx context.Context, book_id string) ([]*pb.Book, error) {
	query := `
		SELECT
			book_id,
			title,
			author,
			year_published
		FROM
			books
		WHERE
			book_id = $1 `

	var bookResponse pb.Book

	row := b.db.QueryRowContext(ctx, query, book_id)
	err := row.Scan(
		&bookResponse.BookId,
		&bookResponse.Title,
		&bookResponse.Author,
		&bookResponse.YearPublished)
	if err != nil {
		log.Println(err)
		return []*pb.Book{}, err
	}

	return []*pb.Book{&bookResponse}, nil
}

func (b *Book) SearchBookByTitle(ctx context.Context, title string) ([]*pb.Book, error) {
	query := `
		SELECT
			book_id,
			title,
			author,
			year_published
		FROM
			books
		WHERE
			title = $1 `

	var bookResponse []*pb.Book

	rows, err := b.db.QueryContext(ctx, query, title)
	if err != nil {
		log.Println(err.Error())
		return []*pb.Book{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var res pb.Book
		err = rows.Scan(&res.BookId, &res.Title, &res.Author, &res.YearPublished)
		if err != nil {
			log.Println(err.Error())
			return []*pb.Book{}, err
		}
		bookResponse = append(bookResponse, &res)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
		return []*pb.Book{}, err
	}

	return bookResponse, nil
}

func (b *Book) BorrowBook(ctx context.Context, borrow m.BorrowBookRequest) (bool, error) {
	query := `
		INSERT INTO rent
			(book_id, user_id)
		VALUES
			($1, $2)`

	_, err := b.db.ExecContext(ctx, query, borrow.Book_id, borrow.User_id)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}
