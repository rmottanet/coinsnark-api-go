package models

import (
    "time"
)


type ConversionRequest struct {
    From   string  `json:"from"`
    To     string  `json:"to"`
    Amount float64 `json:"amount"`
}


type ConversionResponse struct {
    API               string            `json:"api"`
    APIDocumentation  string            `json:"api_documentation"`
    Conversion        map[string]string `json:"conversion"`
    License           string            `json:"license"`
    TermsOfUse        string            `json:"terms_of_use"`
    Timestamp         string            `json:"timestamp"`
}


func NewConversionResponse(conversion map[string]string, timestamp time.Time) *ConversionResponse {
    return &ConversionResponse{
        API:              "CoinSnark",
        APIDocumentation: "https://rmottanet.gitbook.io/coinsnark/",
        Conversion:       conversion,
        License:          "https://raw.githubusercontent.com/rmottanet/coinsnark-api-go/main/LICENSE",
        TermsOfUse:       "https://rmottanet.gitbook.io/coinsnark/coin-snark/coin-snak-api-terms-of-use",
        Timestamp:        timestamp.Format(time.RFC3339),
    }
}
