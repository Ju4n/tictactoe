package tictactoe

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	//"errors"

	"net/http"
)

type Request struct {
	X int `json:"x"`
	Y int `json:"y"`
	Token string `json:"token"`
}

type Game struct {
	Finished bool `json:"finished"`
	Winner string `json:"winner"`
}

type State struct {
	Positions []string `json:"positions"`
}

// board
var emptyBoard [3][3]string
var board = emptyBoard

func GetState(w http.ResponseWriter, r *http.Request) {
	data := make([]string, 9)
	count := 0;
	for index, line := range board {
		for square,token := range line {
			if (token != "") {
				data[count] = strconv.Itoa(index) + "_" + strconv.Itoa(square) + "-" + token
				count++
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(State{Positions: data}); err != nil {
		panic(err)
	}
}

func Verify(w http.ResponseWriter, r *http.Request) {
	var verifyRequest Request
	body, err := io.ReadAll(r.Body)
	json.Unmarshal([]byte(body), &verifyRequest);

	defer r.Body.Close()

	if (err != nil) {
		panic(err)
	}

	fmt.Println("Verify Token")
	finished, token := addAndCheck(verifyRequest)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Game{Finished: finished, Winner: token}); err != nil {
		panic(err)
	}
}

func Restart(w http.ResponseWriter, r *http.Request) {
	board = emptyBoard
	fmt.Println("Reset the board");
}

func addAndCheck(verifyRequest Request) (bool, string) {
	board[verifyRequest.X][verifyRequest.Y] = verifyRequest.Token
	token := verifyRequest.Token;
	if (
		// horizontals
		board[0][0] == token && board[0][1] == token && board[0][2] == token ||
		board[1][0] == token && board[1][1] == token && board[1][2] == token ||
		board[2][0] == token && board[2][1] == token && board[2][2] == token ||
		// verticals
		board[0][0] == token && board[1][0] == token && board[2][0] == token ||
		board[0][1] == token && board[1][1] == token && board[2][1] == token ||
		board[0][2] == token && board[1][2] == token && board[2][2] == token ||
		// diagonals
		board[0][0] == token && board[1][1] == token && board[2][2] == token ||
		board[0][2] == token && board[1][1] == token && board[2][0] == token) {

		fmt.Printf("%s has Won\n", token);

		return true, token
	}

	return false, ""
}
