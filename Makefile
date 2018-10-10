SHELL := /bin/bash

dev:
	# Run chatbot locally on port 8000.
	ENVIRONMENT=development PORT=8000 go run cmd/chatbot/main.go

build:
	# Build chatbot locally.
	go build -o build/chatbot ./cmd/chatbot

test:
	# Run tests locally.
	go test -race ./...

test_coverage:
	# Run tests and generate coverage profile.
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

mocks:
	# Generate mocks for testing.
	rm -rf ./internal/backend/backendmocks
	rm -rf ./pkg/slack/slackmocks
	mockery -dir ./internal/backend -name Backend -outpkg backendmocks -output ./internal/backend/backendmocks
	mockery -dir ./pkg/slack -name Slack -outpkg slackmocks -output ./pkg/slack/slackmocks

heroku_push: build test
	# Build and run tests locally before pushing to heroku.
ifneq ($(shell git rev-parse --abbrev-ref HEAD), master)
	$(error Not on branch master)
else
	git push heroku master
endif

heroku_logs:
	# Tail logs on chatbot heroku app.
	heroku logs --tail
