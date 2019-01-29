package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type AComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *AComponent) Render() vecty.ComponentOrHTML {
	return elem.Anchor(t.Markup, t.Slot)
}
