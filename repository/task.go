package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/snazimen/go_final_project/model"
)

var _ Task = (*TaskRepo)(nil)

type TaskRepo struct {
	Db *sqlx.DB
}

const limit = 50

func NewNewRepository(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{Db: db}
}

func (r *TaskRepo) CreateTask(task *model.Task) (int64, error) {
	res, err := r.Db.Exec(SQLCreateTask, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		log.Debugf("Database.CreateTask: %+v", err)

		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Debugf("Database.CreateTask: %+v", err)

		return 0, err
	}

	return id, nil
}

func (r *TaskRepo) GetTasks() (model.TasksResp, error) {
	tasks := make([]model.Task, 0)

	res, err := r.Db.Query(SQLGetTasks, time.Now().Format(model.TimeFormat), limit)
	if err != nil {
		log.Debugf("Database.GetTasks: %+v", err)

		return model.TasksResp{Tasks: tasks}, err
	}

	defer res.Close()

	var task model.Task

	for res.Next() {
		err = res.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			log.Debugf("Database.GetTasks: %+v", err)

			return model.TasksResp{Tasks: tasks}, err
		}

		tasks = append(tasks, task)
	}

	if err = res.Err(); err != nil {
		return model.TasksResp{Tasks: tasks}, err
	}

	return model.TasksResp{Tasks: tasks}, nil
}

func (r *TaskRepo) GetTasksBySearchString(searchString string) (model.TasksResp, error) {
	tasks := make([]model.Task, 0)

	res, err := r.Db.Query(SQLGetTasksBySearchString, "%"+searchString+"%", limit)
	if err != nil {
		log.Debugf("Database.GetTasksBySearchString: %+v", err)

		return model.TasksResp{Tasks: tasks}, err
	}

	defer res.Close()

	var task model.Task

	for res.Next() {
		err = res.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			log.Debugf("Database.GetTasksBySearchString: %+v", err)

			return model.TasksResp{Tasks: tasks}, err
		}

		tasks = append(tasks, task)
	}

	if err = res.Err(); err != nil {
		return model.TasksResp{Tasks: tasks}, err
	}

	return model.TasksResp{Tasks: tasks}, nil
}

func (r *TaskRepo) GetTasksByDate(searchDate time.Time) (model.TasksResp, error) {
	tasks := make([]model.Task, 0)

	res, err := r.Db.Query(SQLGetTasksByDate, searchDate.Format(model.TimeFormat), limit)
	if err != nil {
		log.Debugf("Database.GetTasksByDate: %+v", err)

		return model.TasksResp{Tasks: tasks}, err
	}

	defer res.Close()

	var task model.Task

	for res.Next() {
		err = res.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			log.Debugf("Database.GetTasksByDate: %+v", err)

			return model.TasksResp{Tasks: tasks}, err
		}

		tasks = append(tasks, task)
	}

	if err = res.Err(); err != nil {
		return model.TasksResp{Tasks: tasks}, err
	}

	return model.TasksResp{Tasks: tasks}, nil
}

func (r *TaskRepo) GetTaskById(id string) (*model.Task, error) {
	var task model.Task

	res, err := r.Db.Query(SQLGetTaskById, id)
	if err != nil {
		log.Debugf("Database.GetTaskById: %+v", err)

		return nil, err
	}
	defer res.Close()

	if res.Next() {
		err = res.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			log.Debugf("Database.GetTaskById: %+v", err)

			return nil, err
		}
	}

	if err = res.Err(); err != nil {
		return nil, err
	}

	if task.Id == "" {
		err = fmt.Errorf("task id %s not found", id)
		log.Debugf("Database.GetTaskById: %+v", err)

		return nil, err
	}

	return &task, nil
}

func (r *TaskRepo) UpdateTask(task *model.Task) error {
	_, err := r.Db.Exec(SQLUpdateTask, task.Id, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		log.Debugf("Database.UpdateTask: %+v", err)

		return err
	}

	return nil
}

func (r *TaskRepo) MakeTaskDone(id string, date string) error {
	res, err := r.Db.Exec(SQLMakeTaskDone, id, date)
	if err != nil {
		log.Debugf("Database.MakeTaskDone: %+v", err)

		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Debugf("Database.MakeTaskDone: %+v", err)

		return err
	}

	if count == 0 {
		err = fmt.Errorf("task id %s not found", id)
		log.Debugf("Database.MakeTaskDone: %+v", err)

		return err
	}

	return nil
}

func (r *TaskRepo) DeleteTask(id string) error {
	res, err := r.Db.Exec(SQLDeleteTask, id)
	if err != nil {
		log.Debugf("Database.DeleteTask: %+v", err)

		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Debugf("Database.DeleteTask: %+v", err)

		return err
	}

	if count == 0 {
		err = fmt.Errorf("task id %s not found", id)
		log.Debugf("Database.DeleteTask: %+v", err)

		return err
	}

	return nil
}
