package service

import (
	"context"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/repository"
)

type UserServiceInterface interface {
	CreateUser(UserReq *UserRequest) error
	GetUserById(id int) (UserResponse, error)
	UpdateUser(UserReq *UserRequest, id int) error
}

type (
	UserService struct {
		repo repository.UserRepositoryInterface
	}

	UserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserResponse struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewUserService(repo repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(UserReq *UserRequest) error {
	user := &ent.User{Name: UserReq.Name, Email: UserReq.Email, Password: UserReq.Password}
	err := u.repo.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetUserById(id int) (UserResponse, error) {
	user, err := u.repo.GetUserById(context.Background(), id)
	if err != nil {
		return UserResponse{}, err
	}

	UserRes := UserResponse{
		ID:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	return UserRes, nil
}

func (u *UserService) UpdateUser(UserReq *UserRequest, id int) error {
	user := &ent.User{Name: UserReq.Name, Email: UserReq.Email, Password: UserReq.Password}
	err := u.repo.UpdateUser(context.Background(), user, id)
	if err != nil {
		return err
	}

	return nil
}
