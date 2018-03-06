package main

import (
	"net/http"
	"time"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// Create a multiplexer that can redirect request to handler
	mux := http.NewServeMux()

	// Serve static files from a given directory
	files := http.FileServer(http.Dir(config.Static))

	// Pass the handler to the Handle function
	// and remove the given prefix from the request URL's paht
	// Basically look for files tarting inside static folder
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// List of all endpoints
	// Handler functions are defined in different files
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	// Defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// Defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// Start the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
