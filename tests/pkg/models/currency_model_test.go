package models_test

import (
    "testing"
    "time"

    "coinsnark/api/pkg/models"
)

func TestNewCurrencyResponse(t *testing.T) {

    currencies := map[string]string{
        "USD": "United States Dollar",
        "EUR": "Euro",
        "GBP": "British Pound Sterling",
    }
    timestamp := time.Now()

    response := models.NewCurrencyResponse(currencies, timestamp)


    if response.API != "CoinSnark" {
        t.Errorf("API não definida corretamente: esperava 'CoinSnark', obteve '%s'", response.API)
    }

    if response.Timestamp != timestamp.Format(time.RFC3339) {
        t.Errorf("Timestamp não definido corretamente: esperava '%s', obteve '%s'", timestamp.Format(time.RFC3339), response.Timestamp)
    }

    for code, name := range currencies {
        if response.Currencies[code] != name {
            t.Errorf("Nome da moeda para código '%s' não definido corretamente: esperava '%s', obteve '%s'", code, name, response.Currencies[code])
        }
    }

}
