package controllers

import (
    "strings"
    "strconv"
    "net/http"
	"time"
    "encoding/json"

    "coinsnark/api/pkg/services"
    "coinsnark/api/pkg/models"
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

    convertedAmount, err := convertCtrl.Service.Convert(from, to, amount)
    if err != nil {
        http.Error(w, "Erro ao realizar a conversão", http.StatusInternalServerError)
        return
    }

    response := models.NewConversionResponse(convertedAmount, time.Now())

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error formatting response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}
