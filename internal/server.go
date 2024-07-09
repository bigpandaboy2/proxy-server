package internal

import (
    "encoding/json"
    "io"
    "net/http"
    "sync"
    "github.com/google/uuid"
)

type Server struct {
    data sync.Map
}

type Request struct {
    Method  string            `json:"method"`
    URL     string            `json:"url"`
    Headers map[string]string `json:"headers"`
}

type Response struct {
    ID      string              `json:"id"`
    Status  int                 `json:"status"`
    Headers map[string][]string `json:"headers"`
    Length  int                 `json:"length"`
}

func NewServer() *Server {
    return &Server{}
}

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    client := &http.Client{}
    httpReq, err := http.NewRequest(req.Method, req.URL, nil)
    if err != nil {
        http.Error(w, "could not create request", http.StatusInternalServerError)
        return
    }
    for k, v := range req.Headers {
        httpReq.Header.Set(k, v)
    }

    httpResp, err := client.Do(httpReq)
    if err != nil {
        http.Error(w, "request failed", http.StatusInternalServerError)
        return
    }
    defer httpResp.Body.Close()

    body, err := io.ReadAll(httpResp.Body)
    if err != nil {
        http.Error(w, "could not read response", http.StatusInternalServerError)
        return
    }

    id := uuid.New().String()
    resp := Response{
        ID:      id,
        Status:  httpResp.StatusCode,
        Headers: httpResp.Header,
        Length:  len(body),
    }
    s.data.Store(id, resp)

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        http.Error(w, "could not encode response", http.StatusInternalServerError)
    }
}