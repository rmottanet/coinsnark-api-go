package integrations

import (
    "fmt"
    "os"
    "log"
    "net/url"
    "encoding/json"

    //"coinsnark/api/pkg/config"
    "coinsnark/api/pkg/cache"
    "coinsnark/api/pkg/util"
)


type ExchangeRatesResponse struct {
    Base        string             `json:"base"`
    LastUpdated int64              `json:"last_updated"`
    ExchangeMap map[string]float64 `json:"exchange_rates"`
}


func GetExchangeRatesData(cache cache.Cache) (*ExchangeRatesResponse, error) {
    apiKey := os.Getenv("EXCHANGE_RATES_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("key 'Exchange Rates Key' not found in environment variables")
    }

    apiURL := GetUrl("exchange_rates")

    params := url.Values{}
    params.Add("api_key", apiKey)
    params.Add("base", "USD")

    fullURL := apiURL + "?" + params.Encode()

    data, err := util.FetchData(fullURL)
    if err != nil {
        return nil, fmt.Errorf("error getting data from API: %v", err)
    }

    var exchangeRates ExchangeRatesResponse
    err = json.Unmarshal(data, &exchangeRates)
    if err != nil {
        return nil, fmt.Errorf("error decoding JSON response: %v", err)
    }

    util.SaveDataToCache(exchangeRates.ExchangeMap, cache)

    log.Println("Exchange Rates cached")

    return &exchangeRates, nil
}
