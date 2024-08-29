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
	board := [3][3]string{}

	if board[x][y] == "" {
		board[x][y] = token
	} else {
		return "", errors.New("invalid move, cell already taken")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// horizontal
			fmt.Println(board[i][j])
			// vertical
			fmt.Println(board[j][i])
		}
		// diagonal
		fmt.Println(board[i][i])
		// diagonal invertida
		fmt.Println(board[i][2-i])
	}

	return token, nil
}
