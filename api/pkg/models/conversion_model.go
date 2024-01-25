package models

import (
    "time"
    
    "coinsnark/api/pkg/util"
)


type ConversionRequest struct {
    From   string  `json:"from"`
    To     string  `json:"to"`
    Amount float64 `json:"amount"`
}


type ConversionResponse struct {
    API					string		`json:"api"`
    APIDocumentation	string		`json:"api_documentation"`
    CacheUpdated		string		`json:"cache_updated"`
    From				string		`json:"from"`
    To					string		`json:"to"`
    Converted			string		`json:"converted"`
    License				string		`json:"license"`
    TermsOfUse			string		`json:"terms_of_use"`
}


func NewConversionResponse(from, to string, convertedAmount float64, cacheUpdated time.Time) *ConversionResponse {

    convertedAmountStr := util.FormatAmount(convertedAmount)

    return &ConversionResponse{
        API:              "CoinSnark",
        APIDocumentation: "https://rmottanet.gitbook.io/coinsnark",
        From:             from,
        To:               to,
        Converted:        convertedAmountStr,
        CacheUpdated:     cacheUpdated.Format(time.RFC3339),
        License:          "https://raw.githubusercontent.com/rmottanet/profitability/main/LICENSE",
        TermsOfUse:       "https://rmottanet.gitbook.io/coinsnark/coin-snark/coin-snak-api-terms-of-use",
    }
}
