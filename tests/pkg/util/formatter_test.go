package util_test

import (
	"testing"
	
	"coinsnark/api/pkg/util"
)

func TestFormatAmount(t *testing.T) {
    amount := 10.5050548
    expected := "10.51"
    result := util.FormatAmount(amount)

    if result != expected {
        t.Errorf("FormatAmount(%.4f) = %s; wanted %s", amount, result, expected)
    }
}
