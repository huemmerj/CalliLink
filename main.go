package main

// simple webserver
import (
	"callisplanics/controllers"
	"callisplanics/db"
	"callisplanics/middleware"
	"callisplanics/pages"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	r := mux.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitMongoDB()
	r.Handle("/", middleware.Layout(pages.Exercises()))

	r.Handle("/about", middleware.Layout(pages.About()))

	r.Handle("/addExercise", middleware.Layout(pages.AddExerciseHandler())).Methods("GET")
	r.HandleFunc("/addExercise", controllers.AddExerciseHandler).Methods("POST")
	// Static serve the dist folder
	r.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
	// start server and log error
	err = http.ListenAndServe(":8082", r)
	if err != nil {
		fmt.Println(err)
	}
}
