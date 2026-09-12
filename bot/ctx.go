package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Ctx is one interaction: the session that received it and the payload.
type Ctx struct {
	S *discordgo.Session
	I *discordgo.InteractionCreate
}

func (c *Ctx) Respond(r *discordgo.InteractionResponse) error {
	return c.S.InteractionRespond(c.I.Interaction, r)
}

func (c *Ctx) Reply(data *discordgo.InteractionResponseData) error {
	return c.Respond(&discordgo.InteractionResponse{Type: discordgo.InteractionResponseChannelMessageWithSource, Data: data})
}

func (c *Ctx) Text(s string) error {
	return c.Reply(&discordgo.InteractionResponseData{Content: s})
}

func (c *Ctx) Ephemeral(s string) error {
	return c.Reply(&discordgo.InteractionResponseData{Content: s, Flags: discordgo.MessageFlagsEphemeral})
}

// Defer acknowledges within the 3s window; follow with Followup or Edit.
func (c *Ctx) Defer() error {
	return c.Respond(&discordgo.InteractionResponse{Type: discordgo.InteractionResponseDeferredChannelMessageWithSource})
}

func (c *Ctx) Followup(s string) (*discordgo.Message, error) {
	return c.S.FollowupMessageCreate(c.I.Interaction, true, &discordgo.WebhookParams{Content: s})
}

func (c *Ctx) Edit(s string) (*discordgo.Message, error) {
	return c.S.InteractionResponseEdit(c.I.Interaction, &discordgo.WebhookEdit{Content: &s})
}

// Option finds a slash option by name, descending through subcommands.
func (c *Ctx) Option(name string) *discordgo.ApplicationCommandInteractionDataOption {
	return findOption(c.I.ApplicationCommandData().Options, name)
}

func findOption(opts []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, o := range opts {
		if o.Name == name {
			return o
		}
		if found := findOption(o.Options, name); found != nil {
			return found
		}
	}
	return nil
}

// CustomID splits a component or modal ID "key:arg" into its parts.
func (c *Ctx) CustomID() (key, arg string) {
	var id string
	switch c.I.Type {
	case discordgo.InteractionMessageComponent:
		id = c.I.MessageComponentData().CustomID
	case discordgo.InteractionModalSubmit:
		id = c.I.ModalSubmitData().CustomID
	}
	key, arg, _ = strings.Cut(id, ":")
	return key, arg
}

// User works for both guild and DM interactions.
func (c *Ctx) User() *discordgo.User {
	if c.I.Member != nil {
		return c.I.Member.User
	}
	return c.I.User
}
