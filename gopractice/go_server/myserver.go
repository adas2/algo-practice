package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	flagPort = flag.String("port", "5000", "Port to listen on")
)

func usage() {
	flag.PrintDefaults()
}

func main() {
	http.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request from %v\n", r.RemoteAddr)
		w.Write([]byte("hello\n"))
	})
	http.HandleFunc("/cert", getCertificate)
	http.HandleFunc("/cert/{hostname}", createCertificate)
	log.Fatal(http.ListenAndServe(":"+*flagPort, nil))
}
