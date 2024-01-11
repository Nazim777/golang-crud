package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"user-api/controller"
)


func InitilizeRouter(){
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/todos", controller.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", controller.GetTodo).Methods("GET")
	router.HandleFunc("/todos", controller.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", controller.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controller.DeleteTodo).Methods("DELETE")

	// Start the server
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}