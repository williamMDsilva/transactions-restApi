package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/williamMDsilva/transactions-restApi/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Port %s", port)

	routes := routes.RoutesApis()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), routes))
}
