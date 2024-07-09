package internal

import (
    "net/http"
    "github.com/swaggo/http-swagger"
    _ "github.com/bigpandaboy2/proxy-server/docs"
)

func (s *Server) SetupRoutes() {
    http.HandleFunc("/health", s.HealthCheck)
    http.HandleFunc("/", s.HandleRequest)
    http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
}