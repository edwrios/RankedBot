package bot

import (
	"RankedBot/config"
	"fmt"

	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
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

	fmt.Println("RankedBot se está ejecutando...")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}
		if m.Content == "!prueba" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "probando probando 123")
			//_, _ = s.ChannelMessageSend("314110449824301056", "probando probando 123")
		}
	}
	//fmt.Println(m.Content)
}
