SHELL := /bin/bash

dev:
	ENVIRONMENT=development PORT=8000 go run cmd/chatbot/main.go
