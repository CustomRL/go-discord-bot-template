package commands

import (
	"github.com/CustomRL/go-discord-bot-template/bot"
	"github.com/bwmarrin/discordgo"
)

var ping = &bot.Command{
	ApplicationCommand: discordgo.ApplicationCommand{Name: "ping", Description: "Check the bot is alive"},
	Run: func(c *bot.Ctx) error {
		return c.Text("Pong! " + c.S.HeartbeatLatency().String())
	},
}
