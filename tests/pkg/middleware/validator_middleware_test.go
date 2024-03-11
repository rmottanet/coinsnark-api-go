package middleware_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "coinsnark/api/pkg/middleware"
)

func TestValidateInput_ValidInput(t *testing.T) {
    // Defina um manipulador de teste fictício
    mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte("ValidInput"))
    })

    // Crie um request HTTP falsificado com dados de entrada válidos
    req := httptest.NewRequest("GET", "/convert?from=USD&to=EUR&amount=10.50", nil)

    // Crie um gravador de resposta falsificado para capturar a resposta
    recorder := httptest.NewRecorder()

    // Crie um middleware de validação usando a função ValidateInput
    validationMiddleware := middleware.ValidateInput(mockHandler)

    // Execute a solicitação HTTP falsificada usando o middleware de validação
    validationMiddleware.ServeHTTP(recorder, req)

    // Verifique se a resposta está correta
    if status := recorder.Code; status != http.StatusOK {
        t.Errorf("Handler retornou status incorreto: obteve %v queria %v",
            status, http.StatusOK)
    }

    expectedResponse := "ValidInput"
    if recorder.Body.String() != expectedResponse {
        t.Errorf("Handler retornou resposta incorreta: obteve %v queria %v",
            recorder.Body.String(), expectedResponse)
    }
}

func TestValidateInput_InvalidInput(t *testing.T) {

    req := httptest.NewRequest("GET", "/convert?from=USDD&to=EUR&amount=abc", nil)

    recorder := httptest.NewRecorder()

    mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

    validationMiddleware := middleware.ValidateInput(mockHandler)

    validationMiddleware.ServeHTTP(recorder, req)

    if status := recorder.Code; status != http.StatusBadRequest {
        t.Errorf("Handler retornou status incorreto: obteve %v queria %v",
            status, http.StatusBadRequest)
    }
}
