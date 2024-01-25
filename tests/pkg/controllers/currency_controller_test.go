package controllers_test

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "coinsnark/api/pkg/controllers"
    "coinsnark/api/pkg/models"
    "coinsnark/tests/pkg/mocks"
)

func TestGetCurrencyNames(t *testing.T) {
    mockCache := &mocks.MockCache{
        Data: map[string]float64{
            "USD": 1.0,
            "EUR": 0.85,
            "GBP": 0.73,
        },
    }

    mockService := mocks.NewMockCurrencyService(mockCache)

    controller := controllers.NewCurrencyController(mockService)

    req, err := http.NewRequest("GET", "/currency", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()

    controller.GetCurrencyNames(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler retornou status %v em vez de %v", status, http.StatusOK)
    }

    var response models.CurrencyResponse
    err = json.Unmarshal(rr.Body.Bytes(), &response)
    if err != nil {
        t.Errorf("erro ao decodificar resposta JSON: %v", err)
    }

    expectedResponse := models.NewCurrencyResponse(map[string]string{
        "USD": "United States Dollar",
        "EUR": "Euro",
        "GBP": "British Pound Sterling",
    }, time.Now())

    if !compareCurrencyResponses(response, *expectedResponse) {
        t.Errorf("resposta incorreta. Resposta recebida: %+v. Resposta esperada: %+v", response, expectedResponse)
    }

}


func compareCurrencyResponses(a, b models.CurrencyResponse) bool {
    if len(a.Currencies) != len(b.Currencies) {
        return false
    }
    for key, value := range a.Currencies {
        if bValue, ok := b.Currencies[key]; !ok || bValue != value {
            return false
        }
    }
    return true
}
