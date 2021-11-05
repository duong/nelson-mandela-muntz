package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// For detecting when other users are speaking:

// g, err := session.State.Guild(guildID)
// if err != nil {
//   return err
// }

// g.VoiceStates // <- look at this

// Variables used for environment variables
var (
	token     string
	channelID string
	guildID   string
)

func init() {
	token = os.Getenv("TOKEN")
	channelID = os.Getenv("CHANNEL_ID")
	guildID = os.Getenv("GUILD_ID")
}

func main() {
	log.Println("hello bot")
	s, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal("error creating Discord session: ", err)
	}

	s.Identify.Intents = discordgo.IntentsGuildVoiceStates

	err = s.Open()
	if err != nil {
		log.Fatal("error opening connection: ", err)
	}

	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, false)
	log.Println(vc)

	if err != nil {
		log.Fatal("failed to join voice channel: ", err)
	}

	log.Println("joined voice")

}