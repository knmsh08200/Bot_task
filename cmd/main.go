package main

import (
	"context"
	"fmt"
	"log"
	"os"

	tgbot "github.com/knmsh08200/Bot_task/internal/bot"
	"github.com/knmsh08200/Bot_task/internal/db"
	"github.com/knmsh08200/Bot_task/service"
)

func main() {
	ctx := context.Background()

	ctxWithCancel, cancelFunc := context.WithCancel(ctx)

	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunc()
	}()

	botToken := os.Getenv("TELEGRAM_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGTAM_TOKEN must be set")
	}

	redisClient := db.ConnectRedis(ctx)
	defer redisClient.Close()

	dbConn := db.ConnectPostgres(ctx)
	defer dbConn.Close()

	bot, err := tgbot.NewTelegramBot("YOUR_TOKEN_HERE", true)
	if err != nil {
		log.Fatal(err)
	}

	updates, err := tgbot.StartBot(bot, 60)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		service.BotWork(bot, update /*interface*/)
	}
}
