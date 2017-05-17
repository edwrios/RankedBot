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
var GuildID = "313889492375175169"

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
	var err error

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}
		//var expresion = regexp.MustCompile(`!info\s.*`)
		//ok, _ = regexp.MatchString(`!ready\s.*`, m.Content)
		ok, _ = regexp.MatchString(`!ready`, m.Content)
		if ok == true {
			//var g *discordgo.Guild

			var gm, _ = s.GuildMember(GuildID, m.Author.ID)

			salida.WriteString("Author: ")
			salida.WriteString(m.Author.String())
			salida.WriteString("\n")
			salida.WriteString("UserName: ")
			salida.WriteString(m.Author.Username)
			salida.WriteString("\n")
			salida.WriteString("Nick: ")
			salida.WriteString(gm.Nick)
			salida.WriteString("\n")
			salida.WriteString("ID: ")
			salida.WriteString(m.Author.ID)
			salida.WriteString("\n")
			salida.WriteString("Mensaje: ")
			salida.WriteString(m.Content)
			salida.WriteString("\n")

			if gm.Nick == "" {
				err = s.GuildMemberNickname(GuildID, m.Author.ID, m.Author.Username+" (R)")
				if err != nil {
					fmt.Println(err.Error())
					_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
					return
				}
			} else {
				err = s.GuildMemberNickname(GuildID, m.Author.ID, gm.Nick+" (R)")
				if err != nil {
					fmt.Println(err.Error())
					_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
					return
				}
			}

			_, _ = s.ChannelMessageSend(m.ChannelID, salida.String())
			//_, _ = s.ChannelMessageSend("314110449824301056", "probando probando 123") //mensaje en canal #prueba

		}
	}
	//fmt.Println(m.Content)
}
