package models

type BookResponse struct {
	Book_id       string `json:"book_id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Yearpublished int32  `json:"yearpublished"`
}

type BookRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Yearpublished int32  `json:"yearpublished"`
}

type BorrowBookRequest struct {
	Book_id string `json:"book_id"`
	User_id string `json:"user_id"`
}
