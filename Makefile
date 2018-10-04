SHELL := /bin/bash

dev:
	# Run chatbot locally on port 8000.
	ENVIRONMENT=development PORT=8000 go run cmd/chatbot/main.go

build:
	# Build chatbot locally.
	go build -o build/chatbot ./cmd/chatbot

test:
	# Run tests locally
	go test -race ./...

heroku_push: build test
	# Build and run tests locally before pushing to heroku.
	git push heroku master

heroku_logs:
	# Tail logs on chatbot heroku app.
	heroku logs --tail
