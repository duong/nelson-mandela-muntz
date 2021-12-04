/*
Package bot implements the main bot logic and loops.

Listens for sound coming from other discord users and will respond with Nelson Muntz's signature "hah-hah"
*/
package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var session *discordgo.Session
var voiceConnection *discordgo.VoiceConnection

func Start(token string, channelID string, guildID string) {
	log.Println("initialising bot...")

	var err error
	session, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal("error creating Discord session: ", err)
	}

	session.Identify.Intents = discordgo.IntentsGuildVoiceStates

	err = session.Open()
	if err != nil {
		log.Fatal("error opening connection: ", err)
	}

	voiceConnection, err = session.ChannelVoiceJoin(guildID, channelID, false, false)
	log.Println("created voice connection")

	if err != nil {
		log.Fatal("failed to join voice channel: ", err)
	}

	// For detecting when other users are speaking:
	log.Println("state enabled", session.StateEnabled)
	voiceConnection.AddHandler(voiceSpeakingUpdateHandler)

	// Keep the program running until it has signal interrupt
	defer session.Close()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Println("Gracefully shutdowning")
}

func voiceSpeakingUpdateHandler(vc *discordgo.VoiceConnection, vs *discordgo.VoiceSpeakingUpdate) {
	log.Println("voice speaking handler")

	log.Printf("%+v\n", vs)

	voiceConnection.AddHandler(voiceSpeakingUpdateHandler)
}
