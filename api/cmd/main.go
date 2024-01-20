package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	// Criar um roteador
	router := mux.NewRouter()

	// Configurar rotas
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome Coin Snark!")
	})

	// Adicionar middleware CORS
	handler := cors.Default().Handler(router)

	// Iniciar o servidor
	fmt.Println("Servidor rodando na porta :8000...")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
