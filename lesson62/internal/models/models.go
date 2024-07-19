package models

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserRespone struct {
	ID string `json:"id"`
}

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRespone struct {
	ID string `json:"id"`
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type DeleteUserResponse struct {
	ID string `json:"id"`
}
