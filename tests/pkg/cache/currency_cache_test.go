package cache_test

import (
    "testing"
    "time"

    "coinsnark/api/pkg/cache"
)

func TestCurrencyCache_GetSet(t *testing.T) {
    c := cache.NewCurrencyCache()

    c.Set("USD", 1.0, 1*time.Hour)

    value, ok := c.Get("USD")

    if !ok || value != 1.0 {
        t.Errorf("Erro ao obter valor da chave USD: esperado 1.0, obtido %f", value)
    }
}

func TestCurrencyCache_GetAll(t *testing.T) {
    c := cache.NewCurrencyCache()

    c.Set("USD", 1.0, 1*time.Hour)
    c.Set("EUR", 0.85, 1*time.Hour)

    currencies, _, err := c.GetAll()

    if err != nil {
        t.Errorf("Error retrieving all cache values: %v", err)
    }

    if len(currencies) != 2 || currencies["USD"] != 1.0 || currencies["EUR"] != 0.85 {
        t.Errorf("Incorrect cache values: %v", currencies)
    }
}

func TestCurrencyCache_GetTimestamp(t *testing.T) {
    c := cache.NewCurrencyCache()

    initialTimestamp := c.GetTimestamp()

    c.Set("USD", 1.0, 1*time.Hour)

    newTimestamp := c.GetTimestamp()

    if newTimestamp.Before(initialTimestamp) {
        t.Error("The new timestamp must be after the initial timestamp")
    }
}
