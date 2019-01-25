package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type OrderedList struct {
	vecty.Core

	State string
	Type  string
	Size  string

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *OrderedList) Render() vecty.ComponentOrHTML {
	return elem.OrderedList(
		vecty.Markup(
			vecty.MarkupIf(t.Type != "", vecty.Attribute("type", t.Type)),
			ClassMap(t.Size, t.State),
		),
		t.Markup, t.Slot,
	)
}
