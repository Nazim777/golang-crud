package controller

import (
	"context"
	"net/http"
	"time"

	"user-api/database"
	"user-api/helpers"
	"user-api/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	cursor, err :=database.Collection.Find(context.Background(), bson.M{})
	if err != nil {
	helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			helpers.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		todos = append(todos, todo)
	}

	helpers.RespondJSON(w, http.StatusOK, todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	objectId, error := primitive.ObjectIDFromHex(params["id"])
	if error != nil{
		helpers.RespondError(w, http.StatusNotFound, "Invalid Id!")
		return
	}
	var todo models.Todo
	err := database.Collection.FindOne(context.Background(), bson.M{"_id":objectId}).Decode(&todo)
	if err != nil {
		helpers.RespondError(w, http.StatusNotFound, "Todo not found")
		return
	}
	helpers.RespondJSON(w, http.StatusOK, todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := helpers.ParseJSON(r, &todo); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	todo.CreatedAt = time.Now()
	result, err := database.Collection.InsertOne(context.Background(), todo)
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Query the document with the generated ID
	err = database.Collection.FindOne(context.Background(), bson.M{"_id": result.InsertedID}).Decode(&todo)
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, todo)
}



func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	objectId, error := primitive.ObjectIDFromHex(params["id"])
	if error != nil{
        helpers.RespondError(w, http.StatusNotFound, "Invalid Id!")
		return
	}
	var todo models.Todo
	if err := helpers.ParseJSON(r, &todo); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err := database.Collection.UpdateOne(context.Background(), bson.M{"_id": objectId}, bson.M{"$set": todo})
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondJSON(w, http.StatusOK, todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	objectId,error := primitive.ObjectIDFromHex(params["id"])
	if error != nil{
        helpers.RespondError(w, http.StatusNotFound, "Invalid Id!")
		return
	}
	result, err := database.Collection.DeleteOne(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if result.DeletedCount == 0 {
		helpers.RespondError(w, http.StatusNotFound, "Todo not found")
		return
	}
	helpers.RespondJSON(w, http.StatusNoContent, nil)
}