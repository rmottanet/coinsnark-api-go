package models

import (
    "time"
)


type ApiResponse struct {
    API              string `json:"api"`
    APIDocumentation string `json:"api_documentation"`
    License          string `json:"license"`
    TermsOfUse       string `json:"terms_of_use"`
    Timestamp        string `json:"timestamp"`
}


func NewApiResponse() *ApiResponse {
    return &ApiResponse{
        API:              "CoinSnark",
        APIDocumentation: "https://rmottanet.gitbook.io/coinsnark/",
        License:          "https://raw.githubusercontent.com/rmottanet/coinsnark-api-go/main/LICENSE",
        TermsOfUse:       "https://rmottanet.gitbook.io/coinsnark/coinsnark/coin-snark-api-terms-of-use",
        Timestamp:        time.Now().Format("2006-01-02T15:04:05"), 
    }
}
