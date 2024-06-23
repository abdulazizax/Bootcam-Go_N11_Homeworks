package postgres

import (
	"context"
	m "user/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserrepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user m.UserRequest) (m.UserResponse, error) {
	query := `
		INSERT INTO users (name, phone, age) VALUES ($1, $2, $3) RETURNING user_id, name, phone, age, created_at, updated_at
	`

	var ResUser m.UserResponse

	row := u.DB.QueryRowxContext(ctx, query, user.Name, user.Phone, user.Age)
	err := row.Scan(&ResUser.UserID, &ResUser.Name, &ResUser.Phone, &ResUser.Age, &ResUser.CreatedAt, &ResUser.UpdatedAt)
	if err != nil {
		return m.UserResponse{}, err
	}

	return ResUser, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, id string) (m.UserResponse, error) {
	query := `
		SELECT user_id, name, phone, age, created_at, updated_at FROM users WHERE user_id = $1
	`

	var ResUser m.UserResponse

	row := u.DB.QueryRowxContext(ctx, query, id)
	err := row.Scan(&ResUser.UserID, &ResUser.Name, &ResUser.Phone, &ResUser.Age, &ResUser.CreatedAt, &ResUser.UpdatedAt)
	if err != nil {
		return m.UserResponse{}, err
	}

	return ResUser, nil
}

func (u *UserRepo) GetUsers(ctx context.Context) ([]m.UserResponse, error) {
	query := `
		SELECT user_id, name, phone, age, created_at, updated_at FROM users
	`

	var ResUsers []m.UserResponse

	rows, err := u.DB.QueryxContext(ctx, query)
	if err != nil {
		return []m.UserResponse{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var ResUser m.UserResponse
		err := rows.Scan(&ResUser.UserID, &ResUser.Name, &ResUser.Phone, &ResUser.Age, &ResUser.CreatedAt, &ResUser.UpdatedAt)
		if err != nil {
			return []m.UserResponse{}, err
		}
		ResUsers = append(ResUsers, ResUser)
	}

	err = rows.Err()
	if err != nil {
		return []m.UserResponse{}, err
	}

	return ResUsers, nil
}

func (u *UserRepo) UpdateUserById(ctx context.Context, id string, user m.UserRequest) (m.UserResponse, error) {
	query := `
		UPDATE users SET name = $1, phone = $2, age = $3 WHERE user_id = $4 RETURNING user_id, name, phone, age, created_at, updated_at
	`

	var ResUser m.UserResponse

	row := u.DB.QueryRowxContext(ctx, query, user.Name, user.Phone, user.Age)
	err := row.Scan(&ResUser.UserID, &ResUser.Name, &ResUser.Phone, &ResUser.Age, &ResUser.CreatedAt, &ResUser.UpdatedAt)
	if err != nil {
		return m.UserResponse{}, err
	}

	return ResUser, nil
}

func (u *UserRepo) DeleteUserByID(ctx context.Context, id string) error {
	query := `
		DELETE FROM users WHERE user_id = $1
	`

	_, err := u.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) GetUserByName(ctx context.Context, name string) (m.UserResponse, error) {
	query := `
		SELECT user_id, name, phone, age, created_at, updated_at FROM users WHERE name = $1
	`

	var ResUser m.UserResponse

	row := u.DB.QueryRowxContext(ctx, query, name)
	err := row.Scan(&ResUser.UserID, &ResUser.Name, &ResUser.Phone, &ResUser.Age, &ResUser.CreatedAt, &ResUser.UpdatedAt)
	if err != nil {
		return m.UserResponse{}, err
	}

	return ResUser, nil	
}
