package tgbot

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

func parse(msg *tgbotapi.Message) (string, int, error) {
	text := strings.TrimSpace(msg.Text)
	parts := strings.Split(text, "\n")

	if len(parts) < 2 {
		return "", 0, errors.New("Неправильный формат ввода")
	}

	reminderText := parts[0]
	reminderInterval, _ := strconv.Atoi(parts[1])

	return reminderText, reminderInterval, nil
}
