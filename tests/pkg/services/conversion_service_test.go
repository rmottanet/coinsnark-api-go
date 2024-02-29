package services_test

import (
	"testing"
	"time"

	"coinsnark/api/pkg/cache"
	"coinsnark/api/pkg/services"
)

func TestConversionService_Convert(t *testing.T) {
	c := cache.NewCurrencyCache()

	c.Set("USD", 1.0, 1*time.Hour)
	c.Set("EUR", 0.85, 1*time.Hour)

	convertSrvc := services.NewConversionService(c)

	response, err := convertSrvc.Convert("USD", "EUR", 100.0)
	if err != nil {
		t.Errorf("Erro inesperado durante a conversão: %v", err)
	}

	expectedAmount := "85.00"
	if response["converted"] != expectedAmount {
		t.Errorf("Valor convertido incorreto. Esperado: %s, Obtido: %s", expectedAmount, response["converted"])
	}

	cacheUpdated, err := time.Parse(time.RFC3339, response["cache_updated"])
	if err != nil {
		t.Errorf("Erro ao analisar o timestamp do cache: %v", err)
	}
	if cacheUpdated.IsZero() {
		t.Error("O timestamp do cache não deve ser zero")
	}

	_, err = convertSrvc.Convert("JPY", "EUR", 100.0)
	if err == nil {
		t.Error("Esperava-se um erro ao converter de uma moeda não encontrada")
	}

	_, err = convertSrvc.Convert("USD", "JPY", 100.0)
	if err == nil {
		t.Error("Esperava-se um erro ao converter para uma moeda não encontrada")
	}
}
