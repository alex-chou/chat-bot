SHELL := /bin/bash

dev:
	ENVIRONMENT=development SERVER_PORT=8000 go run cmd/chatbot/main.go
