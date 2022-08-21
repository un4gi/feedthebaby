package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ConnectDiscord(m string) {
	d, err := discordgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		fmt.Println("Error starting new bot session,", err)
	}

	user, _ := d.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = d.Open()
	if err != nil {
		fmt.Println("Error connecting to Discord.")
	}

	fmt.Println(fmt.Sprint(user), "has connected to Discord!")
	SendMessage(d, m)
}

func SendMessage(d *discordgo.Session, m string) {
	_, err := d.ChannelMessageSend(CHANNEL_ID, m)
	if err != nil {
		fmt.Println(err)
	}
}
