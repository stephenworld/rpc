package game

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"stephen/utils"
	"strings"
)

func CreateGame() (net.Conn, string) {
	fmt.Print("Host a game! Enter a listening Port (or press Enter for 8080): ")
	scanner := bufio.NewScanner(os.Stdin)
	port := "8080"
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input != "" {
			port = input
		}
	}
	utils.ClearScreen()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting host:", err)
		return nil, ""
	}

	fmt.Printf("\nHosting on port %s. Waiting for your friend to connect...\n", port)
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return nil, ""
	}

	fmt.Println("Success! A friend has connected.")
	return conn, port
}
