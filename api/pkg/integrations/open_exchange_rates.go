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


type OpenExchangeRatesResponse struct {
    Disclaimer  string             `json:"disclaimer"`
    License     string             `json:"license"`
    Timestamp   int64              `json:"timestamp"`
    Base        string             `json:"base"`
    Rates       map[string]float64 `json:"rates"`
}


func GetOpenExchangeRatesData(cache cache.Cache) (*OpenExchangeRatesResponse, error) {
	apiKey := os.Getenv("OPEN_EXCHANGE_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("key 'Open Exchange Key' not found in environment variables")
    }
    
    //apiKey, ok := config.EnvVars["OPEN_EXCHANGE_API_KEY"]
    //if !ok {
        //return nil, fmt.Errorf("key 'Open Exchange Key' not found in environment variables")
    //}

    apiURL := GetUrl("open_exchanges_rates")

    params := url.Values{}
    params.Set("app_id", apiKey)
    params.Set("base", "usd")
    params.Set("prettyprint", "false")
    params.Set("show_alternative", "true")

    fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

    data, err := util.FetchData(fullURL)
    if err != nil {
        return nil, fmt.Errorf("error getting data from API: %v", err)
    }

    var exchangeRates OpenExchangeRatesResponse
    err = json.Unmarshal(data, &exchangeRates)
    if err != nil {
        return nil, fmt.Errorf("error decoding JSON response: %v", err)
    }

    util.SaveDataToCache(exchangeRates.Rates, cache)

    log.Println("Open Exchange Rates cached")

    return &exchangeRates, nil
}
