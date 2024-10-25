package util

import (
	"fmt"
	"strconv"
)

func Mention(userID string) string {
	return fmt.Sprintf("<@%s>", userID)
}

func BoardString(board [6][7]int) string {
	result := ""

	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			result += strconv.Itoa(board[i][j])
		}
	}

	return result
}
