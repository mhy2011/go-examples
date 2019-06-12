package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-examples/monitor"
	"log"
	"math/rand"
	"net/http"
	"time"
)


func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})

	http.HandleFunc("/hello", monitor.Monitor(hello))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(writer http.ResponseWriter, request *http.Request)  {
	// 模拟耗时
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(duration)
	fmt.Fprintf(writer, "你好%v", duration)
}