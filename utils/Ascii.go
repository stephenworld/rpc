package utils

import (
	_ "embed"
	"strings"
)

//go:embed standard.txt
var standardArt string

func GetAsciiArt(text string) string {
	standardArt = strings.ReplaceAll(standardArt, "\r\n", "\n")
	lines := strings.Split(standardArt, "\n")
	var result string
	for i := 0; i < 8; i++ {
		for _, char := range text {
			if char >= 32 && char <= 126 {
				idx := 1 + int(char-32)*9 + i
				if idx < len(lines) {
					result += lines[idx]
				}
			}
		}
		result += "\n"
	}
	return result
}

func GetRpcArt(move1, move2, rpcStr, rpcMirrorStr string) string {
	rpcStr = strings.Trim(strings.ReplaceAll(rpcStr, "\r\n", "\n"), "\n")
	rpcMirrorStr = strings.Trim(strings.ReplaceAll(rpcMirrorStr, "\r\n", "\n"), "\n")

	blocks := strings.Split(rpcStr, "\n\n")
	mirrorBlocks := strings.Split(rpcMirrorStr, "\n\n")

	optIdx := map[string]int{
		"rock":     0,
		"paper":    1,
		"scissors": 2,
	}

	idx1, ok1 := optIdx[move1]
	if !ok1 {
		idx1 = 0
	}
	idx2, ok2 := optIdx[move2]
	if !ok2 {
		idx2 = 0
	}

	lines1 := strings.Split(strings.Trim(blocks[idx1], "\n"), "\n")
	lines2 := strings.Split(strings.Trim(mirrorBlocks[idx2], "\n"), "\n")

	var res string
	for i := 0; i < len(lines1) || i < len(lines2); i++ {
		l1 := ""
		if i < len(lines1) {
			l1 = lines1[i]
		}
		l2 := ""
		if i < len(lines2) {
			l2 = lines2[i]
		}

		pad := 25 - len([]rune(l1))
		if pad < 0 {
			pad = 0
		}

		midStr := "      "
		if i == 2 {
			midStr = "  VS  "
		}

		res += l1 + strings.Repeat(" ", pad) + midStr + l2 + "\n"
	}

	return res
}
