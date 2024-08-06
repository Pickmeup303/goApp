package main

import (
	"fmt"
	"kaffein/config"
	"kaffein/internal/router"
	"net/http"
)

func main() {
	cfg, err := config.DefaultConfig()
	if err != nil {
		panic(err)
	}

	// START ROUTE
	r := router.NewRouter()

	address := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	fmt.Printf("Server started at %s\n", address)
	err = http.ListenAndServe(address, r)
	if err != nil {
		panic(err)
	}
}
