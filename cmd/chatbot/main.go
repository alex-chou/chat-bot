package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

var (
	// environment the server is running in
	environment = "development"
	// port to listen to
	port = "8000"
)

func main() {
	if e := os.Getenv("ENVIRONMENT"); e == "" {
		log.Printf("ENVIRONMENT not set. Using: %s", environment)
	} else {
		environment = e
	}

	if p := os.Getenv("PORT"); p == "" {
		log.Printf("PORT not set. Using: %s", port)
	} else {
		port = p
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler)

	log.Printf("Running chatbot in %s on port %s", environment, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
