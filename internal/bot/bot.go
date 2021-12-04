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

// Variables used for environment variables
var (
	token     string
	channelID string
	guildID   string
)

var session *discordgo.Session
var voiceConnection *discordgo.VoiceConnection

func Init() {
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

	start()
}

func start() {
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

	guild, err := session.State.Guild(guildID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", vs)

	for _, key := range guild.VoiceStates {
		log.Printf("%+v\n", key)
	}
	voiceConnection.AddHandler(voiceSpeakingUpdateHandler)
}
