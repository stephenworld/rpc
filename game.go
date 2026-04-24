package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"stephen/game"
	"stephen/utils"
	"strings"
)

type GAME struct {
	gameID  string
	player1 string
	player2 string
}

func Game() {
	fmt.Println("Welcome to Rockky")

	fmt.Println("1. Create Game (Host)")
	fmt.Println("2. Join Game (Search ID)")

	scanner := bufio.NewScanner(os.Stdin)
	var method string

	for {
		fmt.Print("Selection: ")
		if !scanner.Scan() {
			return
		}
		method = strings.TrimSpace(scanner.Text())
		if method == "1" || method == "2" {
			break
		}
		fmt.Println("Kindly choose 1 or 2")
	}
	utils.ClearScreen()

	var conn net.Conn
	var address string
	var isHost bool

	switch method {
	case "1":
		conn, address = game.CreateGame()
		isHost = true
	case "2":
		conn, address = game.JoinGame()
		isHost = false
	}

	if conn == nil {
		fmt.Println("Failed to start or connect to game. Returning to menu...")
		return
	}
	defer conn.Close()

	game.GamePlay(conn, isHost, address)
}
