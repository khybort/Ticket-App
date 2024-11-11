# API for Tickets Project

## Overview
This is the backend API for our application, built with Go. It provides the necessary endpoints and business logic to support the frontend UI.

## Prerequisites
- Go 1.23
- PostgreSQL

## Project Summary:
The project is an API built using the Go programming language.
It follows the Clean Architecture principles, separating concerns into layers (domain, use cases, repositories).
The project uses the Gorilla Mux package for routing and the Testify library for testing.
The project includes a Makefile for building, testing, and running the API.

## Features
- Go-based backend API
- RESTful endpoint design
- OpenAPI support
- Environment-specific configurations
- Production-ready setup
- Database integration (assuming based on typical Go backend setups)
- Docker support for easy deployment and scaling

## Getting Started
1. Ensure you have Go installed on your system.
2. Clone the repository and navigate to the api directory.
3. Install dependencies:
    go mod download
4. Run the server:
    go run main.go

## API Documentation
The API documentation is generated using the OpenAPI specification (Swagger). You can find the API documentation in the docs/swagger.yaml file.

To view the API documentation, you can use "http://localhost:8000/swagger/index.html" link

## API Endpoints
The API provides the following endpoints:

| Method | Endpoint                      | Description                          |
|--------|-------------------------------|--------------------------------------|
| GET    | /api/v1/tickets               | Retrieve a list of all tickets       |
| GET    | /api/v1/tickets/:id           | Retrieve a specific ticket by ID     |
| POST   | /api/v1/tickets               | Create a new ticket                  |
| POST   | /api/v1/tickets/:id/purchases | Purchase a ticket by ID              |

These endpoints are implemented using the Gin framework in the main.go file. The corresponding functions in the controllers/ticket_handler.go file handle the requests and interact with the use cases to perform the necessary operations.

## Configuration
The API is configured to run on port 8000 by default, as indicated by the frontend's `.env.production` file.

## Docker Support
This API is designed to be run in a Docker container. See the root `docker-compose.yml` file for more details on how it's configured in the overall application stack.

## Database
(Provide information about the database used, connection details, migrations, etc.)

## Testing
To run the tests:
    go test -v -cover ./...

## Contributing
Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for more information.

## Additional Information
- The API is designed to work in conjunction with the React frontend.
- For local development, ensure that the API is running on `http://localhost:8000` to match the frontend's expectations.

For more detailed information about the Go setup and available commands, refer to the Go documentation.