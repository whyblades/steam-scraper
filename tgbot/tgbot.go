package tgbot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
)

var chatID int64

func InitBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAMBOT_APITOKEN"))
	if err != nil {
		fmt.Println("Failed to init telegram bot: wrong bot api token")
	}
	chatID, err = strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	if err != nil {
		fmt.Println("Failed to init telegram bot: wrong chat id")
	}
	fmt.Println("initialized tg bot")
	return bot
}

func NotifyNewSkin(b *tgbotapi.BotAPI, str string) {
	msg := tgbotapi.NewMessage(chatID, str)
	_, err := b.Send(msg)
	if err != nil {
		fmt.Println("telegram bot error: failed to send message")
	}
}
