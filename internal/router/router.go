package router

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	// property assets
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// add route pattern
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, world"
		w.Write([]byte(msg)) // Write response to the client
	})

	return mux
}
