package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type ContentComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *ContentComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("content")), t.Markup, t.Slot)
}
