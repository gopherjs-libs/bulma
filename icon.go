package bulma

import (
	"github.com/gopherjs/vecty"
)

func StackedIcon(css vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &StackedIconComponent{Markup: vecty.Markup(append(applyer, css)...), Slot: c}
	}
}

func I(style ...string) vecty.ComponentOrHTML {
	return &IComponent{Markup: vecty.Markup(vecty.Class(style...))}
}

func Icon(css vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &IconComponent{Markup: vecty.Markup(append(applyer, css)...), Slot: c}
	}
}
