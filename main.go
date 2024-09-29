package main

import (
	"fmt"
	"inviter/config"
	"inviter/handlers"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Load configuration
	config.Load()

	// Handle form submission
	http.HandleFunc("/submit", handlers.Submit)

	// Serve the success page
	http.HandleFunc("/success", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./static/success.html")
	})

	// Serve the main page
	http.HandleFunc("/", handlers.MainPage)

	err := handlers.InitCache()
	if err != nil {
		log.Fatalf("Error initializing cache: %v", err)
	}

	if config.IsTlsEnable() {
		go func() {
			// Start HTTP server that redirects all traffic to HTTPS
			log.Println("Starting HTTP to HTTPS redirect")
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort()), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Redirect to HTTPS
				index := strings.Index(r.Host, ":")
				target := fmt.Sprintf("https://%s:%s%s", r.Host[:index], config.HttpsPort(), r.RequestURI)
				http.Redirect(w, r, target, http.StatusMovedPermanently)
			})))
		}()

		// Start HTTPS server
		fmt.Println("Server is running on https://127.0.0.1:" + config.HttpsPort())
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", config.HttpsPort()), config.TlsCert(), config.TlsKey(), nil))
	} else {
		fmt.Println("Server is running on http://127.0.0.1:" + config.HttpPort())
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort()), nil))
	}
}
