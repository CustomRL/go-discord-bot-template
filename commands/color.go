package commands

import (
	"strings"

	"github.com/CustomRL/go-discord-bot-template/bot"
	"github.com/bwmarrin/discordgo"
)

var colors = []string{"red", "orange", "yellow", "green", "blue", "purple"}

// color shows an option with autocomplete.
var color = &bot.Command{
	ApplicationCommand: discordgo.ApplicationCommand{
		Name:        "color",
		Description: "Pick a color",
		Options: []*discordgo.ApplicationCommandOption{{
			Type:         discordgo.ApplicationCommandOptionString,
			Name:         "name",
			Description:  "Color name",
			Required:     true,
			Autocomplete: true,
		}},
	},
	Run: func(c *bot.Ctx) error {
		return c.Text("You picked " + c.Option("name").StringValue())
	},
	Autocomplete: func(c *bot.Ctx) ([]*discordgo.ApplicationCommandOptionChoice, error) {
		typed := strings.ToLower(c.Option("name").StringValue())
		var choices []*discordgo.ApplicationCommandOptionChoice
		for _, col := range colors {
			if strings.HasPrefix(col, typed) {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{Name: col, Value: col})
			}
		}
		return choices, nil
	},
}
