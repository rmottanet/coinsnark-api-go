package middleware

import (
    "net/http"
    
    "coinsnark/api/pkg/util"
)


func ValidateInput(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        from := r.URL.Query().Get("from")
        to := r.URL.Query().Get("to")
        amount := r.URL.Query().Get("amount")

        if !util.ValidateCurrencyCode(from) || !util.ValidateCurrencyCode(to) || !util.ValidateAmount(amount) {
            http.Error(w, "Invalid input data", http.StatusBadRequest)
            return
        }

        next.ServeHTTP(w, r)
    })
}
