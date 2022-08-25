package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"practice_1/initializers"
	"practice_1/models"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var users []models.User

	result, err := initializers.Client.Database("practice_1").Collection("users").Find(context.TODO(), bson.D{{}})
	if err != nil {
		w.WriteHeader(404)
		return
	}

	for result.Next(context.TODO()) {
		var elem models.User
		err := result.Decode(&elem)
		if err != nil {
			panic(err)
		}
		users = append(users, elem)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	_, err := initializers.Client.Database("practice_1").Collection("users").InsertOne(context.TODO(), u)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
