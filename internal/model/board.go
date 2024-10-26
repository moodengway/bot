package model

import "strconv"

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
