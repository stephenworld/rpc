package game

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"stephen/utils"
	"strings"
)

func JoinGame() (net.Conn, string) {
	fmt.Print("Enter friend's IP Address (e.g. 192.168.1.5 or localhost:8080): ")
	scanner := bufio.NewScanner(os.Stdin)
	address := ""
	if scanner.Scan() {
		address = strings.TrimSpace(scanner.Text())
	}
	utils.ClearScreen()

	if !strings.Contains(address, ":") {
		// default to 8080
		address = address + ":8080"
	}

	fmt.Printf("Connecting to %s...\n", address)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to friend:", err)
		return nil, ""
	}

	fmt.Println("Success! Connected to friend.")
	return conn, address
}
