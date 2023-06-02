package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		os.Exit(1)
	}

	channelID := "1114230045154750514"
	_, err = discord.ChannelMessageSend(channelID, "Hello, I am always watching you!")
	if err != nil {
		fmt.Println("Error sending message: ", err)
		os.Exit(1)
	}

	err = discord.Close()
	if err != nil {
		fmt.Println("Error closing Discord session: ", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully!")
	os.Exit(0)
}
