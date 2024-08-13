package main

import (
	"fmt"
	"kaffein/config"
	"kaffein/internal/handler"
	"kaffein/internal/router"
	"net/http"
)

func main() {
	// GET SERVER ADDRESS CONFIG
	cfg, err := config.DefaultConfig()
	if err != nil {
		panic(err)
	}

	// GET SERVER ADDRESS CONFIG
	address := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	fmt.Printf("Server started at %s\n", address)

	// START ROUTE
	r := router.NewRouter()

	// ENABLE CORS
	cors := handler.EnableCors(r)

	// START SERVER ADDRESS
	err = http.ListenAndServe(address, cors)
	if err != nil {
		panic(err)
	}
}
