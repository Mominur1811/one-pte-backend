# OnePTE Project

## Overview
This project is a Go-based application with support for Linux/macOS and Windows development environments.

## Prerequisites
- Docker
- Docker Compose
- Go (latest version recommended)
- Make

## Project Dependencies
- Go Modules
- Air (for live reloading during development)
- golang-migrate (for database migrations)

## Initial Setup

### 1. Clone the Repository
```bash
git clone https://github.com/Mominur1811/one-pte-backend.git
```

### 2. Start Docker Services
Before running the project, start the required Docker services:
```bash
docker-compose up -d
```

### 3. Install Development Dependencies
Install the necessary development tools:
```bash
make prepare
```

This command will:
- Install Air (live reload tool)
- Install golang-migrate
- Download Go modules
- Tidy module dependencies

## Development Workflow

### For Linux/macOS
- Start development server:
  ```bash
  make dev
  ```

### For Windows
- Start development server:
  ```bash
  make dev_win
  ```

## Building the Project

### Linux/macOS
```bash
make build-linux
```

### Windows
```bash
make build-win
```

## Running the Server

### Linux/macOS
```bash
make run-server-linux
```

### Windows
```bash
make run-server-win
```

## Convenience Commands
- `make tidy`: Clean and organize Go modules
- `make fmt`: Format Go files using gofumpt and golines

## API Documentation
API documentation is available in the Postman collection:
- File: `onepte-api-momin.postman_collection.json`
- Import this file into Postman to explore and test the API endpoints

## Troubleshooting
- Ensure Docker is running before starting the project
- Check that all required dependencies are installed
- Verify your Go version is compatible

## License
[Add your project's license information here]
