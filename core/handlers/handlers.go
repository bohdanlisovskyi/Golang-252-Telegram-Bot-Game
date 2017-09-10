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
		{"Buildings"},
	}
	message.ReplyKeyboard("Here is  your planet, build the future!", buttons)
}

func ShowMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Research Centre", "Mines", "City Centre", "Cosmodrome", "Rating", "Planet history"},
	}
	message.ReplyKeyboard("Buildings", buttons)
}

func ShowResearchMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Tutorial", "Help", "Planet search", "Buildings"},
	}
	message.ReplyKeyboard("Research Centre", buttons)
}

func ShowMinesMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Metal mine", "Crystal mine", "Buildings"},
	}
	message.ReplyKeyboard("Mines", buttons)
}
func ShowMetalMinesMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Upgrade", "Employ people", "Buildings", "Status"},
	}
	message.ReplyKeyboard("Metal mine", buttons)
}
func ShowCrystalMinesMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Upgrade", "Employ people", "Buildings", "Status"},
	}
	message.ReplyKeyboard("Crystal mine", buttons)
}

func ShowCityCentreMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Upgrade", "Status", "Buildings"},
	}
	message.ReplyKeyboard("City Centre", buttons)
}

func ShowCosmodromeMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Fleet", "Dockyard", "Buildings"},
	}
	message.ReplyKeyboard("Cosmodrome", buttons)
}

func ShowFleetMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Set new fleet", "Edit existing group", "Status", "Buildings"},
	}
	message.ReplyKeyboard("Fleet", buttons)
}

func ShowDockyardMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Build new ships", "Status", "Buildings"},
	}
	message.ReplyKeyboard("Dockyard", buttons)
}

func ShowRatingMenu(message *tbot.Message) {

	buttons := [][]string{
		{"Resourses used", "Population", "Buildings"},
	}
	message.ReplyKeyboard("Rating", buttons)
}
