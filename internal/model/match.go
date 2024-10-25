package model

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Match struct {
	ID    uint    `gorm:"column:id;primaryKey;autoIncrement"`
	Host  string  `gorm:"column:host"`
	Guest *string `gorm:"column:guest"`
}

const (
	WaitingColor int = 16776960
	PlayingColor int = 1752220
)

func (m Match) MessageEmbed() discordgo.MessageEmbed {
	title := fmt.Sprintf("Match#%d", m.ID)

	host := createMention(m.Host)
	guest := "N/A"
	color := WaitingColor

	if m.Guest != nil {
		guest = createMention(*m.Guest)
		color = PlayingColor
	}

	description := fmt.Sprintf("Red: %s\nYellow: %s", host, guest)

	return discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       color,
	}
}

func createMention(userID string) string {
	return fmt.Sprintf("<@%s>", userID)
}
