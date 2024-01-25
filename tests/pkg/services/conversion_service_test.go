package services_test

import (
	"testing"
	"time"

	"coinsnark/api/pkg/cache"
	"coinsnark/api/pkg/services"
)

func TestConversionService_Convert(t *testing.T) {
	c := cache.NewCurrencyCache()

	c.Set("USD", 1.0, 1*time.Hour)
	c.Set("EUR", 0.85, 1*time.Hour)

	convertSrvc := services.NewConversionService(c)

	convertedAmount, timestamp, err := convertSrvc.Convert("USD", "EUR", 100.0)
	if err != nil {
		t.Errorf("Unexpected error during conversion: %v", err)
	}

	expectedAmount := 100.0 * (0.85 / 1.0)
	if convertedAmount != expectedAmount {
		t.Errorf("Incorrect converted value. Expected: %f, Obtained: %f", expectedAmount, convertedAmount)
	}

	if timestamp.IsZero() {
		t.Error("Cache timestamp must not be zero")
	}

	_, _, err = convertSrvc.Convert("JPY", "EUR", 100.0)
	if err == nil {
		t.Error("An error was expected when converting from a currency not found")
	}

	_, _, err = convertSrvc.Convert("USD", "JPY", 100.0)
	if err == nil {
		t.Error("An error was expected when converting from a currency not found")
	}
}
