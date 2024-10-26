package model_test

import (
	"reflect"
	"testing"

	"github.com/nonya123456/connect4/internal/model"
)

func TestBoard_String(t *testing.T) {
	tests := []struct {
		name string
		b    model.Board
		want string
	}{
		{
			name: "default",
			b:    model.Board{},
			want: "000000000000000000000000000000000000000000",
		},
		{
			name: "normal",
			b: model.Board{
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

func TestParseBoard(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Board
		wantErr bool
	}{
		{
			name:    "invalid length",
			args:    args{s: "0000"},
			wantErr: true,
		},
		{
			name:    "invalid byte",
			args:    args{s: "300500200000000000000000000000000002000001"},
			wantErr: true,
		},
		{
			name: "normal",
			args: args{s: "100000200000000000000000000000000002000001"},
			want: model.Board{
				{1, 0, 0, 0, 0, 0, 2},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.ParseBoard(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Winner(t *testing.T) {
	tests := []struct {
		name string
		b    model.Board
		want int
	}{
		{
			name: "0",
			b: model.Board{
				{1, 0, 0, 0, 0, 0, 2},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 0, 1},
			},
			want: 0,
		},
		{
			name: "horizontal",
			b: model.Board{
				{1, 1, 1, 1, 0, 0, 2},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 0, 1},
			},
			want: 1,
		},
		{
			name: "vertical",
			b: model.Board{
				{1, 1, 1, 2, 1, 2, 2},
				{0, 0, 0, 0, 0, 0, 2},
				{0, 0, 0, 0, 0, 0, 2},
				{0, 0, 0, 0, 0, 0, 2},
				{0, 0, 0, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 0, 1},
			},
			want: 2,
		},
		{
			name: "diagonal 1",
			b: model.Board{
				{1, 1, 1, 2, 1, 2, 2},
				{0, 0, 0, 0, 0, 2, 2},
				{0, 0, 0, 0, 2, 0, 0},
				{0, 0, 0, 2, 0, 0, 2},
				{0, 0, 0, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 0, 1},
			},
			want: 2,
		},
		{
			name: "diagonal 2",
			b: model.Board{
				{1, 1, 1, 2, 1, 2, 2},
				{0, 1, 0, 0, 0, 0, 2},
				{0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 2},
				{0, 0, 0, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 0, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Winner(); got != tt.want {
				t.Errorf("Board.Winner() = %v, want %v", got, tt.want)
			}
		})
	}
}
