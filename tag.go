package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Tag(css vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Span(
			vecty.Markup(vecty.Class("tag"), css),
			vecty.Markup(applyer...),
			Components(c...))
	}
}

func Tags(css vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Span(
			vecty.Markup(vecty.Class("tags"), css),
			vecty.Markup(applyer...),
			Components(c...))
	}
}
