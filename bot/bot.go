package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/krzek/krzek-bot-discord/config"
)

func Start() {
	// Create a new Discord session using the provided bot token.
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	goBot.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	goBot.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = goBot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content)

	if m.Content == "<@!832356118827237397>" {
		s.ChannelMessageSend(m.ChannelID, "I'm here! :3\n\n You can use the !help command to see the full list of available commands that I can execute! <3")
	}

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		// Ignore all messages created by the bot itself
		// This isn't required in this specific example but it's a good practice.
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "!ping" {
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		}

		if m.Content == "!pong" {
			s.ChannelMessageSend(m.ChannelID, "Ping!")
		}

		if m.Content == "!help" {
			s.ChannelMessageSend(m.ChannelID, "Commands:\n!ping\n!pong\n!pain")
		}

		if m.Content == "!pain" {
			s.ChannelMessageSend(m.ChannelID, "A PAIN É CAMPEÃ, PORRAAAAA!")
		}
	}

}
