package repository

import (
	"context"
	"fmt"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/ent/task"
)

type TaskRepositoryInterface interface {
	CreateTask(ctx context.Context, task *ent.Task) (*ent.Task, error)
	//ListTask(ctx context.Context, userid int) ([]*ent.Task, error)
	GetTaskById(ctx context.Context, taskid int, userid int) (*ent.Task, error)
	//UpdateTask(ctx context.Context, task *ent.Task, taskid int, userid int) (*ent.Task, error)
	//DeleteTask(ctx context.Context, taskid int, userid int) error
}

type taskRepository struct {
	client *ent.Client
}

func NewTaskRepository(client *ent.Client) TaskRepositoryInterface {
	return &taskRepository{
		client: client,
	}
}

func (r *taskRepository) CreateTask(ctx context.Context, task *ent.Task) (*ent.Task, error) {
	createdTask, err := r.client.Task.Create().
		SetTitle(task.Title).
		SetDescription(task.Description).
		SetStatus(task.Status).
		SetDueDate(*task.DueDate).
		SetUserID(task.UserID).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create task in repository: %w", err)
	}
	return createdTask, nil
}

func (r *taskRepository) GetTaskById(ctx context.Context, id int, userId int) (*ent.Task, error) {
	gotTask, err := r.client.Task.Query().
		Where(task.IDEQ(id)).
		Where(task.UserIDEQ(userId)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to get task by id (%d) in repository: %w", id, err)
	}
	return gotTask, nil
}
