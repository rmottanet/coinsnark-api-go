package util

import (
    "net/http"
    "encoding/json"

    "coinsnark/api/pkg/models"
)

// writeErrorResponse escreve o ErrorResponse na resposta HTTP
func WriteErrorResponse(w http.ResponseWriter, errorResponse *models.ErrorResponse) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(errorResponse)
}
