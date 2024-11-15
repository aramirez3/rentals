package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	Addr    string
	Handler *http.ServeMux
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env variable is not set")
	}

	// apiConfig := apiConfig{}

	server := Server{
		Addr:    ":" + port,
		Handler: http.NewServeMux(),
	}

	server.Handler.Handle("/", http.FileServer(http.Dir("static/")))
	server.Handler.HandleFunc("/healthz", handlerHealthcheck)

	fmt.Printf("Booting up on http://localhost%v\n", server.Addr)
	err = http.ListenAndServe(server.Addr, middlewareServer(server.Handler))
	if err != nil {
		fmt.Println(err)
	}
}
