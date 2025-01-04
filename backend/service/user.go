package service

import (
	"context"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/repository"
)

type UserServiceInterface interface {
	CreateUser(UserReq *UserRequest) (UserResponse, error)
	GetUserById(id int) (UserResponse, error)
	UpdateUser(UserReq *UserRequest, id int) (UserResponse, error)
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
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func NewUserService(repo repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(userReq *UserRequest) (UserResponse, error) {
	user := &ent.User{Name: userReq.Name, Email: userReq.Email, Password: userReq.Password}
	createdUser, err := u.repo.CreateUser(context.Background(), user)
	if err != nil {
		return UserResponse{}, err
	}

	userRes := UserResponse{
		ID:   int(createdUser.ID),
		Name: createdUser.Name,
	}

	return userRes, nil
}

func (u *UserService) GetUserById(id int) (UserResponse, error) {
	gotUser, err := u.repo.GetUserById(context.Background(), id)
	if err != nil {
		return UserResponse{}, err
	}

	userRes := UserResponse{
		ID:   int(gotUser.ID),
		Name: gotUser.Name,
	}

	return userRes, nil
}

func (u *UserService) UpdateUser(UserReq *UserRequest, id int) (UserResponse, error) {
	user := &ent.User{Name: UserReq.Name, Email: UserReq.Email, Password: UserReq.Password}
	updatedUser, err := u.repo.UpdateUser(context.Background(), user, id)
	if err != nil {
		return UserResponse{}, err
	}

	userRes := UserResponse{
		ID:   int(updatedUser.ID),
		Name: updatedUser.Name,
	}

	return userRes, nil
}
