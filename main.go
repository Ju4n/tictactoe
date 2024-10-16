package main

import (
	"log"
	"net/http"

	"github.com/Ju4n/tictactoe/tictactoe"
)

func main () {
	http.HandleFunc("/api/restart", tictactoe.Restart)
	http.HandleFunc("/api/verify", tictactoe.Verify)
	http.HandleFunc("/api/state", tictactoe.GetState)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil));
}


