package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/pvskp/gourl/database"
	"github.com/spaolacci/murmur3"
)

const (
	serverPort = ":8080"
)

var (
	db database.IDatabase
)

// generateHash
func generateHash(url string) string {
	hash := murmur3.New32()
	hash.Write([]byte(url))
	return hex.EncodeToString(hash.Sum(nil))
}

func redirectRequest(w http.ResponseWriter, r *http.Request) {
	hash := strings.TrimPrefix(r.RequestURI, "/go/")
	log.Println("Looking for hash", hash)
	url, err := db.GetHashValue(hash)
	if err != nil {
		fmt.Fprintln(w, "Error redirecting :(")
		return
	}
	fmt.Println(w, "Redirecting to ", url)
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

// processRequest processes the request
func processRequest(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if isValidURL(url) {
		hash := generateHash(url)
		for db.HashExists(hash) {
			hash = generateHash(url)
		}

		fmt.Fprintf(w, "Hash generated: %s", hash)

		err := db.SaveHash(hash, url)
		if err != nil {
			log.Printf("Error saving hash: %s", err)
			fmt.Fprintln(w, "Error generating your url :(")
			return
		}
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
	http.HandleFunc("/go/", redirectRequest)
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(func() error {
		log.Printf("Server started on http://localhost%s", serverPort)
		err := http.ListenAndServe(serverPort, nil)
		if err != nil {
			log.Println("Failed to start server:", err)
			return err
		}
		return nil
	}())
}

func init() {
	var err error
	db, err = database.InitiateDB("BadgerIM")
	if err != nil {
		log.Fatalf("Failed to init db: %v", err)
	}
}
