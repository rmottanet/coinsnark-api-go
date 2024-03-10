package controllers

import (
    "net/http"
    "strings"
    "strconv"
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
        // Invalid request error (400)
        errorResponse := models.NewErrorResponse("Error converting amount value", http.StatusBadRequest)
        jsonResponse, _ := json.Marshal(errorResponse)
        w.WriteHeader(http.StatusBadRequest)
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResponse)
        return
    }

    convertedAmount, err := convertCtrl.Service.Convert(from, to, amount)
    if err != nil {
        // Internal server error (500)
        errorResponse := models.NewErrorResponse("Error converting.", http.StatusInternalServerError)
        jsonResponse, _ := json.Marshal(errorResponse)
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResponse)
        return
    }

    response := models.NewConversionResponse(convertedAmount, time.Now())

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        // Internal server error (500)
        errorResponse := models.NewErrorResponse("Error formatting response", http.StatusInternalServerError)
        jsonResponse, _ = json.Marshal(errorResponse)
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResponse)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}
