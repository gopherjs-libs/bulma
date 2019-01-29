package bulma

import (
	"github.com/gopherjs/vecty"
)

func Content(applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &ContentComponent{Markup: vecty.Markup(applyer...), Slot: vecty.List(c)}
	}
}

func Cnt(l vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return Content(vecty.Markup(append(applyer, l)...))
}
