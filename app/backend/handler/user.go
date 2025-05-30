package handler

import (
	"fmt"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
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

	h.logger.Info("[INFO] SignIn", zap.String("token", tokenString))

	return c.JSON(http.StatusCreated, tokenString)
}

func (h *UserHandler) GetUserProfile(c echo.Context) error {
	// クレームからidを取得
	user := c.Get("user").(*jwt.Token)
	h.logger.Debug("[Debug] token ", zap.Any("user", user))

	claims := user.Claims.(jwt.MapClaims)
	h.logger.Debug("[Debug] claims ", zap.Any("claims", claims))

	userId := claims["user_id"]
	h.logger.Debug("[Debug] claims ", zap.Any("userId", userId))

	// Serviceの呼び出し
	UserRes, err := h.Service.GetUserProfile(int(userId.(float64)))
	if err != nil {
		h.logger.Error("[ERROR] GetUserProfile", zap.Error(err))
		return c.JSON(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, UserRes)
}

func (h *UserHandler) UpdateUserProfile(c echo.Context) error {
	// クレームからidを取得
	user := c.Get("user").(*jwt.Token)
	h.logger.Debug("[Debug] token ", zap.Any("user", user))

	claims := user.Claims.(jwt.MapClaims)
	h.logger.Debug("[Debug] claims ", zap.Any("claims", claims))

	userId := claims["user_id"]
	h.logger.Debug("[Debug] claims ", zap.Any("userId", userId))

	// requestのBind
	UserReq := &service.UserRequest{}
	if err := c.Bind(UserReq); err != nil {
		err = fmt.Errorf("failed handler.UpdateUserProfile: %v", err)
		h.logger.Error("[ERROR] UpdateUserProfile", zap.Error(err))
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// Serviceの呼び出し
	userRes, err := h.Service.UpdateUserProfile(UserReq, int(userId.(float64)))
	if err != nil {
		h.logger.Error("[ERROR] UpdateUserProfile", zap.Error(err))
		return c.JSON(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, userRes)
}
