package service

import (
	"context"
	"fmt"
	"time"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/ent/task"
	"github.com/geek-teru/simple-task-app/repository"
)

type TaskServiceInterface interface {
	CreateTask(ctx context.Context, taskReq *TaskRequest, userid int) (TaskResponse, error)
	//ListTask(ctx context.Context, userid int, offset int, limit int) ([]*ent.Task, error)
	GetTaskById(ctx context.Context, taskid int, userid int) (TaskResponse, error)
	UpdateTask(ctx context.Context, taskReq *TaskRequest, taskid int, userid int) (TaskResponse, error)
	DeleteTask(ctx context.Context, taskid int, userid int) (TaskResponse, error)
}

// TODO: ListTaskを実装
// Todo: CreateTaskとUpdateTaskのuseridの取り扱いは共通にする。
type (
	TaskService struct {
		repo repository.TaskRepositoryInterface
	}

	TaskRequest struct {
		Title       string      `json:"title"`
		Description string      `json:"description"`
		DueDate     *time.Time  `json:"due_date"`
		Status      task.Status `json:"status"`
	}

	TaskResponse struct {
		ID          int         `json:"id"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		DueDate     *time.Time  `json:"due_date"`
		Status      task.Status `json:"status"`
		UserID      int         `json:"user_id"`
	}
)

func NewTaskService(repo repository.TaskRepositoryInterface) TaskServiceInterface {
	return &TaskService{
		repo: repo,
	}
}

func (u *TaskService) CreateTask(ctx context.Context, taskReq *TaskRequest, userid int) (TaskResponse, error) {
	task := &ent.Task{Title: taskReq.Title, Description: taskReq.Description, DueDate: taskReq.DueDate, Status: taskReq.Status, UserID: userid}
	createdTask, err := u.repo.CreateTask(context.Background(), task)
	if err != nil {
		return TaskResponse{}, fmt.Errorf("%w", err)
	}

	TaskRes := TaskResponse{
		ID:          int(createdTask.ID),
		Title:       createdTask.Title,
		Description: createdTask.Description,
		DueDate:     createdTask.DueDate,
		Status:      createdTask.Status,
		UserID:      createdTask.UserID,
	}

	return TaskRes, nil
}

func (u *TaskService) ListTask(ctx context.Context, userid int, offset int, limit int) ([]TaskResponse, error) {
	storedTask, err := u.repo.ListTask(context.Background(), userid, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	TasksRes := []TaskResponse{}
	for _, v := range storedTask {
		t := TaskResponse{
			ID:          int(v.ID),
			Title:       v.Title,
			Description: v.Description,
			DueDate:     v.DueDate,
			Status:      v.Status,
			UserID:      v.UserID,
		}
		TasksRes = append(TasksRes, t)
	}
	return TasksRes, nil
}

func (u *TaskService) GetTaskById(ctx context.Context, taskid int, userid int) (TaskResponse, error) {
	storedTask, err := u.repo.GetTaskById(context.Background(), taskid, userid)
	if err != nil {
		return TaskResponse{}, fmt.Errorf("%w", err)
	}

	TaskRes := TaskResponse{
		ID:          int(storedTask.ID),
		Title:       storedTask.Title,
		Description: storedTask.Description,
		DueDate:     storedTask.DueDate,
		Status:      storedTask.Status,
		UserID:      storedTask.UserID,
	}

	return TaskRes, nil
}

func (u *TaskService) UpdateTask(ctx context.Context, taskReq *TaskRequest, taskid int, userid int) (TaskResponse, error) {
	task := &ent.Task{Title: taskReq.Title, Description: taskReq.Description, DueDate: taskReq.DueDate, Status: taskReq.Status, UserID: userid}
	updatedTask, err := u.repo.UpdateTask(context.Background(), task, taskid, userid)
	if err != nil {
		return TaskResponse{}, fmt.Errorf("%w", err)
	}

	TaskRes := TaskResponse{
		ID:          int(updatedTask.ID),
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		DueDate:     updatedTask.DueDate,
		Status:      updatedTask.Status,
		UserID:      updatedTask.UserID,
	}

	return TaskRes, nil
}

func (u *TaskService) DeleteTask(ctx context.Context, taskid int, userid int) (TaskResponse, error) {
	err := u.repo.DeleteTask(context.Background(), taskid, userid)
	if err != nil {
		return TaskResponse{}, fmt.Errorf("%w", err)
	}

	TaskRes := TaskResponse{}

	return TaskRes, nil
}
