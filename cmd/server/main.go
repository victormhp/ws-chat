package main

import (
	"flag"
	"log"
	"net/http"
)

type application struct {
	hub *Hub
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "./ui/index.html")
}

func main() {
	addr := flag.String("port", ":8080", "HTTP Network address")
	flag.Parse()

	hub := newHub()
	go hub.run()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /", serveHome)
	mux.HandleFunc("GET /ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
