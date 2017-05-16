package bot

import (
	"RankedBot/config"
	"fmt"

	"strings"

	"bytes"

	"regexp"

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
	var salida bytes.Buffer
	var ok bool

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}
		//var expresion = regexp.MustCompile(`!info\s.*`)
		ok, _ = regexp.MatchString(`!info\s.*`, m.Content)
		if ok == true {
			salida.WriteString("Author: ")
			salida.WriteString(m.Author.String())
			salida.WriteString("\n")
			salida.WriteString("UserName: ")
			salida.WriteString(m.Author.Username)
			salida.WriteString("\n")
			salida.WriteString("ID: ")
			salida.WriteString(m.Author.ID)
			salida.WriteString("\n")
			salida.WriteString("Mensaje: ")
			salida.WriteString(m.Content)
			salida.WriteString("\n")
			_, _ = s.ChannelMessageSend(m.ChannelID, salida.String())
			//_, _ = s.ChannelMessageSend("314110449824301056", "probando probando 123")
		}
	}
	//fmt.Println(m.Content)
}
