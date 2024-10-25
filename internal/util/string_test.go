package util

import (
	"testing"
)

func TestMention(t *testing.T) {
	tests := []struct {
		name   string
		userID string
		want   string
	}{
		{
			name:   "normal",
			userID: "1111",
			want:   "<@1111>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mention(tt.userID); got != tt.want {
				t.Errorf("Mention() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoardString(t *testing.T) {
	tests := []struct {
		name  string
		board [6][7]int
		want  string
	}{
		{
			name:  "default",
			board: [6][7]int{},
			want:  "000000000000000000000000000000000000000000",
		},
		{
			name: "normal",
			board: [6][7]int{
				{1, 0, 0, 0, 0, 0, 2},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 0, 1},
			},
			want: "100000200000000000000000000000000002000001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoardString(tt.board); got != tt.want {
				t.Errorf("BoardString() = %v, want %v", got, tt.want)
			}
		})
	}
}
