package model_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/connect4/internal/model"
	"github.com/nonya123456/connect4/internal/util"
)

func TestMatch_MessageEmbed(t *testing.T) {
	type fields struct {
		ID          uint
		MessageID   string
		Host        string
		Guest       *string
		BoardString string
		RoundNumber int
		EndedAt     *time.Time
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
				RoundNumber: 1,
			},
			want: discordgo.MessageEmbed{
				Title:       "Match#1111",
				Description: "🔴 <@2222>\n\n🟡 N/A\n\n```⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣```",
				Color:       model.Aqua,
			},
		},
		{
			name: "playing red turn",
			fields: fields{
				ID:          1111,
				Host:        "2222",
				Guest:       util.ToPtr("3333"),
				BoardString: "100000200000000000000000000000000002000001",
				RoundNumber: 1,
			},
			want: discordgo.MessageEmbed{
				Title:       "Match#1111",
				Description: "🔴 <@2222>\n\n🟡 <@3333>\n\n```🟡 ⚪ ⚪ ⚪ ⚪ ⚪ 🔴\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n🔴 ⚪ ⚪ ⚪ ⚪ ⚪ 🟡\n1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣```",
				Color:       model.Red,
			},
		},
		{
			name: "playing yellow turn",
			fields: fields{
				ID:          1111,
				Host:        "2222",
				Guest:       util.ToPtr("3333"),
				BoardString: "100000200000000000000000000000000002000001",
				RoundNumber: 2,
			},
			want: discordgo.MessageEmbed{
				Title:       "Match#1111",
				Description: "🔴 <@2222>\n\n🟡 <@3333>\n\n```🟡 ⚪ ⚪ ⚪ ⚪ ⚪ 🔴\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n🔴 ⚪ ⚪ ⚪ ⚪ ⚪ 🟡\n1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣```",
				Color:       model.Yellow,
			},
		},
		{
			name: "ended",
			fields: fields{
				ID:          1111,
				Host:        "2222",
				Guest:       util.ToPtr("3333"),
				BoardString: "100000200000000000000000000000000002000001",
				RoundNumber: 1,
				EndedAt:     &time.Time{},
			},
			want: discordgo.MessageEmbed{
				Title:       "Match#1111",
				Description: "🔴 <@2222>\n\n🟡 <@3333>\n\n```🟡 ⚪ ⚪ ⚪ ⚪ ⚪ 🔴\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n⚪ ⚪ ⚪ ⚪ ⚪ ⚪ ⚪\n🔴 ⚪ ⚪ ⚪ ⚪ ⚪ 🟡\n1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣```",
				Color:       model.Gray,
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
				EndedAt:     tt.fields.EndedAt,
			}
			if got := m.MessageEmbed(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match.MessageEmbed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatch_IsEnded(t *testing.T) {
	type fields struct {
		ID          uint
		MessageID   string
		Host        string
		Guest       *string
		Board       model.Board
		BoardString string
		RoundNumber int
		EndedAt     *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "not ended",
			fields: fields{
				Board: model.Board{
					{1, 0, 0, 0, 0, 0, 2},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{2, 0, 0, 0, 0, 0, 1},
				},
				RoundNumber: 5,
			},
			want: false,
		},
		{
			name: "1",
			fields: fields{
				Board: model.Board{
					{1, 1, 1, 1, 0, 0, 2},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{2, 0, 0, 0, 0, 0, 1},
				},
				RoundNumber: 8,
			},
			want: true,
		},
		{
			name: "2",
			fields: fields{
				Board: model.Board{
					{1, 0, 0, 2, 2, 2, 2},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{2, 0, 0, 0, 0, 0, 1},
				},
				RoundNumber: 8,
			},
			want: true,
		},
		{
			name: "draw",
			fields: fields{
				Board: model.Board{
					{1, 0, 0, 0, 0, 0, 2},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{2, 0, 0, 0, 0, 0, 1},
				},
				RoundNumber: 45,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &model.Match{
				ID:          tt.fields.ID,
				MessageID:   tt.fields.MessageID,
				Host:        tt.fields.Host,
				Guest:       tt.fields.Guest,
				Board:       tt.fields.Board,
				BoardString: tt.fields.BoardString,
				RoundNumber: tt.fields.RoundNumber,
				EndedAt:     tt.fields.EndedAt,
			}
			if got := m.IsEnded(); got != tt.want {
				t.Errorf("Match.IsEnded() = %v, want %v", got, tt.want)
			}
		})
	}
}
