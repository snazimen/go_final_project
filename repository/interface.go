package repository

import (
	"time"

	"github.com/snazimen/go_final_project/model"
)

type Task interface {
	Create(task *model.Task) (int64, error)
	GetTasks() (model.TasksResp, error)
	GetTasksBySearchString(searchString string) (model.TasksResp, error)
	GetTasksByDate(searchDate time.Time) (model.TasksResp, error)
	GetTaskById(id string) (*model.Task, error)
	UpdateTask(task *model.Task) error
	MakeTaskDone(id string, date string) error
	DeleteTask(id string) error
}
