package config

import (
    "os"
    "log"
    "github.com/joho/godotenv"
)


var EnvVars map[string]string


func LoadEnv() {
    EnvVars = make(map[string]string)

    // check .env
    if _, err := os.Stat(".env"); err != nil {
        if os.IsNotExist(err) {
            log.Println("File .env not found, using environment variables from service")
            return
        }
        log.Fatalf("Error checking .env file: %v", err)
    }


    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    EnvVars["ExchangeRatesKey"] = os.Getenv("EXCHANGE_RATES_API_KEY")
    EnvVars["FixerKey"] = os.Getenv("FIXER_API_KEY")
    EnvVars["OpenExchangeKey"] = os.Getenv("OPEN_EXCHANGE_API_KEY")
    EnvVars["port"] = os.Getenv("PORT")
}
