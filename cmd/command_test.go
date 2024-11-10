package cmd

import (
	"math/rand"
	"testing"

	"github.com/bwmarrin/discordgo"
)

var VerbsTest = []string{"test"}
var NountsTest = []string{"test"}

var configMock = Config{
	Token:           "Hello",
	BotPrefix:       "!",
	ServerId:        "1234567890",
	QuotesChannelId: "1234567890",
}

// Check if commands are registered
func TestCommands(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "not found",
			input:    "asfdasdfasdf",
			expected: false,
		},
		{
			name:     "bert found",
			input:    "bert",
			expected: true,
		},
		{
			name:     "ciar found",
			input:    "ciar",
			expected: true,
		},
		{
			name:     "con found",
			input:    "bert",
			expected: true,
		},
		{
			name:     "quot found",
			input:    "quot",
			expected: true,
		},
		{
			name:     "wish found",
			input:    "wish",
			expected: true,
		},
		{
			name:     "loves found",
			input:    "loves",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := Commands[tt.input]
			if ok != tt.expected {
				t.Fatalf("Commands not registeres: %s", tt.name)
			}
		})
	}
}

func TestBert(t *testing.T) {
	// Act
	msg, err := BertHandler([]string{})
	if err != nil {
		t.Errorf("BertHandler returned error: %v", err)
	}
	if msg != "bertin my peanits" {
		t.Error("BertHandler returned invalid msg")
	}
}

func TestCiar(t *testing.T) {
	// Act
	msg, err := CiarHandler([]string{})
	if err != nil {
		t.Errorf("CiarHandler returned error: %v", err)
	}
	if msg != "i love you" {
		t.Error("CiarHandler returned invalid msg")
	}
}

func TestCon(t *testing.T) {
	// Act
	msg, err := ConHandler([]string{})
	if err != nil {
		t.Errorf("ConHandler returned error: %v", err)
	}
	if msg != "Just be a lone wolf rather than alone wolf" {
		t.Error("Conhandler returned invalid msg")
	}
}

func TestLoves(t *testing.T) {
	msg, err := LovesHandler([]string{})
	if err != nil {
		t.Errorf("LoveHandler returned error: %v", err)
	}
	if (msg != "loves") && (msg != "hates") {
		t.Error("LoveHandler returned invalid msg")
	}
}

func TestQuot(t *testing.T) {
	rand.NewSource(1)

	e := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			GuildID: config.ServerId,
		},
	}

	msg, err := QuoteHandler(s, e, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if msg == "" {
		t.Fatalf("Expected a message, got empty string")
	}

	messages, _ := s.ChannelMessages("", 0, "", "", "")
	found := false
	for _, message := range messages {
		if msg == message.Content {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Result %q not found in messages", msg)
	}

	e.Message.GuildID = "wrong-guild-id"
	msg, err = QuoteHandler(s, e, nil)
	if err == nil || err.Error() != "Wrong place bucko" {
		t.Fatalf("expected error 'Wrong place bucko', got %v", err)
	}
}

func TestRtd(t *testing.T) {
	rtdValue := "2d6"

	msg, err := RtdHandler([]string{rtdValue})
	if err != nil {
		t.Errorf("RtdHandler returned error: %v", err)
	}

	if msg == "" {
		t.Error("RtdHandler returned invalid msg")
	}
}
func TestWish(t *testing.T) {
	Nouns = NountsTest
	msg, err := WishHandler([]string{})
	if err != nil {
		t.Errorf("WishHandler returned error: %v", err)
	}

	if msg != "I wish I was a test" {
		t.Errorf("WishHandler returnd invalid msg: %s", msg)
	}

}
func TestDefault(t *testing.T) {
}
