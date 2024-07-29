package main

import (
	"fmt"
	"kaffein/config"
	"kaffein/internal/router"
	"net/http"
)

func main() {
	cfg := config.DefaultConfig()

	// START ROUTE
	r := router.NewRouter()

	address := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	fmt.Printf("Server started at %s\n", address)
	http.ListenAndServe(address, r)
}
