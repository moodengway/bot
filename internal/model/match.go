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
	Aqua   int = 1752220
	Red    int = 15548997
	Yellow int = 16776960
)

func (m Match) MessageEmbed() discordgo.MessageEmbed {
	title := fmt.Sprintf("Match#%d", m.ID)

	host := createMention(m.Host)
	guest := "N/A"
	color := Aqua

	if m.Guest != nil {
		guest = createMention(*m.Guest)
		color = Red
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
