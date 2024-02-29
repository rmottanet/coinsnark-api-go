package controllers

import (
    "net/http"
    "time"
    "encoding/json"

    "coinsnark/api/pkg/models"
    "coinsnark/api/pkg/services"
)


type CurrencyController struct {
    Service services.CurrencyServiceInterface
}


func NewCurrencyController(service services.CurrencyServiceInterface) *CurrencyController {
    return &CurrencyController{
        Service: service,
    }
}

// GetCurrencyNames retrieves currency symbols and names and returns them as a JSON response.
func (currencyCtrl *CurrencyController) GetCurrencyNames(w http.ResponseWriter, r *http.Request) {

    symbolsNames := currencyCtrl.Service.GetSymbolNames()

    filteredCurrencies := make(map[string]string)
    for symbol, name := range symbolsNames {
        filteredCurrencies[symbol] = name
    }

    response := models.NewCurrencyResponse(filteredCurrencies, time.Now())

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error formatting response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)   
    w.Write(jsonResponse)
}
