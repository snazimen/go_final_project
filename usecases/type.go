package usecases

import (
	"time"

	"github.com/snazimen/go_final_project/model"
)

type Task interface {
	GetNextDate(now time.Time, date string, repeat string) (string, error)
	CreateTask(task *model.Task, today bool) (*model.TaskResp, error)
	GetTasks(searchString string) (model.TasksResp, error)
	GetTask(id string) (*model.Task, error)
	UpdateTask(task *model.Task, today bool) error
	MakeTaskDone(id string) error
	DeleteTask(id string) error
}
