package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func StackedIcon(css vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Span(vecty.Markup(vecty.Class("fa-stack"),
			vecty.Markup(applyer...)),
			Components(c...))
	}
}

func I(style ...string) vecty.ComponentOrHTML {
	return elem.Span(vecty.Markup(vecty.Class("fas"), ClassMap(style...)))
}

func Icon(css vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Span(vecty.Markup(vecty.Class("icon"), css),
			vecty.Markup(applyer...),
			Components(c...))
	}
}
