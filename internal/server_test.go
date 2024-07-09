package internal

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func MockServer() *httptest.Server {
    handler := http.NewServeMux()
    handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"message": "success"}`))
    })
    return httptest.NewServer(handler)
}

func TestHandleRequest(t *testing.T) {
    server := NewServer()
    mockServer := MockServer()
    defer mockServer.Close()

    reqBody, _ := json.Marshal(Request{
        Method:  "GET",
        URL:     mockServer.URL,
        Headers: map[string]string{"Content-Type": "application/json"},
    })

    req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    server.HandleRequest(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("expected status OK; got %v", res.Status)
    }

    var response Response
    if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
        t.Errorf("could not decode response: %v", err)
    }

    if response.Status != http.StatusOK {
        t.Errorf("expected status OK in response; got %v", response.Status)
    }

    if response.ID == "" {
        t.Errorf("expected non-empty ID in response")
    }

    invalidReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("invalid json")))
    invalidReq.Header.Set("Content-Type", "application/json")

    invalidW := httptest.NewRecorder()
    server.HandleRequest(invalidW, invalidReq)

    invalidRes := invalidW.Result()
    defer invalidRes.Body.Close()

    if invalidRes.StatusCode != http.StatusBadRequest {
        t.Errorf("expected status BadRequest; got %v", invalidRes.Status)
    }
}

func TestHealthCheck(t *testing.T) {
    server := NewServer()

    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    w := httptest.NewRecorder()
    server.HealthCheck(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("expected status OK; got %v", res.Status)
    }
}

func TestProxyDifferentMethods(t *testing.T) {
    server := NewServer()
    mockServer := MockServer()
    defer mockServer.Close()

    methods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}

    for _, method := range methods {
        reqBody, _ := json.Marshal(Request{
            Method:  method,
            URL:     mockServer.URL,
            Headers: map[string]string{"Content-Type": "application/json"},
        })

        req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
        req.Header.Set("Content-Type", "application/json")

        w := httptest.NewRecorder()
        server.HandleRequest(w, req)

        res := w.Result()
        defer res.Body.Close()

        if res.StatusCode != http.StatusOK {
            t.Errorf("expected status OK; got %v for method %s", res.Status, method)
        }

        var response Response
        if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
            t.Errorf("could not decode response for method %s: %v", method, err)
        }

        if response.ID == "" {
            t.Errorf("expected non-empty ID in response for method %s", method)
        }
    }
}

func TestExternalServiceDown(t *testing.T) {
    server := NewServer()

    reqBody, _ := json.Marshal(Request{
        Method:  "GET",
        URL:     "http://nonexistent.url",
        Headers: map[string]string{"Content-Type": "application/json"},
    })

    req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    server.HandleRequest(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusInternalServerError {
        t.Errorf("expected status InternalServerError; got %v", res.Status)
    }
}