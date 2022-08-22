package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/andersfylling/snowflake"
	"github.com/incidrthreat/discordhook"
)

// struct to define msg data to be sent
type SendMsgData struct {
	Title    string
	Url      string
	Quantity string
}

func SendDiscordMsg(s *SendMsgData) {
	id_token := strings.Replace(DISCORD_WEBHOOK, "https://discord.com/api/webhooks/", "", -1) // removes the prefix to from the URL
	dataSplit := strings.Split(id_token, "/")                                                 // Splits string into 2 parts, Discord ID and Discord Token
	discoID_Uint, _ := strconv.ParseUint(dataSplit[0], 10, 64)

	sm, err := discordhook.NewWebhookAPI(snowflake.Snowflake(discoID_Uint), dataSplit[1], true, nil)
	if err != nil {
		log.Println(err)
	}
	t := time.Now()

	// Sends the msg to discord via provided webhooks
	sm.Execute(context.Background(), &discordhook.WebhookExecuteParams{
		Embeds: []*discordhook.Embed{
			{
				Type:  "rich",
				Title: s.Title,
				Color: 13369344, // Target hexcolor code converted to Hex
				Fields: []*discordhook.EmbedField{
					{
						Name:   "URL",
						Value:  s.Url,
						Inline: false,
					},
					{
						Name:   "Quantity",
						Value:  s.Quantity,
						Inline: false,
					},
				},
				Footer: &discordhook.EmbedFooter{
					Text: "Product prices and availability are accurate as of the date/time indicated and are subject to change.",
				},
				Timestamp: &t,
			},
		},
	}, nil, "")
}
