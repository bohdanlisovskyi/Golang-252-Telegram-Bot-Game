package core

import (
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/loger"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/vendor/gopkg.in/telegram-bot-api.v4"
)

func Run() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		loger.Log.Panic(err)
	}

	bot.Debug = true
}
