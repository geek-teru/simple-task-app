package repository

import (
	"context"
	"fmt"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/ent/user"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *ent.User) (*ent.User, error)
	GetUserById(ctx context.Context, id int) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	UpdateUser(ctx context.Context, user *ent.User, id int) (*ent.User, error)
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepositoryInterface {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	createdUser, err := r.client.User.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in repository: %w", err)
	}
	return createdUser, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (*ent.User, error) {
	gotUser, err := r.client.User.Query().
		Where(user.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id (%d) in repository: %w", id, err)
	}
	return gotUser, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	gotUser, err := r.client.User.Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email (%s) in repository: %w", email, err)
	}
	return gotUser, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *ent.User, id int) (*ent.User, error) {
	updatedUser, err := r.client.User.UpdateOneID(id).
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update user in repository: %w", err)
	}
	return updatedUser, nil
}
