package bulma

import (
	"github.com/gopherjs/vecty"
)

type TextComponent struct {
	vecty.Core

	text string
	Slot []vecty.MarkupOrChild
}

func (t *TextComponent) Render() vecty.ComponentOrHTML {
	return vecty.Text(t.text, t.Slot...)
}
