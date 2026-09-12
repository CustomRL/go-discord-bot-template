// Package commands holds every slash command and component handler.
// Add a command: create a file, append it to All.
// Add a component or modal handler: register it in RegisterHandlers.
package commands

import "github.com/CustomRL/go-discord-bot-template/bot"

var All = []*bot.Command{ping, color, vote, feedback}

func RegisterHandlers(b *bot.Bot) {
	b.Handle("vote", onVote)
	b.Handle("feedback", onFeedback)
}
