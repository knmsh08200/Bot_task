package bot

import(
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "github.com/knmsh08200/Bot_task/internal/broker"
)

type BotService struct {
    Service broker.TicketService
}

type TelegramAPI struct {
    Bot *tgbotapi.BotAPI
}

func NotifyAdmins(ctx context.Context, bot *tgbotapi.BotAPI, redisClient *redis.Client) {
    for {
        // Периодически проверяем наличие новых тикетов для уведомления
        time.Sleep(5 * time.Second) // Интервал проверки

        // Извлекаем тикеты для каждого отдела
        departments := []string{"Support", "IT", "Billing"}
        for _, department := range departments {
            adminListKey := department + "_admin_notifications"

            // Извлекаем тикет
            ticketID, err := redisClient.LPop(ctx, adminListKey).Result()
            if err == redis.Nil { // trash
                continue // Нет новых тикетов
            } else if err != nil {
                continue // Обрабатываем ошибку
            }

            // Отправляем уведомление админу
            message := tgbotapi.NewMessage(/*здесь ID админа */, "Новый тикет:\nID: "+ticketID)
            _, err = bot.Send(message)
            if err != nil {
                log.Printf("Тикет не доставлен администратору") // Обрабатываем ошибку отправки
            }
			
        }
    }
}