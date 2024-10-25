package model_test

import (
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/connect4/internal/model"
)

func TestMatch_MessageEmbed(t *testing.T) {
	type fields struct {
		ID    uint
		Host  string
		Guest *string
	}
	tests := []struct {
		name   string
		fields fields
		want   discordgo.MessageEmbed
	}{
		{
			name: "waiting",
			fields: fields{
				ID:    1111,
				Host:  "2222",
				Guest: nil,
			},
			want: discordgo.MessageEmbed{
				Title:       "Match#1111",
				Description: "Red: <@2222>\nYellow: N/A",
				Color:       model.WaitingColor,
			},
		},
		{
			name: "playing",
			fields: fields{
				ID:    1111,
				Host:  "2222",
				Guest: toStringPointer("3333"),
			},
			want: discordgo.MessageEmbed{
				Title:       "Match#1111",
				Description: "Red: <@2222>\nYellow: <@3333>",
				Color:       model.PlayingColor,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model.Match{
				ID:    tt.fields.ID,
				Host:  tt.fields.Host,
				Guest: tt.fields.Guest,
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
