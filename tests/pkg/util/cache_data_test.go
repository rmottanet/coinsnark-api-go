package util_test

import (
    "testing"

    "coinsnark/api/pkg/cache"
    "coinsnark/api/pkg/util"
)

func TestSaveDataToCache(t *testing.T) {
    c := cache.NewCurrencyCache()

    data := map[string]float64{
        "USD": 1.0,
        "EUR": 0.85,
        "GBP": 0.73,
    }

    util.SaveDataToCache(data, c)

    for key, value := range data {
        cacheValue, found := c.Get(key)
        if !found {
            t.Errorf("Key %s not found in cache", key)
        }
        if cacheValue != value {
            t.Errorf("Stored value for key %s in cache %f, expected %f", key, cacheValue, value)
        }
    }
}
