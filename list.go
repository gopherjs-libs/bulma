package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type List struct {
	vecty.Core

	IsOrder bool
	State   string
	Type    string
	Size    string

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *List) Render() vecty.ComponentOrHTML {
	_l := EleIf(t.IsOrder, elem.OrderedList, elem.UnorderedList)
	return _l(vecty.Markup(
		vecty.MarkupIf(t.Type != "", vecty.Attribute("type", t.Type)),
		ClassMap(t.Size, t.State),
	), t.Markup, t.Slot)
}
