package main

import (
	"os"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/gorilla/mux"

	"coinsnark/api/pkg/cache"
	"coinsnark/api/pkg/config"
	"coinsnark/api/pkg/integrations"
	"coinsnark/api/pkg/middleware"
	"coinsnark/api/pkg/controllers"
	"coinsnark/api/pkg/services"	
	"coinsnark/api/pkg/html"
)

func main() {

	config.LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	corsHandler := cors.New(cors.Options{
	    AllowedOrigins:   []string{"*"},
	    AllowedMethods:   []string{"GET"},
	    AllowedHeaders:   []string{"Content-Type"},
	    AllowCredentials: true,
	})

	exchangeRatesCache := cache.NewCurrencyCache()
	
	router := mux.NewRouter()

	router.Use(corsHandler.Handler)

	currencyService := services.NewCurrencyService(exchangeRatesCache)
	conversionService := services.NewConversionService(exchangeRatesCache)

	currencyController := controllers.NewCurrencyController(currencyService)
	conversionController := controllers.NewConversionController(conversionService)
	

	router.HandleFunc("/api/currency", currencyController.GetCurrencyNames).Methods("GET")

    router.Handle("/api/convert", middleware.ValidateInput(middleware.ConvertHandler(conversionController.ConvertCurrency)))


	if _, err := integrations.GetOpenExchangeRatesData(exchangeRatesCache); err != nil {
		log.Fatalf("Error retrieving exchange rate data: %v", err)
	}

	if _, err := integrations.GetExchangeRatesData(exchangeRatesCache); err != nil {
		log.Fatalf("Error retrieving exchange rate data: %v", err)
	}
	
	
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        welcomeHTML := html.WelcomePageHTML()

        w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(welcomeHTML))
    })
    
	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
