package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
