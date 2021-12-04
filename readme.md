# Nelson Mandela Muntz

An annoying discord bot that talks over your friends

Makes use of [discordgo](https://github.com/bwmarrin/discordgo) package

Based off code from [voice_receive](https://github.com/bwmarrin/discordgo/tree/master/examples/voice_receive) example

## Build
To build, make sure that modules are enabled and run:
```
go build
```

## Usage

To run with secrets:
```
TOKEN=YOUR_TOKEN GUILD_ID=123123123 CHANNEL_ID=456456456 go run main.go 
```

Or to run a built binary:
```
TOKEN=YOUR_TOKEN GUILD_ID=123123123 CHANNEL_ID=456456456 ./nelson-mandela-muntz
```

To run while in development with secrets dir:
```
TOKEN=$(cat ./secrets/TOKEN) GUILD_ID=$(cat ./secrets/GUILD_ID) CHANNEL_ID=$(cat ./secrets/CHANNEL_ID) go run main.go 
```
