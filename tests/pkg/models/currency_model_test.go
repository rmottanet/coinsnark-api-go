package models_test

import (
    "testing"
    "time"

    "coinsnark/api/pkg/models"
)

func TestNewCurrencyResponse(t *testing.T) {
    // Definir valores de exemplo
    currencies := map[string]string{
        "USD": "United States Dollar",
        "EUR": "Euro",
        "GBP": "British Pound Sterling",
    }
    timestamp := time.Now()

    // Criar uma nova instância de CurrencyResponse
    response := models.NewCurrencyResponse(currencies, timestamp)

    // Verificar se os valores foram definidos corretamente na instância
    if response.API != "CoinSnark" {
        t.Errorf("API não definida corretamente: esperava 'CoinSnark', obteve '%s'", response.API)
    }
    if response.Timestamp != timestamp.Format(time.RFC3339) {
        t.Errorf("Timestamp não definido corretamente: esperava '%s', obteve '%s'", timestamp.Format(time.RFC3339), response.Timestamp)
    }
    // Verificar se o mapa de moedas está correto
    for code, name := range currencies {
        if response.Currencies[code] != name {
            t.Errorf("Nome da moeda para código '%s' não definido corretamente: esperava '%s', obteve '%s'", code, name, response.Currencies[code])
        }
    }
    // Você pode verificar outros campos conforme necessário
}
