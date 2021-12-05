/*
Package bot implements the main bot logic and loops.

Listens for sound coming from other discord users and will respond with Nelson Muntz's signature "hah-hah"
*/
package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

var session *discordgo.Session
var voiceConnection *discordgo.VoiceConnection
var guildID string

var audioFile string = "audio/nelson-ha-ha.m4a"

func Start(token string, channelID string, gid string) {
	log.Println("initialising bot...")
	guildID = gid

	var err error
	session, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal("error creating Discord session: ", err)
	}

	session.Identify.Intents = discordgo.IntentsGuildVoiceStates

	// session.AddHandler(voiceStateUpdateHandler)

	err = session.Open()
	if err != nil {
		log.Fatal("error opening connection: ", err)
	}

	voiceConnection, err = session.ChannelVoiceJoin(guildID, channelID, false, false)
	voiceConnection.LogLevel = discordgo.LogDebug
	log.Println("created voice connection")

	if err != nil {
		log.Fatal("failed to join voice channel: ", err)
	}

	// For detecting when other users are speaking:
	log.Println("state enabled", session.StateEnabled)

	log.Println("PlayAudioFile:", audioFile)

	voiceConnection.AddHandler(func(vc *discordgo.VoiceConnection, vs *discordgo.VoiceSpeakingUpdate) {
		log.Printf("user %s with ssrc %d is speaking: %t", vs.UserID, vs.SSRC, vs.Speaking)

		vc.Speaking(true)
		dgvoice.PlayAudioFile(vc, audioFile, make(chan bool))
		vc.Speaking(false)

	})

	// Keep the program running until it has signal interrupt
	defer session.Close()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Println("Gracefully shutting down")
}
