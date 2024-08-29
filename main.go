package main

import (
	"log"
	"net/http"

	"github.com/Ju4n/tictactoe/tictactoe"
)

func main () {
	http.HandleFunc("/api/verify", tictactoe.Verify)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil));
}


