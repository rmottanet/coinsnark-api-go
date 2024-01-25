package services_test

import (
    "testing"
    "time"

    "coinsnark/api/pkg/cache"
    "coinsnark/api/pkg/services"
    "coinsnark/api/pkg/util"
)

func TestCurrencyService_GetSymbolNames(t *testing.T) {
    c := cache.NewCurrencyCache()

    c.Set("USD", 1.0, 1*time.Hour)
    c.Set("EUR", 0.85, 1*time.Hour)
    c.Set("GBP", 0.73, 1*time.Hour)

    currencySrvc := services.NewCurrencyService(c)

    symbolNames := currencySrvc.GetSymbolNames()

    for symbol := range symbolNames {
        if _, ok := util.CurrencyNames[symbol]; !ok {
            t.Errorf("Symbol %s not found in currency names", symbol)
        }
    }

    for symbol, name := range symbolNames {
        if util.CurrencyNames[symbol] != name {
            t.Errorf("Human name %s does not match the symbol %s", name, symbol)
        }
    }
}
