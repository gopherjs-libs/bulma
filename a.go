package bulma

import (
	"github.com/gopherjs/vecty"
)

func A(css vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &AComponent{Markup: vecty.Markup(append(applyer, css)...), Slot: c}
	}
}
