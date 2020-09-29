package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-tpm-shim/tpm_shim"
	"log"
	"net/http"
)

// Cert is the certificate for server
type Cert struct {
	Hostname string `json:"hostname,omitempty"`
	Value    string `json:"value,omitempty"`
}

var certDB []Cert

func getCertificate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting certificate")
	json.NewEncoder(w).Encode(certDB)
}

func createCertificate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	log.Printf("Creating Certificate")
	var cert Cert
	_ = json.NewDecoder(r.Body).Decode(&cert)
	certDB = append(certDB, cert)
	json.NewEncoder(w).Encode(cert)
}

func addCertificate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Adding Certificate")
}

func check() {
	fmt.Println(tpm_shim.IsTpmHost)
}

func help() {
	fmt.Println("This is juts for help")
}
