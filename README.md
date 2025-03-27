# Web Page Analyzer

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [How to Build and Run](#how-to-build-and-run)
  - [Building and Running Backend](#building-and-running-backend)
  - [Building and Running Frontend](#building-and-running-frontend)
  - [Running with Docker Compose](#running-with-docker-compose)
- [Running Tests](#running-tests)
  - [Unit Tests](#unit-tests)
  - [Test Coverage](#test-coverage)
- [API Endpoints](#api-endpoints)
- [Possible Improvements](#possible-improvements)

---

## Introduction

Web Page Analyzer is a web application that analyzes web pages based on user-provided URLs. It extracts information such as the HTML version, title, headings count, link details, and login forms.3.3.

The backend is built using **Golang (Gin Framework)**, and the frontend is developed using **React**.

---

## Prerequisites

Before running the project, ensure you have the following installed:

- **Go 1.24.1** or later
- **Node.js 18+** and **npm**
- **Docker & Docker Compose**
- **Git**

---

## Project Structure

```
web-analyzer/
│── backend/web-analyzer
│   │── handlers/
│   │── main.go
│   │── go.mod
│   │── Dockerfile
│   └── ...
│
│── frontend/web_page_analyzer
│   │── src/
│   │── package.json
│   │── Dockerfile
│   └── ...
│
│── docker-compose.yml
│── README.md
└── .gitignore
```

---

## How to Build and Run

### Building and Running Backend

1. Navigate to the backend directory:
   ```sh
   cd backend/web-analyzer
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the backend server:
   ```sh
   go run main.go
   ```
4. The backend should now be running on `http://localhost:8080`.

### Building and Running Frontend

1. Navigate to the frontend directory:
   ```sh
   cd frontend/web_page_analyzer
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
3. Run the frontend server:
   ```sh
   npm start
   ```
4. The frontend should now be accessible at `http://localhost:3000`.

---

### Running with Docker Compose

1. Ensure Docker and Docker Compose are installed.
2. Run the following command to start the backend and frontend together:
   ```sh
   docker-compose up --build
   ```
3. Access the frontend at `http://localhost:3000` and the backend at `http://localhost:8080`.
4. To stop the services:
   ```sh
   docker-compose down
   ```

---

## Running Tests

### Unit Tests

Navigate to the backend directory and run:

```sh
cd backend/web-analyzer
go test ./...
```

For the frontend, run:

```sh
cd frontend/web_page_analyzer
npm test
```

### Test Coverage

For the backend, check test coverage using:

```sh
go test -cover ./...
```

For a detailed coverage report:

```sh
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

For the frontend, use:

```sh
npm test -- --coverage
```

---

## API Endpoints

| Method | Endpoint   | Description                 |
| ------ | ---------- | --------------------------- |
| GET    | `/status`  | Check service health status |
| GET    | `/metrics` | Get Prometheus metrics      |
| GET    | `/urls`    | Get analyzed URLs history   |
| POST   | `/analyze` | Analyze a given web page    |

---

## Possible Improvements

- Store analysis results in a database for historical tracking.
- Implement authentication for access control.
- Deploy the project to a cloud service (AWS, GCP, Azure) using CI/CD pipelines.

---

## How the UI Works

1. **Submit URL**:  
    - The UI contains a text field where users can input a URL and submit it for processing.  
    - A dropdown menu displays the history of previously submitted URLs for quick access.

2. **Handling Long Processing Times**:  
    - If a URL takes longer than 30 seconds to process, a message box will appear.  
    - This message box allows users to stop the loading process and continue working on other tasks in the UI.

---
## License

This project is licensed under the Apache 2.0 License.
