package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"coinsnark/api/pkg/controllers"
	"coinsnark/api/pkg/models"

)


type MockConversionService struct{}

func (m *MockConversionService) Convert(from, to string, amount float64) (float64, time.Time, error) {
    return 100.0, time.Date(2024, 2, 5, 12, 0, 0, 0, time.UTC), nil
}

func TestConvertCurrency(t *testing.T) {
	mockService := &MockConversionService{}

	controller := controllers.NewConversionController(mockService)

	reqBody := []byte(`{"from": "USD", "to": "EUR", "amount": 100}`)
	req, err := http.NewRequest("GET", "/convert?from=USD&to=EUR&amount=100", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	controller.ConvertCurrency(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou status %v em vez de %v", status, http.StatusOK)
	}

	var response models.ConversionResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("erro ao decodificar resposta JSON: %v", err)
	}

	expectedResponse := models.NewConversionResponse("USD", "EUR", 100.0, time.Date(2024, 2, 5, 12, 0, 0, 0, time.UTC))
	if response.From != expectedResponse.From || response.To != expectedResponse.To || response.Converted != expectedResponse.Converted || response.CacheUpdated != expectedResponse.CacheUpdated {
		t.Errorf("resposta incorreta")
	}
}
