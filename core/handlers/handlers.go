package handlers

import (
	"time"

	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/bot_enteties/buttons"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/bot_enteties/emojies"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/bot_enteties/text_message"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/core/auth"
	"github.com/bohdanlisovskyi/telegrambot/tbot"
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
		planet_name = "Planet"
	}
	auth.UserCreation(message)
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

func DataCenter(message *tbot.Message) {

	message.ReplyKeyboard("DataCenter", buttons.DataCenter, tbot.ResizeKeyboard)
}

func SkipTutorial(message *tbot.Message) {

	message.Reply(text_message.SkipToturial)

	message.ReplyKeyboard("Buildings", buttons.AllBuildings, tbot.ResizeKeyboard)
}

func AllBuildings(message *tbot.Message) {

	message.ReplyKeyboard("Buildings", buttons.AllBuildings, tbot.ResizeKeyboard)
}

func DataCenterStatus(message *tbot.Message) {
	//TODO show some status
	status := "some status here"

	message.Reply(status)
}

func DataCenterToturial(message *tbot.Message) {

	message.Reply(text_message.TutorialText)
}

func DataCenterSearchPlanet(message *tbot.Message) {

	//TODO code for search planet

	message.Reply("Some text as a result to search planet: Show list of neighbor!!!")

	message.ReplyKeyboard("Search Planet", buttons.AfterSearchButton, tbot.ResizeKeyboard)
}

func DataCenterSearchListOfPlanet(message *tbot.Message) {

	//TODO code for getting list of planet

	listOfPlanet := "some list of planet"

	message.Reply(listOfPlanet)

	message.ReplyKeyboard("Go to Cosmodrom", buttons.SearchListPlanet, tbot.ResizeKeyboard)
}

func CosmodromMenu(message *tbot.Message) {

	message.ReplyKeyboard("Cosmodrom", buttons.CosmodromMenu, tbot.ResizeKeyboard)
}
