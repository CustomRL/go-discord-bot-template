package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/CustomRL/go-discord-bot-template/bot"
	"github.com/CustomRL/go-discord-bot-template/commands"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		slog.Error("DISCORD_TOKEN is not set")
		os.Exit(1)
	}
	b, err := bot.New(token, os.Getenv("DISCORD_GUILD_ID"), commands.All...)
	if err != nil {
		slog.Error("create bot", "err", err)
		os.Exit(1)
	}
	commands.RegisterHandlers(b)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := b.Run(ctx); err != nil {
		slog.Error("run", "err", err)
		os.Exit(1)
	}
}
