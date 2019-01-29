package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type OrderedListComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *OrderedListComponent) Render() vecty.ComponentOrHTML {
	return elem.OrderedList(t.Markup, t.Slot)
}

type UnorderedListComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *UnorderedListComponent) Render() vecty.ComponentOrHTML {
	return elem.UnorderedList(t.Markup, t.Slot)
}

type ListItemComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *ListItemComponent) Render() vecty.ComponentOrHTML {
	return elem.ListItem(t.Markup, t.Slot)
}
