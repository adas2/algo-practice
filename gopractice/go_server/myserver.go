package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"adas2.io/practice/my_crypto"
)

// command line parameters
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
		w.Write([]byte(fmt.Sprintf("X509 certificate generated at host: %s for host:port %s\n%s\n", name, r.RemoteAddr, string(cert))))
	}
}

func main() {
	// inline http response handler
	http.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request from %v\n", r.RemoteAddr)
		w.Write([]byte("hello\n"))
	})

	// using a handler struct with state
	hostname, _ := os.Hostname()
	handlers := wrapperStruct{name: hostname}
	http.HandleFunc("/test", handlers.testHandler)

	// another way using a wrapping func
	http.HandleFunc("/gen", testHandler2(hostname))

	flag.Parse()

	fmt.Printf("Starting go server at port %s...\n", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, nil))
}
