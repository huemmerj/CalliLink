package main

// simple webserver
import (
	"fmt"
	"net/http"
	"callisplanics/middleware"
	"callisplanics/pages"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

 func main() {
	http.Handle("/", middleware.Layout(pages.Home()))

	http.Handle("/about", middleware.Layout(pages.About()))
	
	// start server and log error
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println(err)
	}
}
