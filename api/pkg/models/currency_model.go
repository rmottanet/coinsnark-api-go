package models

import (
    "time"
)


type CurrencyResponse struct {
    API               string            `json:"api"`
    APIDocumentation  string            `json:"api_documentation"`
    Timestamp         string            `json:"timestamp"`
    Currencies        map[string]string `json:"currencies"`
    License           string            `json:"license"`
    TermsOfUse        string            `json:"terms_of_use"`
}


func NewCurrencyResponse(currencies map[string]string, timestamp time.Time) *CurrencyResponse {
	
    return &CurrencyResponse{
        API:              "CoinSnark",
        APIDocumentation: "https://rmottanet.gitbook.io/coinsnark",
        Currencies:       currencies,
        Timestamp:        timestamp.Format(time.RFC3339),
        License:          "https://raw.githubusercontent.com/rmottanet/profitability/main/LICENSE",
        TermsOfUse:       "https://rmottanet.gitbook.io/coinsnark/coin-snark/coin-snak-api-terms-of-use",
    }
}
