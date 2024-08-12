package main

// simple webserver
import (
	"callisplanics/db"
	"callisplanics/middleware"
	"callisplanics/pages"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitMongoDB()
	http.Handle("/", middleware.Layout(pages.Exercises()))

	http.Handle("/about", middleware.Layout(pages.About()))

	http.Handle("/addExercise", middleware.Layout(pages.AddExerciseHandler()))
	// Static serve the dist folder
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
	// start server and log error
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println(err)
	}
}
