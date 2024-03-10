package services

import (
	"strconv"
    "time"
    "errors"

    "coinsnark/api/pkg/cache"
)


type ConversionServiceInterface interface {
    Convert(from, to string, amount float64) (map[string]string, error)
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
func (convertSrvc *ConversionService) Convert(from, to string, amount float64) (map[string]string, error) {

    rateFrom, okFrom := convertSrvc.Cache.Get(from)
    rateTo, okTo := convertSrvc.Cache.Get(to)
    if !okFrom || !okTo {
        return nil, errors.New("Exchange rates not found in the cache")
    }

    exchangeRate := rateTo / rateFrom

    convertedAmount := amount * exchangeRate

    cacheTimestamp := convertSrvc.Cache.GetTimestamp()

    response := make(map[string]string)
    response["from"] = from
    response["to"] = to
    response["converted"] = strconv.FormatFloat(convertedAmount, 'f', 2, 64)
    response["cache_updated"] = cacheTimestamp.Format(time.RFC3339)

    return response, nil
}
