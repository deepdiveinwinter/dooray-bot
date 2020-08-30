package main

import (
	"github.com/deepdiveinwinter/dooray-bot/clients"
	"os"
)

var (
	DoorayHookUrlKey  = "DOORAY_HOOK_URL"
	DoorayBotName     = "HelloWorldBot"
	DoorayBotImageUrl = "https://static.dooray.com/static_images/dooray-bot.png"
)

func main() {

	doorayHookUrl := os.Getenv(DoorayHookUrlKey)

	hookClient := clients.NewHookClient(doorayHookUrl, DoorayBotName, DoorayBotImageUrl)

	doorayMessage := clients.DoorayMessage{
		Text: "Hello World!",
		Attachments: []clients.DoorayAttachment{
			{
				Title:     "Attachment",
				TitleLink: "https://docs.toast.com/ko/Dooray/Messenger/ko/incoming-hook-guide/",
				Text:      "message",
				Color:     "blue",
			},
		},
	}

	err := hookClient.SendMessage(doorayMessage)
	if err != nil {
		panic(err)
	}
}
