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
