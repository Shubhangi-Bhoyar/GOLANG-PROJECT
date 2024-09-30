package main

import (
	// "context"
	"log"
	"mongodb-with-golang/controllers"
	"net/http"

	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	route := mux.NewRouter()

	// route.HandleFunc("/", controllers.Page)

	// subroute := route.PathPrefix("/todolist").Subrouter()

	route.HandleFunc("/api/v1/task", controllers.Task).Methods(http.MethodPost)
	route.HandleFunc("/api/v1/task3",controllers.Task3).Methods(http.MethodPatch)
	route.HandleFunc("/api/v1/task1",controllers.Task1).Methods(http.MethodGet)
	route.HandleFunc("/api/v1/task2",controllers.Task2).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", route))
}
