package repository

import (
	"context"
	//"fmt"

	"github.com/geek-teru/simple-task-app/ent"
	//"github.com/geek-teru/simple-task-app/ent/task"
)

type TaskRepositoryInterface interface {
	CreateTask(ctx context.Context, task *ent.Task) (*ent.Task, error)
	ListTask(ctx context.Context, userid int) ([]*ent.Task, error)
	GetTask(ctx context.Context, taskid int, userid int) (*ent.Task, error)
	UpdateTask(ctx context.Context, task *ent.Task, taskid int, userid int) (*ent.Task, error)
	DeleteTask(ctx context.Context, taskid int, userid int) error
}
