package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Define your model
type Todo struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"desc"`
}

var db *gorm.DB

func main() {
	// Connect to MySQL database
	dsn := "root:admin@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate will create the table. Make sure your MySQL server is running.
	db.AutoMigrate(&Todo{})

	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	// Start the server
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler functions

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	db.Find(&todos)
	respondJSON(w, http.StatusOK, todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo Todo
	db.First(&todo, params["id"])
	respondJSON(w, http.StatusOK, todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := parseJSON(r, &todo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	db.Create(&todo)
	respondJSON(w, http.StatusCreated, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo Todo
	if err := parseJSON(r, &todo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	db.Model(&todo).Where("id = ?", params["id"]).Updates(todo)
	respondJSON(w, http.StatusOK, todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo Todo
	db.Delete(&todo, params["id"])
	respondJSON(w, http.StatusNoContent, nil)
}

// Helper functions

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

func parseJSON(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}
