package service

import (
	"context"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	CreateUser(UserReq UserRequest) (UserResponse, error)
	GetUserById(id int) (UserResponse, error)
	UpdateUser(UserReq UserRequest, id int) (UserResponse, error)
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

func (u *UserService) CreateUser(userReq UserRequest) (UserResponse, error) {
	// passwordをbcryptで暗号化
	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		return UserResponse{}, err
	}

	passwordEncryptedUser := &ent.User{Name: userReq.Name, Email: userReq.Email, Password: string(hash)}
	createdUser, err := u.repo.CreateUser(context.Background(), passwordEncryptedUser)
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

func (u *UserService) UpdateUser(UserReq UserRequest, id int) (UserResponse, error) {
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
