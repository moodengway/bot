package model

import (
	"errors"
	"strconv"
)

type Board [6][7]int

func (b Board) String() string {
	s := ""
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			s += strconv.Itoa(b[i][j])
		}
	}

	return s
}

func (b Board) Winner() int {
	// Check horizontal
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			if b[i][j] == 0 {
				continue
			}

			if b[i][j] == b[i][j+1] && b[i][j] == b[i][j+2] && b[i][j] == b[i][j+3] {
				return b[i][j]
			}
		}
	}

	// Check vertical
	for j := 0; j < 7; j++ {
		for i := 0; i < 3; i++ {
			if b[i][j] == 0 {
				continue
			}

			if b[i][j] == b[i+1][j] && b[i][j] == b[i+2][j] && b[i][j] == b[i+3][j] {
				return b[i][j]
			}
		}
	}

	// Check right diagonal
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if b[i][j] == 0 {
				continue
			}

			if b[i][j] == b[i+1][j+1] && b[i][j] == b[i+2][j+2] && b[i][j] == b[i+3][j+3] {
				return b[i][j]
			}
		}
	}

	// Check left diagonal
	for i := 0; i < 3; i++ {
		for j := 3; j < 7; j++ {
			if b[i][j] == 0 {
				continue
			}

			if b[i][j] == b[i+1][j-1] && b[i][j] == b[i+2][j-2] && b[i][j] == b[i+3][j-3] {
				return b[i][j]
			}
		}
	}

	return 0
}

func ParseBoard(s string) (Board, error) {
	if len(s) != 42 {
		return Board{}, errors.New("invalid board string length")
	}

	mapper := make(map[byte]int)
	mapper['0'] = 0
	mapper['1'] = 1
	mapper['2'] = 2

	var board Board
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			b := s[i*7+j]
			num, ok := mapper[b]
			if !ok {
				return Board{}, errors.New("invalid byte in board string")
			}

			board[i][j] = num
		}
	}

	return board, nil
}
