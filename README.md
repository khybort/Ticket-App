# Tickets Projects

## Overview
This project is a full-stack application consisting of a React frontend (UI), a Go backend API, and is orchestrated using Docker Compose for easy development and deployment.

## Components
1. **UI**: NGINX and React-based frontend [Go to UI README](./ui/README.md)
2. **API**: Go-based backend API [Go to API README](./api/README.md)
3. **Docker Compose**: Orchestration for multi-container Docker applications

## Docker Compose Configuration
The `docker-compose.yml` file in the root directory defines the services, networks, and volumes for our application. It typically includes:

- **UI Service**: 
  - Builds the React frontend
  - Exposes it on port 3000 but Nginx uses it on port 80

- **API Service**:
  - Builds the Go backend API
  - Exposes it on port 8000
  - May include environment variables for database connections, etc.

- **Database Service** (if applicable):
  - Likely uses a standard database image (e.g., PostgreSQL)
  - Defines volume for data persistence

## Getting Started
1. Ensure Docker and Docker Compose are installed on your system.
2. Clone the repository.
3. From the root directory, run:
    make build-prod 
    make up-prod
4. Access the application:
- Frontend: `http://localhost`
- Backend API: `http://localhost:8000`
- OpenAPI documentation: `http://localhost:8000/swagger/index.html`
5. You can also look at the Makefile for more operations

## Project Structure
    .
    ├── api
    │   ├── cmd
    │   ├── docs
    │   ├── internal
    │   │   ├── config
    │   │   ├── controllers
    │   │   ├── database
    │   │   ├── domain
    │   │   ├── repositories
    │   │   └── usecases
    │   ├── go.mod
    │   └── main.go
    ├── docker-compose.yml
    ├── Makefile
    ├── README.md
    └── ui
        ├── public
        ├── src
        │   ├── components
        │   ├── pages
        │   ├── services
        │   └── App.js
        ├── package.json
        └── README.md

### Explanation of the Project Structure

- **`api/`**: This folder contains the backend application.
  - **`cmd/`**: Typically stores the entry point of the application (e.g., main.go).
  - **`docs/`**: Documentation files related to the backend.
  - **`internal/`**: Contains all internal logic, typically structured to include:
    - **`config/`**: Configuration files for the backend.
    - **`controllers/`**: Handlers for HTTP requests, routing.
    - **`database/`**: Database connection and management.
    - **`domain/`**: Core domain models.
    - **`repositories/`**: Data persistence and database interactions.
    - **`usecases/`**: Business logic.
  - **`go.mod`**: Go module definition for dependency management.
  - **`main.go`**: The entry point of the Go application.
  
- **`docker-compose.yml`**: Defines and runs multi-container Docker applications.
- **`Makefile`**: Contains automation tasks (build, test, deploy).
- **`README.md`**: Project documentation.

- **`ui/`**: The frontend part of the project.
  - **`public/`**: Static files (e.g., images, favicon).
  - **`src/`**: Source code for the frontend application.
    - **`components/`**: Reusable React components.
    - **`pages/`**: Page components.
    - **`services/`**: External API calls or utilities.
    - **`App.js`**: Main React component.
  - **`package.json`**: Node.js dependencies and project metadata.
  - **`README.md`**: Documentation for the frontend application.

## Test
  make test

## Monitor logging
  `make logb-prod` for backend logging
  `make logf-prod` for frontend logging

## Update Documentation
  swag init
  swag fmt

## Development Workflow
- For frontend development, you can work directly in the `ui` directory and use standard React development commands.
- For backend development, work in the `api` directory and use Go commands for testing and running locally.
- Use Docker Compose for integrated testing of the full stack.

## Production Deployment
The Docker Compose setup can be used as a basis for production deployment, but you may need to adjust settings for security and scalability in a production environment.

## Additional Notes
- The `.env.production` file in the UI directory suggests that the application is configured for a production environment.
- Ensure all necessary environment variables are properly set in the Docker Compose file or through a `.env` file for secure configuration management.

For more detailed information about each component, refer to the README files in the respective `ui` and `api` directories.