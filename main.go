package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Token           string `json:"token"`
	BotPrefix       string `json:"botPrefix"`
	ServerId        string `json:"serverId"`
	QuotesChannelId string `json:"quotesChannelId"`
}

func ReadConfig() (*Config, error) {
	fmt.Println("Reading config...")
	data, err := os.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}
	fmt.Println("Json decode...")
	var cfg Config
	err = json.Unmarshal([]byte(data), &cfg)
	if err != nil {
		fmt.Println("Error decoding json")
		return nil, err
	}

	return &cfg, nil
}

var BotId string
var config Config

func main() {
	configPtr, err := ReadConfig()
	if err != nil {
		fmt.Println("Failed reading config: ", err)
		return
	}
	config = *configPtr

	soup, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Failed initialising Discord session: ", err)
		return
	}

	u, err := soup.User("@me")
	if err != nil {
		fmt.Println("Failed getting current user: ", err)
		return
	}

	BotId = u.ID

	soup.AddHandler(messageHandler)

	err = soup.Open()
	if err != nil {
		fmt.Println("Failed connecting to Discord: ", err)
	}

	<-make(chan struct{})
}

func messageHandler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.ID == BotId {
		return
	}

	prefix := config.BotPrefix
	if strings.HasPrefix(e.Content, prefix) {
		args := strings.Fields(e.Content)[strings.Index(e.Content, prefix):]
		cmd := args[0][len(prefix):]
		//arguments := args[1:]

		switch cmd {
		case "bert":
			_, err := s.ChannelMessageSend(e.ChannelID, "Bertin my peanits, straight up zorkin it!")
			ErrorHandler("Failed sending bert Command Response: ", err)

		case "ciar":
			_, err := s.ChannelMessageSend(e.ChannelID, "i love you")
			ErrorHandler("Failed sending ciar Command Response: ", err)

		case "con":
			_, err := s.ChannelMessageSend(e.ChannelID, "Just be a lone wolf rather than alone wolf")
			ErrorHandler("Failed sending con Command Response: ", err)

		case "rtd":
			messages, err := s.ChannelMessages(config.QuotesChannelId, 100, "", "", "")
			ErrorHandler("Failed getting quotes: ", err)

			rnd := rand.Intn(100)

			_, err = s.ChannelMessageSend(e.ChannelID, messages[rnd].Content)
			ErrorHandler("Failed sending quotes: ", err)
		case "loves":
			var msg string
			rnd := rand.Intn(100)
			if rnd > 50 {
				msg = "loves"
			} else {
				msg = "hates"
			}
			_, err := s.ChannelMessageSend(e.ChannelID, msg)
			ErrorHandler("Failed sending quotes: ", err)
		default:
			_, err := s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Unknown command %q.", cmd))
			ErrorHandler("Failed sending Unknown Command Response: ", err)
		}

	}
}

func ErrorHandler(errMsg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v", errMsg, err)
	}
}