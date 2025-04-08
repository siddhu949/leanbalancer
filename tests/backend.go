package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "✅ Hello from backend %d!\n", port)
	})

	addr := fmt.Sprintf(":%d", port)
	log.Printf("✅ Backend running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func main() {
	go handler(9001)
	go handler(9002)
	handler(9003) // blocking, so the main goroutine stays alive
}
