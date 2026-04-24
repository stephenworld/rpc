package game

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"stephen/banner"
	"stephen/utils"
	"strings"
)

type gameRound struct {
	user1  string
	user2  string
	result string
}

func GamePlay(conn net.Conn, isHost bool, GAMEID string) {
	fmt.Println(banner.Rpc)
	fmt.Printf("\n--- Game Connected (%s) ---\n", GAMEID)

	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(conn)

	playerRole := "Host"
	if !isHost {
		playerRole = "Guest"
	}
	fmt.Printf("You are playing as %s.\n\n", playerRole)

	for {
		fmt.Print("Enter rock, paper, scissors (or 'exit' to quit): ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if input == "exit" {
			fmt.Println("You left the game.")
			conn.Write([]byte("exit\n"))
			break
		}

		if !utils.CheckOptions(input) {
			fmt.Println("Invalid input. Please choose rock, paper, or scissors.")
			continue
		}

		// Send move to opponent
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Println("Connection lost.")
			break
		}

		fmt.Println("Waiting for opponent to choose...")

		// Wait for opponent's move
		opponentRaw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Opponent disconnected.")
			break
		}

		opponentMove := strings.TrimSpace(strings.ToLower(opponentRaw))
		if opponentMove == "exit" {
			fmt.Println("Opponent left the game.")
			break
		}

		fmt.Printf("\nYou chose: %s\n", input)
		fmt.Printf("Opponent chose: %s\n", opponentMove)

		result := utils.DetermineWinner(input, opponentMove)

		utils.ClearScreen()

		switch result {

		case "draw":
			fmt.Println("Result: It's a draw!")
		case "user1":
			fmt.Println("Result: You win this round!")
		case "user2":
			fmt.Println("Result: Opponent wins this round!")
		}
		fmt.Println()
	}
}
