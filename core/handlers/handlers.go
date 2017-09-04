package handlers

import (
	"time"

	"github.com/yanzay/tbot"
)

func EchoHandler(message *tbot.Message) {
	message.Reply(message.Text())
}

func StartGame(message *tbot.Message) {
	message.Replyf("Hello, %s!", message.From)
	time.Sleep(1 * time.Second)
	message.Reply("Let's start game")
	time.Sleep(500 * time.Millisecond)

	buttons := [][]string{
		{"Menu"},
	}
	message.ReplyKeyboard("Buttons example", buttons)
}

func ShowMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Units", "Building"},
	}
	message.ReplyKeyboard("Menu", buttons)
}
