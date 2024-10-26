package model

import "testing"

func TestBoard_String(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want string
	}{
		{
			name: "default",
			b:    Board{},
			want: "000000000000000000000000000000000000000000",
		},
		{
			name: "normal",
			b: Board{
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
			if got := tt.b.String(); got != tt.want {
				t.Errorf("Board.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
