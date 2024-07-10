# Proxy Server

## Overview

This is an HTTP server that proxies requests to external services, processes these requests, and returns responses to the client in JSON format. The server also saves requests and responses locally.

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/bigpandaboy2/proxy-server.git
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

[https://proxy-server-ghjq.onrender.com](https://proxy-server-ghjq.onrender.com)

## API Endpoints

- `POST /` - Proxy requests to external services.
- `GET /health` - Healthcheck endpoint.
- `GET /swagger/` - Swagger API documentation.

## Request Format

```sh
{
  "method": "GET",
  "url": "http://google.com",
  "headers": { "Authorization": "Basic bG9naW46cGFzc3dvcmQ=" }
}
```
## Response Format

```sh
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "status": 200,
  "headers": {
    "Content-Type": ["application/json"]
  },
  "length": 1234
}
