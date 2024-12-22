package repository

import (
	"context"
	"fmt"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/ent/user"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *ent.User) error
	GetUserById(ctx context.Context, id int) (*ent.User, error)
	UpdateUser(ctx context.Context, user *ent.User, id int) error
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
	err := r.client.User.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create user in repository: %w", err)
	}
	return nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (*ent.User, error) {
	user, err := r.client.User.Query().
		Where(user.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id (%d) in repository: %w", id, err)
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *ent.User, id int) error {
	_, err := r.client.User.UpdateOneID(id).
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update user in repository: %w", err)
	}
	return nil
}
