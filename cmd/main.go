package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"

	"github.com/snazimen/go_final_project/api"
	"github.com/snazimen/go_final_project/config"
	"github.com/snazimen/go_final_project/middle"
	"github.com/snazimen/go_final_project/repository"
	"github.com/snazimen/go_final_project/usecases"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(cfg.LogLevel)

	db, err := repository.NewDB(cfg.DBFile)
	if err != nil {
		log.Fatalf("Error connect to repository: %+v", err)
	}
	defer db.Close()

	// init repository
	repo := repository.NewNewRepository(db)

	// init usecases
	taskUC := usecases.NewTaskUsecase(repo)
	taskHandler := api.NewTaskHandler(taskUC)

	// init middleware
	authMiddleware := middle.New(cfg)

	// init auth
	authHandler := middle.NewAuthHandler(cfg)

	webDir := "./web"
	r := chi.NewRouter()
	fileServer := http.FileServer(http.Dir(webDir))
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if filepath.Ext(r.URL.Path) == ".css" {
			w.Header().Set("Content-Type", "text/css")
		}
		fileServer.ServeHTTP(w, r)
	})
	r.Post("/api/signin", authHandler.GetAuthByPassword)
	r.Get("/api/nextdate", taskHandler.GetNextDate)
	r.Post("/api/task", authMiddleware.Auth(taskHandler.CreateTask))
	r.Get("/api/tasks", authMiddleware.Auth(taskHandler.GetTasks))
	r.Get("/api/task", authMiddleware.Auth(taskHandler.GetTask))
	r.Put("/api/task", authMiddleware.Auth(taskHandler.UpdateTask))
	r.Post("/api/task/done", authMiddleware.Auth(taskHandler.MakeTaskDone))
	r.Delete("/api/task", authMiddleware.Auth(taskHandler.DeleteTask))

	serverAddress := fmt.Sprintf("localhost:%s", cfg.Port)
	log.Infoln("Listening on " + serverAddress)
	if err = http.ListenAndServe(serverAddress, r); err != nil {
		log.Panicf("Start server error: %+v", err.Error())
	}
}
