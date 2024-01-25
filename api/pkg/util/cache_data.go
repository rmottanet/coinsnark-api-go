package util

import (
    "time"
    "coinsnark/api/pkg/cache"
)

const defaultExpiration = 6 * time.Hour

func SaveDataToCache(data map[string]float64, cache cache.Cache) {
    for key, value := range data {
        cache.Set(key, value, defaultExpiration)
    }
}
