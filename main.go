package main

// simple webserver
import (
	"callisplanics/middleware"
	"callisplanics/pages"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.Handle("/", middleware.Layout(pages.Home()))

	http.Handle("/about", middleware.Layout(pages.About()))
	// Static serve the dist folder
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
	// start server and log error
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println(err)
	}
}
