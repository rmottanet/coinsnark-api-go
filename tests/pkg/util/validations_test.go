package util_test

import (
	"testing"
	
	"coinsnark/api/pkg/util"
)

func TestValidateCurrencyCode(t *testing.T) {
    if !util.ValidateCurrencyCode("USD") {
        t.Errorf("ValidateCurrencyCode(\"USD\") = false; wanted true")
    }

    if util.ValidateCurrencyCode("USDD") {
        t.Errorf("ValidateCurrencyCode(\"USDD\") = true; wanted false")
    }

    if util.ValidateCurrencyCode("US") {
        t.Errorf("ValidateCurrencyCode(\"US\") = true; wanted false")
    }
}

func TestValidateAmount(t *testing.T) {
    if !util.ValidateAmount("10.50") {
        t.Errorf("ValidateAmount(\"10.50\") = false; wanted true")
    }

    if util.ValidateAmount("10.555") {
        t.Errorf("ValidateAmount(\"10.555\") = true; wanted false")
    }

    if util.ValidateAmount("abc") {
        t.Errorf("ValidateAmount(\"abc\") = true; wanted false")
    }
}
