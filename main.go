package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Printf("PORT: [%v]\n", port)
	if port == "" {
		log.Fatal("PORT env variable is not set")
	}

	// apiConfig := apiConfig{}

	server := http.Server{
		Addr:        ":" + port,
		Handler:     http.NewServeMux(),
		ReadTimeout: time.Duration(5 * time.Second),
	}

	fmt.Println("Hello world")

	server.Handler.Handle("/", http.FileServer(http.Dir("static/")))
	fmt.Printf("Booting up on http://localhost%v\n", server.Addr)
	err = http.ListenAndServe(server.Addr, server.Handler)
	if err != nil {
		fmt.Println(err)
	}
}
