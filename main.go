package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/snazimen/go_final_project/tree/total/database"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	err = database.ConnectDB()
	if err != nil {
		log.Fatalf("Error connect to database: %s", err)
	}

	port := os.Getenv("TODO_PORT")
	if len(port) == 0 {
		port = "7540"
	}
	webDir := "./web"
	r := chi.NewRouter()
	r.Handle("/", http.FileServer(http.Dir(webDir)))
	serverAddress := fmt.Sprintf("localhost:%s", port)
	log.Println("Listening on " + serverAddress)
	if err = http.ListenAndServe(serverAddress, http.FileServer(http.Dir(webDir))); err != nil {
		log.Panicf("Start server error: %s", err.Error())
	}
}
