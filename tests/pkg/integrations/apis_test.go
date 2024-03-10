package integrations_test

import (
	"testing"

	"coinsnark/api/pkg/integrations"
)

func TestGetUrl(t *testing.T) {
	testCases := map[string]string{
		"bcb_quotes":  "https://www.bcb.gov.br/api/servico/sitebcb/indicadorCambio",
		"exchange_rates":  "https://exchange-rates.abstractapi.com/v1/live/",
		"fixer":  "https://api.apilayer.com/fixer/latest",
		"open_exchanges_rates":  "https://openexchangerates.org/api/latest.json",
	}

	for index, expectedURL := range testCases {
		url := integrations.GetUrl(index)

		if url != expectedURL {
			t.Errorf("For the index %s, the URL returned is %s, expected %s", index, url, expectedURL)
		}
	}
}
