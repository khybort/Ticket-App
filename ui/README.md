# UI for Tickets Project

## Overview
This project is a React-based user interface for our application. It's designed to work in conjunction with our backend API and is containerized for easy deployment.

## Features
- React-based frontend
- NGINX
- Environment-specific configurations
- Production-ready setup

## Environment Configuration
The project uses environment-specific configuration files:

- `.env.production`: Contains production environment settings
  - `REACT_APP_BACKEND_BASE_URL`: Set to communicate with the backend API
  - `NODE_ENV`: Set to 'production' for optimized build
  - `FAST_REFRESH`: Disabled in production for stability

## Getting Started
1. Install dependencies:
    npm install
2. Start the development server:
    npm start
3. Build for production:
    npm run build


## Docker Support
This project is designed to be run in a Docker container. See the root `docker-compose.yml` file for more details on how it's configured in the overall application stack.

## Additional Information
- The project is set up to communicate with a backend running at `http://localhost:8000` in the production environment.
- Fast Refresh is disabled in production for improved stability and performance.

For more detailed information about the React setup and available scripts, refer to the Create React App documentation.
