package repository

import (
	"context"
	"fmt"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/ent/user"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *ent.User) error
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	UpdateUser(ctx context.Context, user *ent.User) error
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepositoryInterface {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *ent.User) error {
	_, err := r.client.User.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	user, err := r.client.User.Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email (%s): %w", email, err)
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *ent.User) error {
	_, err := r.client.User.Update().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
