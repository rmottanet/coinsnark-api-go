// tests/pkg/integrations/mock_cache.go

package mocks

import (
    "time"
)

// MockCache é uma implementação fictícia do cache.Cache para testes
type MockCache struct {
    Data map[string]float64
}

func (mc *MockCache) Get(key string) (float64, bool) {
    value, ok := mc.Data[key]
    return value, ok
}

func (mc *MockCache) Set(key string, value float64, expiration time.Duration) {
    mc.Data[key] = value
}

func (mc *MockCache) GetAll() (map[string]float64, time.Time, error) {
    return mc.Data, time.Now(), nil
}

func (mc *MockCache) GetTimestamp() time.Time {
    return time.Now()
}
