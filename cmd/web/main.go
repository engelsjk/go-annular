package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	goannular "github.com/engelsjk/go-annular"
)

type Handler struct {
	annular *goannular.Annular
}

func (h *Handler) SVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	h.annular.Draw()
	if err := h.annular.Render(w, "svg"); err != nil {
		log.Println(err.Error())
	}
}

func (h *Handler) PNG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	h.annular.Draw()
	if err := h.annular.Render(w, "png"); err != nil {
		log.Println(err.Error())
	}
}

func main() {

	annular := goannular.NewAnnular()

	handler := &Handler{annular: annular}

	http.Handle("/png", http.HandlerFunc(handler.PNG))
	http.Handle("/svg", http.HandlerFunc(handler.SVG))

	port := "2003"

	fmt.Printf("listening at http://localhost:%s/svg and listening at http://localhost:%s/png\n", port, port)
	err := http.ListenAndServe(net.JoinHostPort("", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
