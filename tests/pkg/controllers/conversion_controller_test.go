package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"coinsnark/api/pkg/controllers"
	"coinsnark/api/pkg/models"
)

type MockConversionService struct{}

func (m *MockConversionService) Convert(from, to string, amount float64) (map[string]string, error) {
	response := map[string]string{
		"from":         from,
		"to":           to,
		"converted":    "100.00",
		"cache_updated": time.Date(2024, 2, 5, 12, 0, 0, 0, time.UTC).Format(time.RFC3339),
	}
	return response, nil
}

func TestConvertCurrency(t *testing.T) {
	mockService := &MockConversionService{}

	controller := controllers.NewConversionController(mockService)

	req, err := http.NewRequest("GET", "/convert?from=USD&to=EUR&amount=100", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	controller.ConvertCurrency(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler retornou status %v em vez de %v", status, http.StatusOK)
	}

	var response models.ConversionResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Erro ao decodificar resposta JSON: %v", err)
	}

	// Verifica se os campos da resposta estão corretos
	if response.API != "CoinSnark" {
		t.Errorf("API incorreta. Esperava 'CoinSnark', obteve '%s'", response.API)
	}
	if response.APIDocumentation != "https://rmottanet.gitbook.io/coinsnark/" {
		t.Errorf("URL de documentação da API incorreta. Esperava 'https://rmottanet.gitbook.io/coinsnark/', obteve '%s'", response.APIDocumentation)
	}
	if _, ok := response.Conversion["from"]; !ok || response.Conversion["from"] != "USD" {
		t.Errorf("Campo 'from' na conversão está incorreto. Esperava 'USD'")
	}
	if _, ok := response.Conversion["to"]; !ok || response.Conversion["to"] != "EUR" {
		t.Errorf("Campo 'to' na conversão está incorreto. Esperava 'EUR'")
	}
	if _, ok := response.Conversion["converted"]; !ok {
		t.Error("Campo 'converted' na conversão está ausente")
	}
	if _, ok := response.Conversion["cache_updated"]; !ok {
		t.Error("Campo 'cache_updated' na conversão está ausente")
	}
}
