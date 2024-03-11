package controllers

import (
    "net/http"
    "encoding/json"

    "coinsnark/api/pkg/models" 
)


type ApiController struct {
}


func NewApiController() *ApiController {
	return &ApiController{}
}


func (apiCtrl *ApiController) ApiInfo(w http.ResponseWriter, r *http.Request) {

    apiResponse := models.NewApiResponse()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
        // Internal server error (500)
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}
