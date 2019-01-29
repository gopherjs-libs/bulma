package bulma

import (
	"github.com/gopherjs/vecty"
)

func OrderedList(l vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &OrderedListComponent{Markup: vecty.Markup(append(applyer, l)...), Slot: c}
	}
}

func UnorderedList(l vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &UnorderedListComponent{Markup: vecty.Markup(append(applyer, l)...), Slot: c}
	}
}

func ListItem(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &ListItemComponent{Slot: c}
}
