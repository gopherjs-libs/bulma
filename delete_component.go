package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type DeleteComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *DeleteComponent) Render() vecty.ComponentOrHTML {
	return elem.Button(vecty.Markup(vecty.Class("delete")),t.Markup,t.Slot)
}
