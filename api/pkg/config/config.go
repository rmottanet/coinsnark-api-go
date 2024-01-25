package config

import (
    "os"
    "log"
    "github.com/joho/godotenv"
)


var EnvVars map[string]string


func LoadEnv() {
    EnvVars = make(map[string]string)

    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error uploading file .env")
    }

    EnvVars["ExchangeRatesKey"] = os.Getenv("EXCHANGE_RATES_API_KEY")
    EnvVars["FixerKey"] = os.Getenv("FIXER_API_KEY")
    EnvVars["OpenExchangeKey"] = os.Getenv("OPEN_EXCHANGE_API_KEY")
    EnvVars["port"] = os.Getenv("PORT")
}
