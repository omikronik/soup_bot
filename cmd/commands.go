package cmd

import "github.com/bwmarrin/discordgo"

var commandMap = map[string]func(s *discordgo.Session, e *discordgo.MessageCreate) error{
	"help": help,
}

func help(s *discordgo.Session, e *discordgo.MessageCreate) error {
	help := `
	bert: bert
	ciar: ciar
	con: con
	quot:
	rolls the dice to show a fun quote from the last 100 quotes!
	rtd:
	actually rolls the dice, format <diceCount>d<diceSize>, e.g. 2d20
	loves: 50/50
	wish: generate a lovely wish :)
	second: I'll noun you in a second
	`
	_, err := s.ChannelMessageSend(e.ChannelID, help)
	if err != nil {
		return err
	}
	return nil
}
