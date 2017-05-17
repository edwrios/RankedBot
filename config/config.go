package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string
	GuildID   string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	GuildID   string `json:"GuildID"`
}

func ReadConfig() error {
	fmt.Println("Cargando arvhivo de configuracion...")

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	GuildID = config.GuildID

	return nil
}
