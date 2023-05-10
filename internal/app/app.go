package app

import (
	"VK_BOT_RAW/internal/usecase/longpoll"
	"VK_BOT_RAW/internal/usecase/messages"
	"log"
	"os"
)

func Run() {
	log.Print("|INFO| Bot start!")
	token := os.Getenv("VK_TOKEN")
	group := os.Getenv("VK_GROUP")
	v := os.Getenv("VK_API_VERSION")
	wait := os.Getenv("VK_WAIT_TIME")

	messageHandler := messages.NewMessageHandler(token, v)

	longpoll := longpoll.NewLongPoll(token, v, group, messageHandler, wait)
	err := longpoll.Run()
	if err != nil {
		log.Fatal("Fatal error:", err)
	}
}
