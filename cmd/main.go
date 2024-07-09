package main

import (
    "log"
    "net/http"
    "github.com/bigpandaboy2/proxy-server/internal"
    _ "github.com/bigpandaboy2/proxy-server/docs" // local swag docs
)

// @title Proxy Server API
// @version 1.0
// @description This is a sample server for proxying HTTP requests.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
    server := internal.NewServer()
    server.SetupRoutes()
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}