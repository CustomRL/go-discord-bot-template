package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestRouteComponentByPrefix(t *testing.T) {
	b := &Bot{handlers: map[string]Handler{}}
	var got string
	b.Handle("vote", func(c *Ctx) error { _, got = c.CustomID(); return nil })

	c := &Ctx{I: &discordgo.InteractionCreate{Interaction: &discordgo.Interaction{
		Type: discordgo.InteractionMessageComponent,
		Data: discordgo.MessageComponentInteractionData{CustomID: "vote:yes"},
	}}}
	if err := b.route(c); err != nil || got != "yes" {
		t.Fatalf("route: err=%v arg=%q", err, got)
	}
	c.I.Data = discordgo.MessageComponentInteractionData{CustomID: "nope"}
	if err := b.route(c); err == nil {
		t.Fatal("expected error for unregistered custom id")
	}
}

func TestOptionDescendsSubcommands(t *testing.T) {
	c := &Ctx{I: &discordgo.InteractionCreate{Interaction: &discordgo.Interaction{
		Type: discordgo.InteractionApplicationCommand,
		Data: discordgo.ApplicationCommandInteractionData{Options: []*discordgo.ApplicationCommandInteractionDataOption{{
			Name: "sub", Type: discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandInteractionDataOption{{Name: "n", Type: discordgo.ApplicationCommandOptionInteger, Value: float64(7)}},
		}}},
	}}}
	if o := c.Option("n"); o == nil || o.IntValue() != 7 {
		t.Fatalf("Option(n) = %v", o)
	}
	if c.Option("missing") != nil {
		t.Fatal("expected nil for missing option")
	}
}
