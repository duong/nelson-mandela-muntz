/*
Package main nelson-mandela-muntz
*/
package main

import (
	"log"
	"os"

	"github.com/duong/nelson-mandela-muntz/internal/bot"
)

// Variables used for environment variables
var (
	token     string
	channelID string
	guildID   string
)

func main() {
	token = os.Getenv("TOKEN")
	channelID = os.Getenv("CHANNEL_ID")
	guildID = os.Getenv("GUILD_ID")

	if token == "" {
		log.Fatal("no TOKEN provided")
	} else if channelID == "" {
		log.Fatal("no CHANNEL_ID provided")
	} else if guildID == "" {
		log.Fatal("no GUILD_ID provided")
	}

	bot.Start(token, channelID, guildID)
}
