package handler

import (
	"fmt"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
	zap "go.uber.org/zap"

	"github.com/geek-teru/simple-task-app/service"
)

type TaskHandler struct {
	Service service.TaskServiceInterface
	logger  *zap.Logger
	// validator *validator.Validate
}

// func NewTaskHandler(service service.TaskServiceInterface, log *zap.Logger) *TaskHandler {
func NewTaskHandler(service service.TaskServiceInterface, log *zap.Logger) *TaskHandler {
	return &TaskHandler{
		Service: service,
		logger:  log,
		// validator: validator.New(),
	}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	// クレームからidを取得
	user := c.Get("user").(*jwt.Token)
	h.logger.Debug("[Debug] token ", zap.Any("user", user))

	claims := user.Claims.(jwt.MapClaims)
	h.logger.Debug("[Debug] claims ", zap.Any("claims", claims))

	userId := claims["user_id"]
	h.logger.Debug("[Debug] claims ", zap.Any("userId", userId))

	// requestのBind
	TaskReq := &service.TaskRequest{}
	if err := c.Bind(TaskReq); err != nil {
		err = fmt.Errorf("failed handler.CreateTask: %v", err)
		h.logger.Error("[ERROR] CreateTask", zap.Error(err))
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// Serviceの呼び出し
	TaskRes, err := h.Service.CreateTask(TaskReq, int(userId.(float64)))
	if err != nil {
		h.logger.Error("[ERROR] CreateTask", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, TaskRes)
}

func (h *TaskHandler) ListTask(c echo.Context) error {
	// クレームからidを取得
	user := c.Get("user").(*jwt.Token)
	h.logger.Debug("[Debug] token ", zap.Any("user", user))

	claims := user.Claims.(jwt.MapClaims)
	h.logger.Debug("[Debug] claims ", zap.Any("claims", claims))

	userId := claims["user_id"]
	h.logger.Debug("[Debug] claims ", zap.Any("userId", userId))

	// クエリパラメータからページ番号を取得
	page := c.QueryParam("page")

	// Serviceの呼び出し
	tokenString, err := h.Service.ListTask(TaskReq, page)
	if err != nil {
		h.logger.Error("[ERROR] ListTask", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, tokenString)
}

func (h *TaskHandler) GetTaskById(c echo.Context) error {
	// クレームからidを取得
	user := c.Get("user").(*jwt.Token)
	h.logger.Debug("[Debug] token ", zap.Any("user", user))

	claims := user.Claims.(jwt.MapClaims)
	h.logger.Debug("[Debug] claims ", zap.Any("claims", claims))

	userId := claims["user_id"]
	h.logger.Debug("[Debug] claims ", zap.Any("userId", userId))

	// パスパラメータからTaskIdを取得

	// Serviceの呼び出し
	TaskRes, err := h.Service.GetTaskById(int(TaskId.(float64)), int(userId.(float64)))
	if err != nil {
		h.logger.Error("[ERROR] GetTaskById", zap.Error(err))
		return c.JSON(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, TaskRes)
}
