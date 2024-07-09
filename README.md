# Proxy Server

## Overview

This is an HTTP server that proxies requests to external services, processes these requests, and returns responses to the client in JSON format. The server also saves requests and responses locally.

## Setup

1. Clone the repository:
    ```sh
    git clone <your-repo-url>
    cd proxy-server
    ```

2. Build and run the server using Docker Compose:
    ```sh
    make build
    make up
    ```

3. Access the server at `http://localhost:8080`.

## API Endpoints

- `POST /` - Proxy requests to external services.
- `GET /health` - Healthcheck endpoint.
- `GET /swagger/` - Swagger API documentation.

## Request Format

```json
{
  "method": "GET",
  "url": "http://google.com",
  "headers": { "Authentication": "Basic bG9naW46cGFzc3dvcmQ=" }
}

## Response Format

```json
{
  "id": "unique-request-id",
  "status": 200,
  "headers": { "Content-Type": "application/json" },
  "length": 123
}

### Conclusion

The server code was implemented to handle HTTP requests, validate inputs, and return JSON responses with unique request IDs. Local storage was achieved using sync.Map, and a healthcheck endpoint was added. Additionally, we used the swaggo/swag package for Swagger documentation and containerized the application for deployment on Render.# proxy-server
