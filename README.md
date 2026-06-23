# Agent Gateway

Phase 1 implementation of a Go-based AI Agent Gateway.

## Features

- REST API
- Mock LLM Provider
- Provider abstraction
- Unit tests

## Installing sqlite driver
go get github.com/mattn/go-sqlite3

## Run

go run cmd/server/main.go

## Running Unit Test

go test ./...

## Test

### Request
curl -X POST \
http://localhost:8080/chat \
-H "Content-Type: application/json" \
-d '{"message":"hello"}'

### Response
{
  "response": "Hello! How can I help you?"
}