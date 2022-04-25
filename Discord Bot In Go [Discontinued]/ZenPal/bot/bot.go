package bot

import (
	"fmt"
	"strings"

	"ZenPal/config"

	"github.com/bwmarrin/discordgo"
)

// BotID is ...
var BotID string
var goBot *discordgo.Session

// Start is ...
func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("Error creating Discord session", err)
		return
	}

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

		if m.Content == "<ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")

		} else if m.Content == "<version" {
			_, _ = s.ChannelMessageSend(m.ChannelID, config.Version)

		} else if m.Content == "<owner" {
			_, _ = s.ChannelMessageSend(m.ChannelID, config.Owner)

			/*else if m.Content == "<avatar "+ userID {
			// UserAvatarDecode returns an image.Image of a user's Avatar
			// user : The user which avatar should be retrieved
			// UserAvatarDecode(userID)
			/*func UserAvatarDecode(u *User,err error)

			if err != nil {
				fmt.Println(err.Error())
				return
			}*/

		} else if m.Content == "<pfp" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Paste a link to set your pfp")

			_, err = dg.UserUpdate("", "", "", avatar, "")
			if err != nil {
				fmt.Println("Error updating pfp through file")
			}

		}
	}

}

// UserAvatarDecode returns an image.Image of a user's Avatar
// user : The user which avatar should be retrieved
// UserAvatarDecode(userID)
//func UserAvatarDecode(s *discordgoSession, u *User, img image.Image) (err error) {

/*body, err := s.RequestWithBucketID("GET", EndpointUserAvatar(u.ID, u.Avatar), nil, EndpointUserAvatar("", ""))

	if err != nil {
		return
	}

	img, _, err = image.Decode(bytes.NewReader(body))
	return
} */
