package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type BoxComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *BoxComponent) Render() vecty.ComponentOrHTML {
	vecty.If()
	return elem.Div(vecty.Markup(vecty.Class("box")), t.Markup, t.Slot)
}
