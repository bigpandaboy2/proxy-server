package main

import (
    "log"
    "net/http"
    "github.com/bigpandaboy2/proxy-server/internal"
    _ "github.com/bigpandaboy2/proxy-server/docs" 
)

func main() {
    server := internal.NewServer()
    server.SetupRoutes()
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}