package services

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-mongdb-CRUD/models"
	"net/http"
	"time"
)

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user models.User
	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])
	json.NewDecoder(request.Body).Decode(&user)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := Collection.UpdateOne(ctx, models.User{ID: id}, bson.M{"$set": user})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}
