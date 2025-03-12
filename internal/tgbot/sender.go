package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"reminder-bot/internal/database"
	"time"
)

func SendMsg(db *database.Database, bot *tgbotapi.BotAPI) {
	for {
		reminders, err := db.GetReminders(true)
		if err != nil {
			log.Println(err)
			continue
		}

		now := time.Now().UTC()
		for _, reminder := range reminders {
			lastChecked := reminder.LastChecked.UTC()
			nextCheckTime := lastChecked.Add(reminder.Interval)

			if now.After(nextCheckTime) || now.Equal(nextCheckTime) {
				userID, err := db.GetChatID(int64(reminder.UserID))
				if err != nil {
					log.Println(err)
					continue
				}

				msg := tgbotapi.NewMessage(userID, reminder.Content)
				bot.Send(msg)

				log.Printf("sending reminder %d to %d", reminder.UserID, userID)

				err = db.UpdateLastCheched(int64(reminder.ID))
				if err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}
}
