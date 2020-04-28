package main

import (
	"log"

	"github.com/amroessam/SlackBullet/handlers"
)

func main() {
	log.Println("Starting SlackBullet...")

	//Read yaml configurations
	var config handlers.Config = handlers.ReadConfig()

	//Loop through all configured pushbullet accounts
	for _, el := range config.PushbulletAccounts {
		handlers.WSHandler(el)
	}
}
