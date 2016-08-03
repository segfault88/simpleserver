package main

import (
	"io"
	"log"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "hello, world!\n")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
