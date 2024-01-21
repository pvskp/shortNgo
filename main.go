package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/spaolacci/murmur3"
)

// generateHash generates a hash for a given URL
func generateHash(url string) string {
	hash := murmur3.New32()
	hash.Write([]byte(url))
	return hex.EncodeToString(hash.Sum(nil))
}

// processRequest processes the request
func processRequest(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if isValidURL(url) {
		hash := generateHash(url)
		fmt.Fprintf(w, "Hash generated: %s", hash)
		return
	}
	fmt.Fprintln(w, "The provided URL is not valid")
}

// isValidURL verifica se a string é uma URL válida
func isValidURL(u string) bool {
	parsedURL, err := url.ParseRequestURI(u)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}

func main() {
	http.HandleFunc("/process", processRequest)
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
