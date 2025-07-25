package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

// logRequest logs all relevant details of an incoming HTTP request
func logRequest(r *http.Request) {
	// Method and URL
	log.Println("----- New Request ------")
	log.Printf("Method: %s", r.Method)
	log.Printf("URL: %s", r.URL.String())
	log.Printf("Path: %s", r.URL.Path)

	// Headers
	log.Println("[Headers]")
	for name, values := range r.Header {
		for _, value := range values {
			log.Printf("\t%s: %s", name, value)
		}
	}

	// Query parameters
	log.Println("[Query Parameters]")
	for name, values := range r.URL.Query() {
		for _, value := range values {
			log.Printf("\t%s: %s", name, value)
		}
	}

	// Body (if any)
	if r.Body != nil {
		defer r.Body.Close()
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
		} else if len(bodyBytes) > 0 {
			log.Println("[Body]")
			bodyStr := string(bodyBytes)
			// Optional: Trim long body
			if len(bodyStr) > 1000 {
				bodyStr = bodyStr[:1000] + "...(truncated)"
			}
			log.Println("\t", bodyStr)
			// Replace the body so it can be read again
			r.Body = io.NopCloser(strings.NewReader(string(bodyBytes)))
		}
	}
	log.Println("------------------------")
}

// requestHandler processes incoming requests
func requestHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request received and logged.\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", requestHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown setup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("Starting server on :8080 \n---------------------------------------------")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt
	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped.")
}
