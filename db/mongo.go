package db

import (
	"callisplanics/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var MongoClient *mongo.Client

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	coll := client.Database("sample_restaurants").Collection("restaurants")
	parameters := models.ExerciseParameters{Time: true, Reps: true, Weight: true}
	newExercise := models.Exercise{Name: "Pushup", Parameters: parameters}

	result, err := coll.InsertOne(context.TODO(), newExercise)

	if err != nil {

		panic(err)
	}
	log.Print(result)

}

func GetMongoClient() *mongo.Client {
	return MongoClient
}
