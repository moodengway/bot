package util

import (
	"fmt"
)

func Mention(userID string) string {
	return fmt.Sprintf("<@%s>", userID)
}

func ToPtr(s string) *string {
	return &s
}
