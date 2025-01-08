package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

func (h *UserHandler) Create(c echo.Context) error {
	// requestのBind
	UserReq := service.UserRequest{}
	if err := c.Bind(UserReq); err != nil {
		err = fmt.Errorf("failed handler.CreateUser: %v", err)
		h.logger.Error("[ERROR] CreateUser", zap.Error(err))
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// Serviceの呼び出し
	userRes, err := h.Service.CreateUser(UserReq)
	if err != nil {
		h.logger.Error("[ERROR] CreateUser", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (h *UserHandler) GetById(c echo.Context) error {
	// パスパラメータからidを取得する
	// TODO: クレームからidを取得するようにする
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = fmt.Errorf("failed handler.GetUser: %v", err)
		h.logger.Error("[ERROR] GetUser", zap.Error(err))
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// JWTのクレームからidを取得
	// id, err := jwtutil.GetClaim(c, "id")
	// if err != nil {
	// 	err = fmt.Errorf("failed handler.GetUser: %v, code: %w", err, apperror.ErrInvalidParams)
	// 	h.logger.Error("[ERROR] GetUser", zap.Error(err))
	// 	return c.JSON(http.StatusBadRequest, "Bad Request")
	// }

	// Serviceの呼び出し
	UserRes, err := h.Service.GetUserById(id)
	if err != nil {
		h.logger.Error("[ERROR] GetUser", zap.Error(err))
		return c.JSON(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, UserRes)
}

func (h *UserHandler) Update(c echo.Context) error {
	// パスパラメータからIDを取得する
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = fmt.Errorf("failed handler.UpdateUser: %v", err)
		h.logger.Error("[ERROR] UpdateUser", zap.Error(err))
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// requestのBind
	UserReq := service.UserRequest{}
	if err := c.Bind(UserReq); err != nil {
		err = fmt.Errorf("failed handler.UpdateUser: %v", err)
		h.logger.Error("[ERROR] UpdateUser", zap.Error(err))
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// Serviceの呼び出し
	userRes, err := h.Service.UpdateUser(UserReq, id)
	if err != nil {
		h.logger.Error("[ERROR] UpdateUser", zap.Error(err))
		return c.JSON(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, userRes)
}
