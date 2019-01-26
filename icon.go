package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func StackedIcon(style ...string) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Span(vecty.Markup(vecty.Class("fa-stack"),
			ClassMap(style...)),
			Components(c...))
	}
}

func I(style ...string) vecty.ComponentOrHTML {
	return elem.Span(vecty.Markup(vecty.Class("fas"), ClassMap(style...)))
}

func Icon(style ...string) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Span(vecty.Markup(vecty.Class("icon"),
			ClassMap(style...)),
			Components(c...))
	}
}
