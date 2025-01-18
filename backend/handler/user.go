package handler

import (
	"fmt"
	"net/http"

	jwt "github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	zap "go.uber.org/zap"

	"github.com/geek-teru/simple-task-app/service"
)

type UserHandler struct {
	Service service.UserServiceInterface
	logger  *zap.Logger
	// validator *validator.Validate
}

// func NewUserHandler(service service.UserServiceInterface, log *zap.Logger) *UserHandler {
func NewUserHandler(service service.UserServiceInterface, log *zap.Logger) *UserHandler {
	return &UserHandler{
		Service: service,
		logger:  log,
		// validator: validator.New(),
	}
}

func (h *UserHandler) SignUp(c echo.Context) error {
	// requestのBind
	UserReq := &service.UserRequest{}
	if err := c.Bind(UserReq); err != nil {
		err = fmt.Errorf("failed handler.SignUp: %v", err)
		h.logger.Error("[ERROR] SignUp", zap.Error(err))
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// Serviceの呼び出し
	userRes, err := h.Service.SignUp(UserReq)
	if err != nil {
		h.logger.Error("[ERROR] SignUp", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (h *UserHandler) SignIn(c echo.Context) error {
	// requestのBind
	UserReq := &service.UserRequest{}
	if err := c.Bind(UserReq); err != nil {
		err = fmt.Errorf("failed handler.SignIn: %v", err)
		h.logger.Error("[ERROR] SignUp", zap.Error(err))
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// Serviceの呼び出し
	tokenString, err := h.Service.SignIn(UserReq)
	if err != nil {
		h.logger.Error("[ERROR] SignIn", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, tokenString)
}

func (h *UserHandler) GetUserProfile(c echo.Context) error {
	// クレームからidを取得
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid claims"})
	}

	userIdRaw, exists := claims["user_id"]
	if !exists {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "user_id not found"})
	}

	userId, ok := userIdRaw.(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "user_id must be a int"})
	}

	// Serviceの呼び出し
	UserRes, err := h.Service.GetUserProfile(userId)
	if err != nil {
		h.logger.Error("[ERROR] GetUserProfile", zap.Error(err))
		return c.JSON(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, UserRes)
}

func (h *UserHandler) UpdateUserProfile(c echo.Context) error {
	// クレームからidを取得
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid claims"})
	}

	userIdRaw, exists := claims["user_id"]
	if !exists {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "user_id not found"})
	}

	userId, ok := userIdRaw.(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "user_id must be a int"})
	}

	// requestのBind
	UserReq := &service.UserRequest{}
	if err := c.Bind(UserReq); err != nil {
		err = fmt.Errorf("failed handler.UpdateUserProfile: %v", err)
		h.logger.Error("[ERROR] UpdateUserProfile", zap.Error(err))
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// Serviceの呼び出し
	userRes, err := h.Service.UpdateUserProfile(UserReq, userId)
	if err != nil {
		h.logger.Error("[ERROR] UpdateUserProfile", zap.Error(err))
		return c.JSON(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, userRes)
}
