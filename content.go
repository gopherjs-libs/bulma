package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Content(applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Div(vecty.Markup(append(applyer, ClassMap("content"))...),
			Components(c...))
	}
}

func Cnt(l vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return Content(vecty.Markup(append(applyer, l)...))
}

func Block(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("block"),
		),
		Components(c...),
	)
}
