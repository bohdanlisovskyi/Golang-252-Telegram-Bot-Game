package auth

import (
	"testing"

	"fmt"

	"github.com/bohdanlisovskyi/telegrambot/tbot"
)

//var mockadapter = adapter.Message{
//	1,
//	"",
//	"",
//	nil,
//	"greckas",
//	2,
//	true,
//	true,
//	nil,
//	true,
//	true,
//}

func mockMessage() *tbot.Message {

	varta := make(map[string]string)
	varta["planet_name"] = "Earffth"
	m := &tbot.Message{
		nil,
		varta,
	}
	return m
}

func TestUserCreation(t *testing.T) {

	str, err := UserCreation(mockMessage())
	fmt.Println(str, err)
}
