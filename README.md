# go-discord-bot-template

Minimal, fast Discord bot skeleton in Go on [discordgo](https://github.com/bwmarrin/discordgo).
One map lookup per interaction, no reflection, no framework.

Supports slash commands (with subcommands and options), autocomplete, buttons and
select menus, modals, deferred replies, and raw gateway events.

## Run

```sh
cp .env.example .env   # fill in DISCORD_TOKEN, optionally DISCORD_GUILD_ID
export $(cat .env | xargs)
go run .
```

Commands are synced on startup with one bulk overwrite, so removing a command
from code removes it from Discord.

## Layout

```
main.go             token, signal handling, run
bot/                session, dispatcher, Ctx helpers
commands/           one file per command; register in commands.go
```

## Add a slash command

```go
var hello = &bot.Command{
	ApplicationCommand: discordgo.ApplicationCommand{Name: "hello", Description: "Say hi"},
	Run: func(c *bot.Ctx) error { return c.Text("hi " + c.User().Mention()) },
}
```

Append it to `All` in `commands/commands.go`.

## Buttons, selects, modals

Give components a `CustomID` of `key` or `key:arg`, then `b.Handle("key", fn)`.
Inside `fn`, `c.CustomID()` returns the key and arg. See `commands/vote.go`
and `commands/feedback.go`.

## Slow work

Discord requires a response within 3 seconds. Call `c.Defer()`, do the work,
then `c.Edit(...)` or `c.Followup(...)`.

## Other gateway events

`b.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) { ... })`.
Add the matching intent in `bot.New` if the event needs one.

## Ctx cheat sheet

| Method | Use |
|---|---|
| `Text(s)` / `Ephemeral(s)` | plain reply, public or private |
| `Reply(data)` | reply with embeds, components, files |
| `Respond(resp)` | raw response, e.g. open a modal |
| `Defer()` then `Edit(s)` / `Followup(s)` | work that takes longer than 3s |
| `Option(name)` | slash option, searches into subcommands |
| `CustomID()` | component or modal `key, arg` |
| `User()` | invoking user, works in guilds and DMs |
