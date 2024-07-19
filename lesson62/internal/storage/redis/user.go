package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"csb/internal/models"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	rd *redis.Client
}

func NewRedisRepository() *RedisRepository {
	return &RedisRepository{rd: connectDB()}
}

func (r *RedisRepository) CreateUser(ctx context.Context, in *models.CreateUserRequest) (*models.CreateUserRespone, error) {
	id := uuid.New().String()

	userJSON, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user: %w", err)
	}

	key := "user:" + id

	err = r.rd.Set(ctx, key, userJSON, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to set user in Redis: %w", err)
	}

	return &models.CreateUserRespone{ID: id}, nil
}

func (r *RedisRepository) GetUser(ctx context.Context, in *models.GetUserRequest) (*models.GetUserResponse, error) {
	key := "user:" + in.ID

	userJSON, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user from Redis: %w", err)
	}

	var user models.GetUserRequest
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &models.GetUserResponse{
		ID: user.ID,
	}, nil
}

func (r *RedisRepository) UpdateUser(ctx context.Context, in *models.UpdateUserRequest) (*models.UpdateUserRespone, error) {
	key := "user:" + in.ID

	exists, err := r.rd.Exists(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}
	if exists == 0 {
		return nil, fmt.Errorf("user not found")
	}

	userJSON, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user: %w", err)
	}

	err = r.rd.Set(ctx, key, userJSON, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to update user in Redis: %w", err)
	}

	return &models.UpdateUserRespone{ID: in.ID}, nil
}

func (r *RedisRepository) DeleteUser(ctx context.Context, in *models.DeleteUserRequest) (*models.DeleteUserResponse, error) {
	key := "user:" + in.ID

	deleted, err := r.rd.Del(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to delete user from Redis: %w", err)
	}
	if deleted == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &models.DeleteUserResponse{ID: in.ID}, nil
}
