package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type ButtonsModel struct {
	AreSmall  bool
	AreMedium bool
	AreLarge  bool
}

type Buttons struct {
	vecty.Core
	Size string

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *Buttons) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(
		vecty.Class("buttons"),
		ClassMap(t.Size),
	), t.Markup, t.Slot)
}
