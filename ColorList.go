package SimpleChatCommon

import (
	"github.com/nsf/termbox-go"
)

type ColorList struct {
	colors   map[string]termbox.Attribute
	fallback termbox.Attribute
	isInited bool
}

func (this *ColorList) setColors() {
	this.colors = make(map[string]termbox.Attribute)
	this.colors["black"] = termbox.ColorBlack
	this.colors["blue"] = termbox.ColorBlue
	this.colors["cyan"] = termbox.ColorCyan
	this.colors["default"] = termbox.ColorDefault
	this.colors["green"] = termbox.ColorGreen
	this.colors["magenta"] = termbox.ColorMagenta
	this.colors["red"] = termbox.ColorRed
	this.colors["white"] = termbox.ColorWhite
	this.colors["yellow"] = termbox.ColorYellow

	this.fallback = termbox.ColorWhite
	this.isInited = true
}

func (this *ColorList) IsColor(colStr string) bool {
	if !this.isInited {
		this.setColors()
	}

	ret := false

	if this.colors[colStr] != 0 {
		ret = true
	}

	return ret
}

func (this *ColorList) Get(color string) termbox.Attribute {
	if !this.isInited {
		this.setColors()
	}

	attr := this.fallback

	if this.colors[color] != 0 {
		attr = this.colors[color]
	}

	return attr
}

func (this *ColorList) FromColor(col termbox.Attribute) string {
	if !this.isInited {
		this.setColors()
	}

	str := ""

	for k, att := range this.colors {
		if att == col {
			str = k
		}
	}

	return str
}
