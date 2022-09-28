package services

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-mongdb-CRUD/models"
	"net/http"
	"time"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user models.User
	json.NewDecoder(request.Body).Decode(&user)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := Collection.InsertOne(ctx, user)
	id := result.InsertedID
	user.ID = id.(primitive.ObjectID)
	json.NewEncoder(response).Encode(user)
}
