package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"user-api/database"
	"user-api/models"
	"user-api/helpers"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	database.Db.Find(&todos)
	helpers.RespondJSON(w, http.StatusOK, todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	database.Db.First(&todo, params["id"])
	helpers.RespondJSON(w, http.StatusOK, todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := helpers.ParseJSON(r, &todo); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	database.Db.Create(&todo)
	helpers.RespondJSON(w, http.StatusCreated, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	if err := helpers.ParseJSON(r, &todo); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	database.Db.Model(&todo).Where("id = ?", params["id"]).Updates(todo)
	helpers.RespondJSON(w, http.StatusOK, todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	database.Db.Delete(&todo, params["id"])
	helpers.RespondJSON(w, http.StatusNoContent, nil)
}
