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
