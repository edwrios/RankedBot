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
var ReadyLista []player

type player struct {
	Username string
	ID       string
	Nick     string
	Author   string
}

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
	fmt.Println(`Presionar CTRL+C para salir.`)
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var salida bytes.Buffer
	var ok bool
	var err error
	var gm, _ = s.GuildMember(config.GuildID, m.Author.ID)
	//var g *discordgo.Guild

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}

		ok, _ = regexp.MatchString(`!r`, m.Content)
		if ok {
			tiene, _ := regexp.MatchString(`.*\s\[R\]`, gm.Nick)
			if gm.Nick == "" {
				err = s.GuildMemberNickname(config.GuildID, m.Author.ID, m.Author.Username+" [R]")
				if err != nil {
					fmt.Println(err.Error())
					_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
					return
				}
				_, _ = s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> está listo para aceptar ranked matches.")
				//agregar a ReadyList
				jug := player{Username: m.Author.Username, Nick: gm.Nick, ID: m.Author.ID, Author: m.Author.String()}
				ReadyLista = append(ReadyLista, jug)
			} else if !tiene {
				err = s.GuildMemberNickname(config.GuildID, m.Author.ID, gm.Nick+" [R]")
				if err != nil {
					fmt.Println(err.Error())
					_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
					return
				}
				_, _ = s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> está listo para aceptar ranked matches.")
				//agregar a ReadyList
				jug := player{Username: m.Author.Username, Nick: gm.Nick, ID: m.Author.ID, Author: m.Author.String()}
				ReadyLista = append(ReadyLista, jug)
			}
		}

		ok, _ = regexp.MatchString(`!nr`, m.Content)
		if ok && gm.Nick != "" {
			tiene, _ := regexp.MatchString(`.*\s\[R\]`, gm.Nick)
			if tiene {
				temp := strings.TrimSuffix(gm.Nick, " [R]")
				err = s.GuildMemberNickname(config.GuildID, m.Author.ID, temp)
				if err != nil {
					fmt.Println(err.Error())
					_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
					return
				}

				for i := range ReadyLista {
					if ReadyLista[i].Author == m.Author.String() {
						ReadyLista = append(ReadyLista[:i], ReadyLista[i+1:]...)
						//break
					}
				}
				/*
					for i := len(ReadyLista) - 1; i >= 0; i-- {
						//actual := ReadyLista[i]
						if ReadyLista[i].Author == m.Author.String() {
							ReadyLista = append(ReadyLista[:i], ReadyLista[i+1:]...)
							fmt.Println("encontrado: " + ReadyLista[i].Author)

						}
					}*/
				_, _ = s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> ya no aceptará más retos.")
			}
		}

		ok, _ = regexp.MatchString(`!info`, m.Content)
		if ok {
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
			salida.WriteString("Token: ")
			salida.WriteString(m.Author.Token)
			salida.WriteString("\n")
			salida.WriteString("Mensaje: ")
			salida.WriteString(m.Content)
			salida.WriteString("\n")

			_, _ = s.ChannelMessageSend(m.ChannelID, salida.String())
			//_, _ = s.ChannelMessageSend("314110449824301056", "probando probando 123") //mensaje en canal #prueba
		}

		ok, _ = regexp.MatchString(`!l`, m.Content)
		if ok {
			if len(ReadyLista) == 0 {
				_, _ = s.ChannelMessageSend(m.ChannelID, "No hay jugadores disponibles para ranked.")
			} else {
				printLista := ""
				//Listar jugadores listos
				for i := 0; i < len(ReadyLista); i++ {
					fmt.Println(ReadyLista[i].Username + " " + ReadyLista[i].Nick + " " + ReadyLista[i].Author)
					printLista = printLista + ReadyLista[i].Username + " " + ReadyLista[i].Nick + " " + ReadyLista[i].Author + "\n"
				}
				fmt.Println()
				_, _ = s.ChannelMessageSend(m.ChannelID, printLista)
			}
		}

	}
	//fmt.Println(m.Content)
}
