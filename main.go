package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/tamerlang/go-quest-demo/controllers"
	"github.com/tamerlang/go-quest-demo/models"
)

const (
	defaultPort = "8008"

	idleTimeout       = 30 * time.Second
	writeTimeout      = 180 * time.Second
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 10 * time.Second
)

func main() {
  err := godotenv.Load()

  if err != nil {
    log.Fatal("Error loading .env file")
  }

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	handler := controllers.New() 

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: handler,

		IdleTimeout:       idleTimeout,
		WriteTimeout:      writeTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
	}

  models.ConnectDatabase()

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("ERR ListenAndServe:", err)
	}
}
