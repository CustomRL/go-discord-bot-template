package bot

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) onInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	c := &Ctx{S: s, I: i}
	defer func() {
		if r := recover(); r != nil {
			slog.Error("handler panicked", "panic", r)
			_ = c.Ephemeral("Something went wrong.")
		}
	}()
	if err := b.route(c); err != nil {
		slog.Error("interaction failed", "type", i.Type, "err", err)
		_ = c.Ephemeral("Something went wrong.")
	}
}

func (b *Bot) route(c *Ctx) error {
	switch c.I.Type {
	case discordgo.InteractionApplicationCommand:
		cmd, err := b.command(c)
		if err != nil {
			return err
		}
		return cmd.Run(c)
	case discordgo.InteractionApplicationCommandAutocomplete:
		cmd, err := b.command(c)
		if err != nil {
			return err
		}
		if cmd.Autocomplete == nil {
			return fmt.Errorf("command %q has no autocomplete", cmd.Name)
		}
		choices, err := cmd.Autocomplete(c)
		if err != nil {
			return err
		}
		return c.Respond(&discordgo.InteractionResponse{
			Type: discordgo.InteractionApplicationCommandAutocompleteResult,
			Data: &discordgo.InteractionResponseData{Choices: choices},
		})
	case discordgo.InteractionMessageComponent, discordgo.InteractionModalSubmit:
		key, _ := c.CustomID()
		h := b.handlers[key]
		if h == nil {
			return fmt.Errorf("no handler for custom id %q", key)
		}
		return h(c)
	}
	return nil
}

func (b *Bot) command(c *Ctx) (*Command, error) {
	name := c.I.ApplicationCommandData().Name
	cmd := b.commands[name]
	if cmd == nil {
		return nil, fmt.Errorf("unknown command %q", name)
	}
	return cmd, nil
}
