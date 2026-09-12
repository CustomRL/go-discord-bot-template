package bot

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session  *discordgo.Session
	guildID  string
	commands map[string]*Command
	handlers map[string]Handler
}

// New builds a session with the given commands. guildID "" registers
// commands globally (propagation takes up to an hour); set it while developing.
func New(token, guildID string, cmds ...*Command) (*Bot, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	s.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages
	b := &Bot{
		Session:  s,
		guildID:  guildID,
		commands: make(map[string]*Command, len(cmds)),
		handlers: map[string]Handler{},
	}
	for _, c := range cmds {
		b.commands[c.Name] = c
	}
	s.AddHandler(b.onInteraction)
	s.AddHandler(func(_ *discordgo.Session, r *discordgo.Ready) {
		slog.Info("ready", "user", r.User.String(), "guilds", len(r.Guilds))
	})
	return b, nil
}

// Handle routes components and modals whose CustomID is "key" or "key:arg" to h.
func (b *Bot) Handle(key string, h Handler) { b.handlers[key] = h }

// Run opens the gateway, syncs slash commands, and blocks until ctx is done.
func (b *Bot) Run(ctx context.Context) error {
	if err := b.Session.Open(); err != nil {
		return err
	}
	defer b.Session.Close()
	if err := b.syncCommands(); err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}
