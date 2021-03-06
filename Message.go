package SimpleChatCommon

import (
	"fmt"
	"github.com/gabriel-comeau/tbuikit"
	"github.com/nsf/termbox-go"
	"regexp"
	"strings"
)

// Separator string to split the message up
const MESSAGE_TOKEN string = "|%|"

//
const WHO_MESSAGE_TOKEN string = "|%WHO%|"
const WHO_MESSAGE_REGEX_STR string = `^\|%WHO%\|`

// Converts a colorized string object to a string
func Pack(colorizedString *tbuikit.ColorizedString) string {
	colors := new(ColorList)
	colStr := colors.FromColor(colorizedString.Color)

	return fmt.Sprintf("%v%v%v\n", colorizedString.Text, MESSAGE_TOKEN, colStr)
}

// Converts a string to a ColorizedString object
func Unpack(raw string) *tbuikit.ColorizedString {
	cs := new(tbuikit.ColorizedString)
	trimmed := strings.Trim(raw, " \n")
	parts := strings.Split(trimmed, MESSAGE_TOKEN)

	cs.Color = termbox.ColorWhite

	for i, part := range parts {
		if i == 0 {
			cs.Text = part
		} else if i == 1 {
			colors := new(ColorList)
			col := colors.Get(part)
			cs.Color = col
		}
	}

	return cs
}

// Check if this message from the server is a "who" message
func IsWhoMessage(colorizedString *tbuikit.ColorizedString) bool {
	whoMsgRegex := regexp.MustCompile(WHO_MESSAGE_REGEX_STR)
	if whoMsgRegex.MatchString(colorizedString.Text) {
		return false
	} else {
		return false
	}
}

// Creates a new colorized string object, from a body string and color string
func Create(text, color string) *tbuikit.ColorizedString {
	cs := new(tbuikit.ColorizedString)
	cs.Text = text

	cs.Color = termbox.ColorWhite
	colors := new(ColorList)
	col := colors.Get(color)
	cs.Color = col

	return cs
}

// For convenience, takes a slice of strings and a single color and then
// converts them to slice of message objects.  Obviously, they'll all have
// the same color but in many cases this is desirable.
func ColStringSliceFromStringSlice(texts []string, color string) []*tbuikit.ColorizedString {
	cstrings := make([]*tbuikit.ColorizedString, 0)
	for _, t := range texts {
		cstrings = append(cstrings, Create(t, color))
	}

	return cstrings
}
