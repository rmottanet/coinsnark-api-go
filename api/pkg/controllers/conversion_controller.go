package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "coinsnark/api/pkg/models"
    "coinsnark/api/pkg/services"
)


type ConversionController struct {
    Service services.ConversionServiceInterface
}


func NewConversionController(service services.ConversionServiceInterface) *ConversionController {
    return &ConversionController{
        Service: service,
    }
}


// ConvertCurrency handles the conversion of currency based on the request parameters.
func (convertCtrl *ConversionController) ConvertCurrency(w http.ResponseWriter, r *http.Request) {

    from := strings.ToUpper(r.URL.Query().Get("from"))
    to := strings.ToUpper(r.URL.Query().Get("to"))
    amountStr := r.URL.Query().Get("amount")

    amount, err := strconv.ParseFloat(amountStr, 64)
    if err != nil {
        http.Error(w, "Error converting amount value", http.StatusBadRequest)
        return
    }

    convertedAmount, cacheTimestamp, err := convertCtrl.Service.Convert(from, to, amount)
    if err != nil {
        http.Error(w, "Conversion error", http.StatusInternalServerError)
        return
    }

    response := models.NewConversionResponse(from, to, convertedAmount, cacheTimestamp)

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error formatting response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(http.StatusOK)

    w.Write(jsonResponse)
}
