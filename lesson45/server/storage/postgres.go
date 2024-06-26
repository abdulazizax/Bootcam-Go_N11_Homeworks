package storage

import (
	"context"
	m "lesson45/server/models"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	db *sqlx.DB
}

func NewBook(db *sqlx.DB) *Book {
	return &Book{db: db}
}

func (b *Book) AddBook(ctx context.Context, book m.BookRequest) (string, error) {
	query := `
		INSERT INTO books
    		(title, author, year_published)
		VALUES
    		($1, $2, $3)
		RETURNING
    		book_id`

	var res string

	row := b.db.QueryRowContext(ctx, query, book.Title, book.Author, book.Yearpublished)
	err := row.Scan(&res)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (b *Book) SearchBookByID(ctx context.Context, book_id string) ([]m.BookResponse, error) {
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

	var bookResponse m.BookResponse

	row := b.db.QueryRowContext(ctx, query, book_id)
	err := row.Scan(&bookResponse.Book_id, &bookResponse.Title, &bookResponse.Author, &bookResponse.Yearpublished)
	if err != nil {
		return []m.BookResponse{}, err
	}

	return []m.BookResponse{bookResponse}, nil
}

func (b *Book) SearchBookByTitle(ctx context.Context, title string) ([]m.BookResponse, error) {
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

	var bookResponse []m.BookResponse

	rows, err := b.db.QueryContext(ctx, query, title)
	if err != nil {
		return []m.BookResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var res m.BookResponse
		err = rows.Scan(&res.Book_id, &res.Title, &res.Author, &res.Yearpublished)
		if err != nil {
			return []m.BookResponse{}, err
		}
		bookResponse = append(bookResponse, res)
	}

	err = rows.Err()
	if err != nil {
		return []m.BookResponse{}, err
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
		return false, err
	}

	return true, nil
}
