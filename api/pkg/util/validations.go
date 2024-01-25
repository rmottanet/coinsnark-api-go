package util

import (
    "regexp"
    "strconv"
)

func ValidateCurrencyCode(code string) bool {
    regex := regexp.MustCompile(`^[a-zA-Z]{3}$`)
    return regex.MatchString(code)
}

func ValidateAmount(amountStr string) bool {
    regex := regexp.MustCompile(`^\d+(\.\d{1,2})?$`)
    if !regex.MatchString(amountStr) {
        return false
    }

	// valid amount number
    _, err := strconv.ParseFloat(amountStr, 64)
    if err != nil {
        return false
    }

    return true
}
