package router

import (
	"kaffein/internal/handler"
	"kaffein/internal/handler/cryptography_handler"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Handle static files with correct stripping of the "/resources/" prefix
	mux.Handle("/resources/", http.StripPrefix("/resources/", handler.ResourcesHandler()))

	mux.HandleFunc("/", cryptography_handler.HandlerEncrypt)
	mux.HandleFunc("/decrypt", cryptography_handler.HandlerDecrypt)
	mux.HandleFunc("/information_capacity", cryptography_handler.HandlerGetCapacity)

	return mux
}
