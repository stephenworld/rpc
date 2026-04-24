package utils

import "strings"

func CheckOptions(input string) bool {
	switch strings.ToLower(input) {
	case "rock", "paper", "scissors":
		return true
	}
	return false
}

func DetermineWinner(p1, p2 string) string {
	p1 = strings.ToLower(p1)
	p2 = strings.ToLower(p2)

	if p1 == p2 {
		return "draw"
	}
	switch p1 {
	case "rock":
		if p2 == "scissors" {
			return "user1"
		}
	case "paper":
		if p2 == "rock" {
			return "user1"
		}
	case "scissors":
		if p2 == "paper" {
			return "user1"
		}
	}
	return "user2"
}
