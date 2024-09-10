package main

import (
	"fmt"
	"inviter/config"
	"inviter/handlers"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	tlsEnable := config.Load()

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle form submission
	http.HandleFunc("/submit", handlers.Submit)

	// Serve the main page
	http.HandleFunc("/", handlers.MainPage)

	if tlsEnable {
		fmt.Println("Server is running on https://127.0.0.1:" + config.Port())
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", config.Port()), config.TlsCert(), config.TlsKey(), nil))

	} else {
		fmt.Println("Server is running on http://127.0.0.1:" + config.Port())
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port()), nil))
	}
}
