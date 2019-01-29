package bulma

import (
	"github.com/gopherjs/vecty"
)

func Box(style ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &BoxComponent{Markup: vecty.Markup(style...), Slot: vecty.List(c)}
	}
}
