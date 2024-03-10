package middleware_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "coinsnark/api/pkg/middleware"
)

func TestConvertHandler(t *testing.T) {

    mockHandler := func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte("ConvertHandlerFunc foi chamada com sucesso"))
    }

    req := httptest.NewRequest("GET", "/convert", nil)

    recorder := httptest.NewRecorder()

    handler := middleware.ConvertHandler(mockHandler)

    handler.ServeHTTP(recorder, req)

    if status := recorder.Code; status != http.StatusOK {
        t.Errorf("Handler retornou status incorreto: obteve %v queria %v",
            status, http.StatusOK)
    }

    expectedResponse := "ConvertHandlerFunc foi chamada com sucesso"
    if recorder.Body.String() != expectedResponse {
        t.Errorf("Handler retornou resposta incorreta: obteve %v queria %v",
            recorder.Body.String(), expectedResponse)
    }
}
