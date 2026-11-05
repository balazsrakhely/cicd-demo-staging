package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// port returns the port to listen on, from PORT, falling back to 8080.
// Twelve-factor config: the same build runs anywhere, told from outside.
func port() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}

// newMux builds the router. Session 1 registered these handlers on the
// default mux inside main; pulling them into their own function is the one
// change CI needs, because a test can only reach a handler it can construct.
func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello from my container")
	})
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	return mux
}

func main() {
	p := port()
	log.Printf("listening on :%s", p)
	log.Fatal(http.ListenAndServe(":"+p, newMux()))
}
