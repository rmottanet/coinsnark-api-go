package services

import (
    "time"
    "errors"

    "coinsnark/api/pkg/cache"
)


type ConversionServiceInterface interface {
    Convert(from, to string, amount float64) (float64, time.Time, error)
}


type ConversionService struct {
    Cache cache.Cache
}


func NewConversionService(cache cache.Cache) *ConversionService {
    return &ConversionService{
        Cache: cache,
    }
}

// Convert converts the specified amount from one currency to another.
func (convertSrvc *ConversionService) Convert(from, to string, amount float64) (float64, time.Time, error) {

    rateFrom, okFrom := convertSrvc.Cache.Get(from)
    rateTo, okTo := convertSrvc.Cache.Get(to)

    if !okFrom || !okTo {
        return 0, time.Time{}, errors.New("Exchange rates not found in the cache")
    }

    exchangeRate := rateTo / rateFrom

    convertedAmount := amount * exchangeRate

    cacheTimestamp := convertSrvc.Cache.GetTimestamp()

    return convertedAmount, cacheTimestamp, nil
}
