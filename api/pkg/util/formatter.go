package util

import "fmt"


func FormatAmount(amount float64) string {
    return fmt.Sprintf("%.2f", amount)
}


