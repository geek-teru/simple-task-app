package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	SignUp(UserReq *UserRequest) (UserResponse, error)
	SignIn(UserReq *UserRequest) (string, error)
	GetUserProfile(id int) (UserResponse, error)
	UpdateUserProfile(UserReq *UserRequest, id int) (UserResponse, error)
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

func (u *UserService) SignUp(userReq *UserRequest) (UserResponse, error) {
	// passwordをbcryptで暗号化
	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		return UserResponse{}, fmt.Errorf("[ERROR] failed to SignUp in service: %w", err)
	}

	passwordEncryptedUser := &ent.User{Name: userReq.Name, Email: userReq.Email, Password: string(hash)}
	createdUser, err := u.repo.CreateUser(context.Background(), passwordEncryptedUser)
	if err != nil {
		return UserResponse{}, fmt.Errorf("%w", err)
	}

	userRes := UserResponse{
		ID:   int(createdUser.ID),
		Name: createdUser.Name,
	}

	return userRes, nil
}

func (u *UserService) SignIn(userReq *UserRequest) (string, error) {
	storedUser, err := u.repo.GetUserByEmail(context.Background(), userReq.Email)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(userReq.Password))
	if err != nil {
		return "", fmt.Errorf("[ERROR] failed to SignIn in service: %w", err)
	}

	// // JWTトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", fmt.Errorf("[ERROR] failed to SignIn in service: %w", err)
	}

	return tokenString, nil
}

func (u *UserService) GetUserProfile(id int) (UserResponse, error) {
	storedUser, err := u.repo.GetUserById(context.Background(), id)
	if err != nil {
		return UserResponse{}, fmt.Errorf("%w", err)
	}

	userRes := UserResponse{
		ID:   int(storedUser.ID),
		Name: storedUser.Name,
	}

	return userRes, nil
}

func (u *UserService) UpdateUserProfile(UserReq *UserRequest, id int) (UserResponse, error) {
	user := &ent.User{Name: UserReq.Name, Email: UserReq.Email, Password: UserReq.Password}
	updatedUser, err := u.repo.UpdateUser(context.Background(), user, id)
	if err != nil {
		return UserResponse{}, fmt.Errorf("%w", err)
	}

	userRes := UserResponse{
		ID:   int(updatedUser.ID),
		Name: updatedUser.Name,
	}

	return userRes, nil
}
