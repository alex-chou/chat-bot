package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alex-chou/chat-bot/internal/server"
	"github.com/alex-chou/chat-bot/pkg/slack"
)

var (
	// environment the server is running in
	environment = "development"
	// port to listen to
	port = "8000"
)

func main() {
	if err := readConfig(); err != nil {
		log.Fatal(err)
	}

	server := configureServer()

	log.Printf("Running chatbot in %s on port %s", environment, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}

func readConfig() error {
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
	return nil
}

func configureSlack() slack.Slack {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("SLACK_TOKEN not set")
	}
	return slack.New(token)
}

func configureServer() *server.Server {
	return server.NewServer(configureSlack())
}
