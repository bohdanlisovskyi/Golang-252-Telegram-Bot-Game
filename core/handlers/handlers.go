package handlers

import (
	"time"

	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/bot_enteties/buttons"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/bot_enteties/emojies"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/bot_enteties/text_message"
	"github.com/yanzay/tbot"
)

func EchoHandler(message *tbot.Message) {
	message.Reply(message.Text())
}

func StartGame(message *tbot.Message) {
	message.Replyf("Hello, %s!", message.From)
	time.Sleep(1 * time.Second)
	message.Reply(text_message.MessageAboutGame)
	time.Sleep(500 * time.Millisecond)
	message.Reply(text_message.EnterPlanetName, tbot.WithMarkdown)
	time.Sleep(500 * time.Millisecond)

}

func PlanetName(message *tbot.Message) {

	//TODO check planet

	planet_name := message.Vars["planet_name"]

	if planet_name == "" {
		return
	}

	message.Replyf(text_message.PlanetCreated, planet_name)

	//TODO create all resource and buildings

	message.Reply(text_message.NextStep, tbot.WithMarkdown)

	message.ReplyKeyboard(emojies.DownwardsArrow+emojies.DownwardsArrow+emojies.DownwardsArrow+emojies.DownwardsArrow, buttons.Tutorial_Skip)
}

func GameTutorial(message *tbot.Message) {

	message.Reply(text_message.TutorialText)

	time.Sleep(3 * time.Second)

	message.ReplyKeyboard("Buildings", buttons.AllBuildings, tbot.ResizeKeyboard)
}
