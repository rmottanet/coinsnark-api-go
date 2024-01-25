package services

import (
	"log"
	
    "coinsnark/api/pkg/cache"
    "coinsnark/api/pkg/util"
    "coinsnark/api/pkg/integrations"
)


type CurrencyServiceInterface interface {
    GetSymbolNames() map[string]string
}


type CurrencyService struct {
    Cache cache.Cache
}


func NewCurrencyService(cache cache.Cache) *CurrencyService {
    return &CurrencyService{
        Cache: cache,
    }
}

// GetSymbolNames returns a map of currency symbols and names.
func (currencySrvc *CurrencyService) GetSymbolNames() map[string]string {

    currencyNames := util.CurrencyNames

	currencyRates, _, err := currencySrvc.Cache.GetAll()
	if err != nil {
		log.Printf("Error getting exchange rates from cache: %v", err)
	}

	// Check if the cache is empty or if there was an error retrieving rates.
	if len(currencyRates) == 0 || err != nil {

		if _, err := integrations.GetOpenExchangeRatesData(currencySrvc.Cache); err != nil {
			log.Fatalf("Error retrieving exchange rate data: %v", err)
		}

		if _, err := integrations.GetExchangeRatesData(currencySrvc.Cache); err != nil {
			log.Fatalf("Error retrieving exchange rate data: %v", err)
		}

		currencyRates, _, err = currencySrvc.Cache.GetAll()
		if err != nil {
			log.Printf("Error getting exchange rates from cache after update: %v", err)
			return nil
		}
	}
	    
    symbolNames := make(map[string]string)

    for currency := range currencyRates {
        if name, ok := currencyNames[currency]; ok {
            symbolNames[currency] = name
        }
    }

    return symbolNames
}
