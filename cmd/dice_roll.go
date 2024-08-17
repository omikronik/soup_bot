package cmd

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

var firstQuartileMessages = []string{
	"Oof, did your character accidentally roll a potato instead of dice? Maybe they should have stayed in bed today.",
	"Looks like the dice gods have spoken, and they said, Nope!",
	"Your roll is so low, even goblins are laughing at you.",
	"Not even a healing potion can fix that disaster of a roll. Better luck next time, brave adventurer!",
}

var secondQuartileMessages = []string{
	"It's not great, but at least you didn't trip over your own feet this time!",
	"Your roll is like a tavern ale, average at best, but it'll get the job done, sort of.",
	"Hey, it could be worse… You could be rolling in the first quartile again.",
	"Your hero takes a cautious step forward, and immediately regrets it.",
}

var thirdQuartileMessages = []string{
	"Not bad! Your roll has some heroic potential. The bards might sing about this, maybe.",
	"Solid roll! Your adventurer might actually survive this encounter. Keep it up!",
	"Your dice whisper, 'Not today, death!' and deliver a roll worthy of a competent adventurer.",
	"You're on the path to greatness! Just a few more rolls like this, and you might actually impress the party.",
}

var fourthQuartileMessages = []string{
	"Legendary! The dice practically bow to your hero's overwhelming prowess.",
	"That roll is so good, the DM just gave you an approving nod, and that never happens!",
	"You've summoned the favor of the dice gods, behold their blessing and smite your enemies!",
	"This is the kind of roll that makes dragons second-guess their life choices.",
}

type DiceRoll struct {
	DiceCount  int
	DiceSize   int
	DiceRolls  []int
	DiceOutput string
}

func NewDiceRoll(diceCount int, diceSize int) (*DiceRoll, error) {
	if diceCount == 0 || diceSize == 0 {
		return nil, errors.New("invalid diceCount or diceSize")
	}

	d := &DiceRoll{
		DiceCount:  diceCount,
		DiceSize:   diceSize,
		DiceRolls:  []int{},
		DiceOutput: "",
	}

	d.RollDice()

	return d, nil
}

func (d *DiceRoll) RollDice() {
	for range d.DiceCount {
		d.DiceRolls = append(d.DiceRolls, (rand.Intn(d.DiceSize) + 1))
	}

	var builder strings.Builder
	builder.WriteString("And your roll is: ")
	total := 0
	for i, v := range d.DiceRolls {
		total += v
		if i == (len(d.DiceRolls) - 1) {
			builder.WriteString(fmt.Sprintf("%d = %d", v, total))
		} else {
			builder.WriteString(fmt.Sprintf("%d + ", v))
		}

	}

	potentialMax := d.DiceCount * d.DiceSize
	quartile := potentialMax / 4

	switch {
	case total == d.DiceCount:
		builder.WriteString("damn bro...")
	case total <= quartile:
		builder.WriteString(getRandomMessage(firstQuartileMessages))
	case total <= quartile*2:
		builder.WriteString(getRandomMessage(secondQuartileMessages))
	case total <= quartile*3:
		builder.WriteString(getRandomMessage(thirdQuartileMessages))
	case total <= quartile*4:
		builder.WriteString(getRandomMessage(fourthQuartileMessages))
	}

	d.DiceOutput = builder.String()
}

func rtd(dice string) (*string, error) {
	reg := regexp.MustCompile(`^([1-9]\d*)d([1-9]\d*)$`)

	matches := reg.FindStringSubmatch(dice)

	if len(matches) != 3 {
		return nil, errors.New("Invalid input :(")
	}

	var diceSettings []int
	for i := 1; i < 3; i++ {
		intConv, err := strconv.Atoi(matches[i])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Invalid input, bad int conversion: %s", matches[i]))
		}
		diceSettings = append(diceSettings, intConv)
	}

	rolls, err := NewDiceRoll(
		diceSettings[0],
		diceSettings[1],
	)

	if err != nil {
		return nil, errors.New("Invalid input :(")
	}

	return &rolls.DiceOutput, nil
}

func getRandomMessage(messages []string) string {
	return messages[rand.Intn(len(messages))]
}
