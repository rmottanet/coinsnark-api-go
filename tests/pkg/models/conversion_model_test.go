package models_test

import (
    "testing"
    "time"

    "coinsnark/api/pkg/models"
)

func TestNewConversionResponse(t *testing.T) {
    // Definir valores de exemplo
    from := "USD"
    to := "EUR"
    convertedAmount := 0.85
    cacheUpdated := time.Now()

    // Criar uma nova instância de ConversionResponse
    response := models.NewConversionResponse(from, to, convertedAmount, cacheUpdated)

    // Verificar se os valores foram definidos corretamente na instância
    if response.API != "CoinSnark" {
        t.Errorf("API não definida corretamente: esperava 'CoinSnark', obteve '%s'", response.API)
    }
    if response.From != from {
        t.Errorf("Valor 'From' não definido corretamente: esperava '%s', obteve '%s'", from, response.From)
    }
    if response.To != to {
        t.Errorf("Valor 'To' não definido corretamente: esperava '%s', obteve '%s'", to, response.To)
    }
    if response.Converted != "0.85" { // O valor 0.85 formatado com duas casas decimais
        t.Errorf("Valor 'Converted' não definido corretamente: esperava '0.8500', obteve '%s'", response.Converted)
    }
    if response.CacheUpdated != cacheUpdated.Format(time.RFC3339) {
        t.Errorf("Valor 'CacheUpdated' não definido corretamente: esperava '%s', obteve '%s'", cacheUpdated.Format(time.RFC3339), response.CacheUpdated)
    }
    // Você pode verificar outros campos conforme necessário
}
