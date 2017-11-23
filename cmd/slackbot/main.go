package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")
			case *slack.MessageEvent:
				log.Print("Message: %v\n", ev)
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", ev.Channel))
			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1
			}
		}
	}
}

func main() {
	token := os.Getenv("BOTAPITOKEN")
	if token == "" {
		log.Fatal("missing API token")
	}
	api := slack.New()
	os.Exit(run(api))
}
