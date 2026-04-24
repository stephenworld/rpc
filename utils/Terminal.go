package utils

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func GetTerminalWidth() int {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 80 // Fallback
	}

	trimSpace := strings.TrimSpace(string(out))

	width, err := strconv.Atoi(trimSpace)
	if err != nil {
		return 80 //Falback
	}

	return width
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Println("Unable to clear screen:", err)
	}
}
