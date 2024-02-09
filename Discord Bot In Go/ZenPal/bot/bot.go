package bot

import (
	"ZenPal/config"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session
var err error

func Start() {
	goBot, err = discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("Error creating Discord session", err)
		return
	}

	//goBot.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent
	//issue fixed by modifying bot settings on developer portal

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {

		if m.Author.ID == BotID {
			return
		}

		if m.Content == config.BotPrefix+"ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")

		} else if m.Content == "<version" {
			_, _ = s.ChannelMessageSend(m.ChannelID, config.Version)

		} else if m.Content == "<owner" {
			_, _ = s.ChannelMessageSend(m.ChannelID, config.Owner)

		}
	}

}
