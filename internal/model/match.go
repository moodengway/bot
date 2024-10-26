package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/connect4/internal/util"
	"gorm.io/gorm"
)

type Match struct {
	ID          uint       `gorm:"column:id;primaryKey;autoIncrement"`
	MessageID   string     `gorm:"column:message_id"`
	Host        string     `gorm:"column:host"`
	Guest       *string    `gorm:"column:guest"`
	Board       Board      `gorm:"-"`
	BoardString string     `gorm:"column:board_string"`
	RoundNumber int        `gorm:"column:round_number"`
	EndedAt     *time.Time `gorm:"column:ended_at"`
}

const (
	Aqua   int = 1752220
	Red    int = 15548997
	Yellow int = 16776960
)

func (m *Match) MessageEmbed() discordgo.MessageEmbed {
	title := fmt.Sprintf("Match#%d", m.ID)

	host := util.Mention(m.Host)
	guest := "N/A"
	color := Aqua

	if m.Guest != nil {
		guest = util.Mention(*m.Guest)
		color = Red
		if m.RoundNumber%2 == 0 {
			color = Yellow
		}
	}

	board, _ := m.boardEmbedString()
	description := fmt.Sprintf("🔴 %s\n\n🟡 %s\n\n```%s```", host, guest, board)

	return discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       color,
	}
}

func (m *Match) boardEmbedString() (string, error) {
	if len(m.BoardString) != 42 {
		return "", errors.New("invalid board string length")
	}

	result := ""

	mapper := make(map[byte]rune)
	mapper['0'] = '⚪'
	mapper['1'] = '🔴'
	mapper['2'] = '🟡'

	for i := 5; i >= 0; i-- {
		row := ""
		for j := 0; j < 7; j++ {
			b := m.BoardString[i*7+j]
			emoji, ok := mapper[b]
			if !ok {
				return "", errors.New("invalid byte in board string")
			}

			row += fmt.Sprintf("%c ", emoji)
		}

		row = row[0 : len(row)-1]
		result += row + "\n"
	}

	result += "1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣"
	return result, nil
}

func (m *Match) BeforeSave(tx *gorm.DB) (err error) {
	m.BoardString = m.Board.String()
	return nil
}

func (m *Match) AfterFind(tx *gorm.DB) (err error) {
	board, err := ParseBoard(m.BoardString)
	if err != nil {
		return err
	}

	m.Board = board
	return nil
}
