package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-mongdb-CRUD/services"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func homeLink(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome to goMongo!")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connectURI := "mongodb+srv://uet-class:uetclass@customerinfo.zhan85p.mongodb.net/test"
	clientOptions := options.Client().ApplyURI(connectURI)
	client, _ := mongo.Connect(ctx, clientOptions)
	services.Collection = client.Database("user").Collection("userInfo")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	router.HandleFunc("/api/getAllUser", services.GetAllUser).Methods("GET")
	router.HandleFunc("/api/getUserByID/{id}", services.GetUserByID).Methods("GET")
	router.HandleFunc("/api/createUser", services.CreateUser).Methods("POST")
	router.HandleFunc("/api/updateUser/{id}", services.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/deleteUser/{id}", services.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
