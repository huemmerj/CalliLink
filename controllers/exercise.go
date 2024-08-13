package controllers

import (
	"callisplanics/db"
	"callisplanics/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func GetExercisesHandler() []models.Exercise {
	client := db.GetMongoClient()
	coll := client.Database("sample_restaurants").Collection("restaurants")

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var results []models.Exercise
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}
func AddExerciseHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Unable to parse form")
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	description := r.FormValue("description")

	log.Print(name)
	log.Print(description)
	w.WriteHeader(http.StatusCreated)
	return GetExercisesHandler()
}
func AddExerciseHandle(exercise models.Exercise) {
	client := db.GetMongoClient()
	coll := client.Database("sample_restaurants").Collection("restaurants")

	result, err := coll.InsertOne(context.TODO(), exercise)

	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
}
