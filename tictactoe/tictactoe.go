package tictactoe

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	X int `json:"x"`
	Y int `json:"y"`
	Token string `json:"token"`
}

type Response struct {
	Finished string `json:"finished"`
	Winner string `json:"winner"`
}


func Verify(w http.ResponseWriter, r *http.Request) {
	var jrequest Request
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), &jrequest);

	defer r.Body.Close()

	if (err != nil) {
		panic(err)
	}

	game(jrequest.X, jrequest.Y, jrequest.Token)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Response{Finished: "true", Winner: "X"}); err != nil {
		panic(err)
	}
}

func game(x int, y int, token string) (string, error) {
	
}
