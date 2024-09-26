package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tgbot "github.com/knmsh08200/Bot_task/internal/bot"
	"github.com/knmsh08200/Bot_task/internal/broker"
	"github.com/knmsh08200/Bot_task/internal/db"
	"github.com/knmsh08200/Bot_task/internal/handler"
	"github.com/knmsh08200/Bot_task/internal/router"
)

func main() {
	ctx := context.Background()

	ctxWithCancel, cancelFunc := context.WithCancel(ctx)

	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunc()
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	botToken := os.Getenv("TELEGRAM_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGTAM_TOKEN must be set")
	}

	redisClient := db.ConnectRedis(ctx)
	defer redisClient.Close()

	dbConn := db.ConnectPostgres(ctxWithCancel)
	defer dbConn.Close()

	bot, err := tgbot.NewTelegramBot("YOUR_TOKEN_HERE", true)
	if err != nil {
		log.Fatal(err)
	}

	updates, err := tgbot.StartBot(bot, 60)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	rdb := broker.NewRepRedis(redisClient)
	TicketRep := tgbot.NewRep(rdb)
	ticketHandler := handler.NewTicketHandler(TicketRep)

	go func() {
		for update := range updates {
			router.BotWork(ctx, bot.Bot, update, ticketHandler)
		}
		close(done)
	}()

	// реализация graceful shutdown
	select {
	case <-sigs:
		fmt.Println("Received shutdown signal")

		cancelFunc()

		select {
		case <-done:
			fmt.Println("Bot gracefully stopped")
		case <-time.After(5 * time.Second): // Таймаут на graceful shutdown
			fmt.Println("Bot shutdown timed out")
		}
	case <-done:
		fmt.Println("Bot stopped")
	}

}
