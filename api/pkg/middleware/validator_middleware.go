package middleware

import (
    "net/http"
    
    "coinsnark/api/pkg/util"
    "coinsnark/api/pkg/models"
)


func ValidateInput(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        from := r.URL.Query().Get("from")
        to := r.URL.Query().Get("to")
        amount := r.URL.Query().Get("amount")

        if !util.ValidateCurrencyCode(from) || !util.ValidateCurrencyCode(to) || !util.ValidateAmount(amount) {
            errorResponse := models.NewErrorResponse("Dados de entrada inválidos", http.StatusBadRequest)
            util.WriteErrorResponse(w, errorResponse)
            return
        }

        next.ServeHTTP(w, r)
    })
}
