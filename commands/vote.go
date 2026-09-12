package commands

import (
	"github.com/CustomRL/go-discord-bot-template/bot"
	"github.com/bwmarrin/discordgo"
)

// vote shows message components; onVote handles every button via the "vote:" prefix.
var vote = &bot.Command{
	ApplicationCommand: discordgo.ApplicationCommand{Name: "vote", Description: "Start a yes/no vote"},
	Run: func(c *bot.Ctx) error {
		return c.Reply(&discordgo.InteractionResponseData{
			Content: "Cast your vote:",
			Components: []discordgo.MessageComponent{discordgo.ActionsRow{Components: []discordgo.MessageComponent{
				discordgo.Button{Label: "Yes", Style: discordgo.SuccessButton, CustomID: "vote:yes"},
				discordgo.Button{Label: "No", Style: discordgo.DangerButton, CustomID: "vote:no"},
			}}},
		})
	},
}

func onVote(c *bot.Ctx) error {
	_, choice := c.CustomID()
	return c.Ephemeral("You voted " + choice)
}
