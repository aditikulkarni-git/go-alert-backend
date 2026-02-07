# Go Alert Backend

A simple Go backend service built to demonstrate clean architecture, REST APIs,
basic observability concepts, and CI using GitHub Actions.




---

## Features

- Health check endpoint
- Alerts API (list alerts)
- In-memory task store
- Structured project layout
- GitHub Actions CI pipeline
- Ready for future observability & database extensions

---

## Tech Stack

- Go (net/http)
- GitHub Actions (CI)
- In-memory data store (map)

---

## Project Structure

go-alert-backend/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── model/
│   └── store/
├── .github/
│   └── workflows/
│       └── ci.yml
├── go.mod
├── go.sum
└── README.md

---

## Running Locally

go mod tidy
go run cmd/server/main.go

Server runs on:

http://localhost:8080

Health check:

curl http://localhost:8080/health

---

## CI Pipeline

This project uses GitHub Actions to automatically build and validate the code
on every push to the main branch.

A successful pipeline run is indicated by a green checkmark in GitHub.

---

## Future Improvements

- PostgreSQL integration
- Metrics endpoint
- Docker support
- Tracing & logging

---

## Author

Aditi Kulkarni

