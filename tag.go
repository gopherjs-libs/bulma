package bulma

import (
	"github.com/gopherjs/vecty"
)

func Tag(css vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &TagComponent{css: css, Markup: vecty.Markup(applyer...), Slot: vecty.List(c)}
	}
}

func Tags(css vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &TagsComponent{css: css, Markup: vecty.Markup(applyer...), Slot: vecty.List(c)}
	}
}
