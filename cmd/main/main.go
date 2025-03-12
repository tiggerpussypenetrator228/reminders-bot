package main

import (
	"reminder-bot/internal/database"
	"reminder-bot/internal/tgbot"
)

func main() {
	cfg := database.DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "root",
		Password: "nyanyan",
		DBName:   "reminders",
	}

	db, err := database.New(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tgbot.LaunchTheBot(db)
}
