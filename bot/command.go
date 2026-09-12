package bot

import "github.com/bwmarrin/discordgo"

type Handler func(c *Ctx) error

type Command struct {
	discordgo.ApplicationCommand
	Run          Handler
	Autocomplete func(c *Ctx) ([]*discordgo.ApplicationCommandOptionChoice, error)
}

// syncCommands overwrites the registered set in one API call, which also
// removes commands that no longer exist in code.
func (b *Bot) syncCommands() error {
	defs := make([]*discordgo.ApplicationCommand, 0, len(b.commands))
	for _, c := range b.commands {
		defs = append(defs, &c.ApplicationCommand)
	}
	_, err := b.Session.ApplicationCommandBulkOverwrite(b.Session.State.User.ID, b.guildID, defs)
	return err
}
