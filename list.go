package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func OrderedList(l vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.OrderedList(vecty.Markup(append(applyer, l)...), Components(c...))
	}
}

func UnorderedList(l vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.UnorderedList(vecty.Markup(append(applyer, l)...), Components(c...))
	}
}

func ListItem(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.ListItem(Components(c...))
}
