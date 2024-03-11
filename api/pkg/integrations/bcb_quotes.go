package integrations

import (
	"fmt"
	"encoding/json"
	"time"
    "log"
    
	"coinsnark/api/pkg/cache"
	"coinsnark/api/pkg/util"
)


type BCBQuotesResponse struct {
	Moeda         string    `json:"moeda"`
	ValorCompra   float64   `json:"valorCompra"`
	DataIndicador time.Time `json:"dataIndicador"`
	TipoCotacao   string    `json:"tipoCotacao"`
}


func GetBCBQuotesData(cache *cache.CurrencyCache) (*map[string]float64, error) {
	
	apiURL := GetUrl("bcb_quotes")

	data, err := util.FetchData(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error getting data from API: %v", err)
	}

	var responseData struct {
		Conteudo []BCBQuotesResponse `json:"conteudo"`
	}
	err = json.Unmarshal(data, &responseData)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	currencyRates := make(map[string]float64)
	for _, quote := range responseData.Conteudo {
		switch quote.Moeda {
		case "Dólar":
			currencyRates["USD"] = 1.0
			currencyRates["BRL"] = quote.ValorCompra
		case "Euro":
			currencyRates["EUR"] = quote.ValorCompra
		}
	}

	brl_to_eur := 1 / currencyRates["EUR"]
	usd_to_eur := currencyRates["BRL"] * brl_to_eur

	exchangeRates := map[string]float64{
		"USD": 1.0,
		"BRL": currencyRates["BRL"],
		"EUR": usd_to_eur,
	}
	util.SaveDataToCache(exchangeRates, cache)

    log.Println("BCB Quotes cached")
    
	return &exchangeRates, nil
}
