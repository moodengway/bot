package model

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/connect4/internal/util"
)

type Match struct {
	ID        uint    `gorm:"column:id;primaryKey;autoIncrement"`
	MessageID string  `gorm:"column:message_id"`
	Host      string  `gorm:"column:host"`
	Guest     *string `gorm:"column:guest"`
}

const (
	Aqua   int = 1752220
	Red    int = 15548997
	Yellow int = 16776960
)

func (m Match) MessageEmbed() discordgo.MessageEmbed {
	title := fmt.Sprintf("Match#%d", m.ID)

	host := util.Mention(m.Host)
	guest := "N/A"
	color := Aqua

	if m.Guest != nil {
		guest = util.Mention(*m.Guest)
		color = Red
	}

	description := fmt.Sprintf("Red: %s\nYellow: %s", host, guest)

	return discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       color,
	}
}
