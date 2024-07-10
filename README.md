# Proxy Server

## Overview

This is an HTTP server that proxies requests to external services, processes these requests, and returns responses to the client in JSON format. The server also saves requests and responses locally.

## Setup

1. Clone the repository:
    ```sh
    git remote add origin https://github.com/bigpandaboy2/proxy-server.git
    cd proxy-server
    ```

2. Build and run the server using Docker Compose:
    ```sh
    make build
    make up
    ```

3. Access the server at `http://localhost:8080`.

## **Deployed on RENDER:**

**Base URL**

[https://proxy-server-ghjq.onrender.com](url)

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
  "method": "GET",
  "url": "http://google.com",
  "headers": { "Authentication": "Basic bG9naW46cGFzc3dvcmQ=" }
}
