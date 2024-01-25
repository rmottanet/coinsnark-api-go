// tests/pkg/mocks/mock_currency_service.go
package mocks

import (
    "coinsnark/api/pkg/cache"
    "log"
    "coinsnark/api/pkg/util"
)

// MockCurrencyService é uma implementação fictícia do serviço de moeda para testes
type MockCurrencyService struct {
    Cache cache.Cache
}

// NewMockCurrencyService cria uma nova instância de MockCurrencyService com o cache global fornecido
func NewMockCurrencyService(cache cache.Cache) *MockCurrencyService {
    return &MockCurrencyService{
        Cache: cache,
    }
}

// GetSymbolNames obtém as siglas e nomes humanos das moedas do cache global
func (mcs *MockCurrencyService) GetSymbolNames() map[string]string {
    currencyNames := util.CurrencyNames
    currencyRates, _, err := mcs.Cache.GetAll()
    if err != nil {
        log.Printf("Erro ao obter taxas de câmbio do cache: %v", err)
        return nil
    }
    symbolNames := make(map[string]string)
    for currency := range currencyRates {
        if name, ok := currencyNames[currency]; ok {
            symbolNames[currency] = name
        }
    }
    return symbolNames
}
