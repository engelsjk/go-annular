package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	goannular "github.com/engelsjk/go-annular"
)

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	port := "2003"
	fmt.Printf("listening at http://localhost:%s\n", port)
	err := http.ListenAndServe(net.JoinHostPort("", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	goannular.Run(w)
}
