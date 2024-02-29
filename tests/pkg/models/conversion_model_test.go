package models_test

import (
    "testing"
    "time"

    "coinsnark/api/pkg/models"
)

func TestNewConversionResponse(t *testing.T) {

    conversion := map[string]string{
        "USD": "0.85",
        "EUR": "1.00",
    }
    cacheUpdated := time.Now()

    response := models.NewConversionResponse(conversion, cacheUpdated)

    if response.API != "CoinSnark" {
        t.Errorf("API não definida corretamente: esperava 'CoinSnark', obteve '%s'", response.API)
    }
    if response.Conversion["USD"] != "0.85" {
        t.Errorf("Valor de conversão 'USD' não definido corretamente: esperava '0.85', obteve '%s'", response.Conversion["USD"])
    }
    if response.Conversion["EUR"] != "1.00" {
        t.Errorf("Valor de conversão 'EUR' não definido corretamente: esperava '1.00', obteve '%s'", response.Conversion["EUR"])
    }
    if response.Timestamp != cacheUpdated.Format(time.RFC3339) {
        t.Errorf("Valor 'Timestamp' não definido corretamente: esperava '%s', obteve '%s'", cacheUpdated.Format(time.RFC3339), response.Timestamp)
    }

}
