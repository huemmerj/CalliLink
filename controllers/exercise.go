package controllers

import (
	"callisplanics/db"
	"callisplanics/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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
