package model_test

import (
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/connect4/internal/model"
)

func TestMatch_MessageEmbed(t *testing.T) {
	type fields struct {
		ID          uint
		MessageID   string
		Host        string
		Guest       *string
		BoardString string
		RoundNumber int
	}
	tests := []struct {
		name   string
		fields fields
		want   discordgo.MessageEmbed
	}{
		{
			name: "waiting",
			fields: fields{
				ID:          1111,
				Host:        "2222",
				Guest:       nil,
				BoardString: "000000000000000000000000000000000000000000",
			},
			want: discordgo.MessageEmbed{
				Title:       "Match#1111",
				Description: "🔴 <@2222>\n\n🟡 N/A\n\n```⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣```",
				Color:       model.Aqua,
			},
		},
		{
			name: "playing",
			fields: fields{
				ID:          1111,
				Host:        "2222",
				Guest:       toStringPointer("3333"),
				BoardString: "100000200000000000000000000000000002000001",
			},
			want: discordgo.MessageEmbed{
				Title:       "Match#1111",
				Description: "🔴 <@2222>\n\n🟡 <@3333>\n\n```🟡 ⚪ ⚪ ⚪ ⚪ ⚪ 🔴\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n🔴 ⚪ ⚪ ⚪ ⚪ ⚪ 🟡\n1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣```",
				Color:       model.Red,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model.Match{
				ID:          tt.fields.ID,
				MessageID:   tt.fields.MessageID,
				Host:        tt.fields.Host,
				Guest:       tt.fields.Guest,
				BoardString: tt.fields.BoardString,
				RoundNumber: tt.fields.RoundNumber,
			}
			if got := m.MessageEmbed(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match.MessageEmbed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func toStringPointer(s string) *string {
	return &s
}

func TestMatch_Board(t *testing.T) {
	type fields struct {
		ID          uint
		MessageID   string
		Host        string
		Guest       *string
		BoardString string
		RoundNumber int
	}
	tests := []struct {
		name    string
		fields  fields
		want    [6][7]int
		wantErr bool
	}{
		{
			name:    "invalid length",
			fields:  fields{BoardString: "0000"},
			wantErr: true,
		},
		{
			name:    "invalid byte",
			fields:  fields{BoardString: "300500200000000000000000000000000002000001"},
			wantErr: true,
		},
		{
			name:   "normal",
			fields: fields{BoardString: "100000200000000000000000000000000002000001"},
			want: [6][7]int{
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
			m := model.Match{
				ID:          tt.fields.ID,
				MessageID:   tt.fields.MessageID,
				Host:        tt.fields.Host,
				Guest:       tt.fields.Guest,
				BoardString: tt.fields.BoardString,
				RoundNumber: tt.fields.RoundNumber,
			}
			got, err := m.Board()
			if (err != nil) != tt.wantErr {
				t.Errorf("Match.Board() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match.Board() = %v, want %v", got, tt.want)
			}
		})
	}
}
