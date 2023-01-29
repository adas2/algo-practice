package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"adas2.io/practice/my_crypto"
)

var (
	flagPort = flag.String("port", "8088", "Port to listen on")
)

func usage() {
	flag.PrintDefaults()
}

type wrapperStruct struct {
	name string
}

func (ws wrapperStruct) testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, from Host: " + ws.name))
}

// alternate way of wrapping handler funcs
func testHandler2(name string) http.HandlerFunc {
	cert, _ := my_crypto.GenX509Cert()
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Generated cert for: " + name + "\n" + string(cert)))
	}
}

func main() {
	http.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request from %v\n", r.RemoteAddr)
		w.Write([]byte("hello\n"))
	})
	// http.HandleFunc("/cert", getCertificate)

	handlers := wrapperStruct{name: "Localhost"}
	http.HandleFunc("/test", handlers.testHandler)

	http.HandleFunc("/gen", testHandler2("Localhost"))

	fmt.Printf("Starting server for testing HTTP POST...\n")
	log.Fatal(http.ListenAndServe(":"+*flagPort, nil))
}
