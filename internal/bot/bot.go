/*
Package bot implements the main bot logic and loops.

Listens for sound coming from other discord users and will respond with Nelson Muntz's signature "hah-hah"
*/
package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/dgvoice"
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

	// session.AddHandler(voiceStateUpdateHandler)

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

	listenAndPlay(voiceConnection)

	// Keep the program running until it has signal interrupt
	defer session.Close()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Println("Gracefully shutting down")
}

// Takes inbound audio and sends it right back out.
func listenAndPlay(v *discordgo.VoiceConnection) {

	log.Println("listen and play started")

	recv := make(chan *discordgo.Packet, 2)
	go dgvoice.ReceivePCM(v, recv)

	send := make(chan []int16, 2)
	go dgvoice.SendPCM(v, send)

	v.Speaking(true)
	defer v.Speaking(false)

	audioFile := "audio/nelson-ha-ha.m4a"

	for {
		p, ok := <-recv
		if !ok {
			log.Println("stopping echo")
			return
		}
		// fmt.Printf("%+v\n", *p)

		// Start loop and attempt to play all files in the given folder
		fmt.Println("PlayAudioFile:", audioFile)

		dgvoice.PlayAudioFile(voiceConnection, audioFile, make(chan bool))
	}
}
