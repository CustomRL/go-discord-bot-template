package commands

import (
	"log/slog"

	"github.com/CustomRL/go-discord-bot-template/bot"
	"github.com/bwmarrin/discordgo"
)

// feedback opens a modal; onFeedback receives the submission.
var feedback = &bot.Command{
	ApplicationCommand: discordgo.ApplicationCommand{Name: "feedback", Description: "Send feedback"},
	Run: func(c *bot.Ctx) error {
		return c.Respond(&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{
				CustomID: "feedback",
				Title:    "Feedback",
				Components: []discordgo.MessageComponent{discordgo.ActionsRow{Components: []discordgo.MessageComponent{
					discordgo.TextInput{CustomID: "text", Label: "What's on your mind?", Style: discordgo.TextInputParagraph, Required: true},
				}}},
			},
		})
	},
}

func onFeedback(c *bot.Ctx) error {
	row := c.I.ModalSubmitData().Components[0].(*discordgo.ActionsRow)
	text := row.Components[0].(*discordgo.TextInput).Value
	slog.Info("feedback", "user", c.User().ID, "text", text)
	return c.Ephemeral("Thanks!")
}
