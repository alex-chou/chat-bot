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
	// serverPort to listen to
	serverPort = "8000"
)

func main() {
	if e := os.Getenv("ENVIRONMENT"); e == "" {
		log.Printf("ENVIRONMENT not set. Using: %s", environment)
	} else {
		environment = e
	}

	if sp := os.Getenv("SERVER_PORT"); sp == "" {
		log.Printf("SERVER_PORT not set. Using: %s", serverPort)
	} else {
		serverPort = sp
	}

	log.Printf("Running chatbot in %s on port %s", environment, serverPort)

	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", serverPort), mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
