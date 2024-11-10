package cmd

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type Command struct {
	Name           string
	MinArgs        int
	Usage          string
	IsComplex      bool
	SimpleHandler  func(args []string) (string, error)
	ComplexHandler func(s *discordgo.Session, e *discordgo.MessageCreate, args []string) (string, error)
}

// Register commands here
// ComplexHandler for things that require the
// session and message context
var Commands = map[string]Command{
	"bert": {
		Name:          "bert",
		MinArgs:       0,
		Usage:         "returns a simple message",
		IsComplex:     false,
		SimpleHandler: BertHandler,
	},
	"ciar": {
		Name:          "ciar",
		MinArgs:       0,
		Usage:         "returns a simple message",
		IsComplex:     false,
		SimpleHandler: CiarHandler,
	},
	"con": {
		Name:          "con",
		MinArgs:       0,
		Usage:         "returns a simple message",
		IsComplex:     false,
		SimpleHandler: ConHandler,
	},
	"quot": {
		Name:           "quot",
		MinArgs:        0,
		Usage:          "returns a simple message",
		IsComplex:      true,
		ComplexHandler: QuoteHandler,
	},
	"loves": {
		Name:          "loves",
		MinArgs:       0,
		Usage:         "returns a simple message",
		IsComplex:     false,
		SimpleHandler: LovesHandler,
	},
	"wish": {
		Name:          "wish",
		MinArgs:       0,
		Usage:         "returns a simple message",
		IsComplex:     false,
		SimpleHandler: WishHandler,
	},
	"second": {
		Name:          "second",
		MinArgs:       0,
		Usage:         "returns a simple message",
		IsComplex:     false,
		SimpleHandler: SecondHandler,
	},
	"rtd": {
		Name:          "rtd",
		MinArgs:       1,
		Usage:         "returns a simple message",
		IsComplex:     false,
		SimpleHandler: RtdHandler,
	},
	"default": {
		Name:          "default",
		MinArgs:       0,
		Usage:         "returns a simple message",
		IsComplex:     false,
		SimpleHandler: DefaultHandler,
	},
}

func BertHandler(args []string) (string, error) {
	return "bertin my peanits", nil
}

func CiarHandler(args []string) (string, error) {
	return "i love you", nil
}

func ConHandler(args []string) (string, error) {
	return "Just be a lone wolf rather than alone wolf", nil
}

func QuoteHandler(s *discordgo.Session, e *discordgo.MessageCreate, args []string) (string, error) {
	if e.GuildID != config.ServerId {
		return "", errors.New("Wrong place bucko")
	}

	messages, err := s.ChannelMessages(config.QuotesChannelId, 100, "", "", "")
	if err != nil {
		return "", err
	}

	rnd := rand.Intn(100)
	return messages[rnd].Content, nil
}

func LovesHandler(args []string) (string, error) {
	if rand.Intn(100) > 50 {
		return "loves", nil
	}

	return "hates", nil
}

func WishHandler(args []string) (string, error) {
	rnd := rand.Intn(len(Nouns))
	msg := fmt.Sprintf("I wish I was a %s", Nouns[rnd])
	return msg, nil
}

func SecondHandler(args []string) (string, error) {
	rnd := rand.Intn(len(Verbs))
	msg := fmt.Sprintf("I'll _%s_ you in a second", Verbs[rnd])
	return msg, nil
}

// Parameters:
// args[0]: string in the format <diceCount>d<maxValue>
//
// Returns:
// Will return a string that is the dice roll added up and
// the sum, e.g. 5 + 2 + 3 = 10 with some extra fluff
func RtdHandler(args []string) (string, error) {
	rolls, err := rtd(args[0])
	if err != nil {
		return "", err
	}

	return *rolls, nil
}

func DefaultHandler(args []string) (string, error) {
	return "Unknown command !help for list of commands", nil
}
