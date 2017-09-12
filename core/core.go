package core

import (
	"os"

	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/loger"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/router"
	"github.com/bohdanlisovskyi/telegrambot/tbot"
)

func Run() {
	token := os.Getenv("TELEGRAM_TOKEN")

	if token == "" {
		loger.Log.Fatal("Token Empty")
	}

	_botServerUp(token)
}

func _botServerUp(token string) {
	bot, err := tbot.NewServer(token)

	if err != nil {
		loger.Log.Fatal(err)
	}

	loger.Log.Fatal(router.NewRouter(bot).ListenAndServe())
}
