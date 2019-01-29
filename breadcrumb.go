package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Breadcrumb(l vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Navigation(
			vecty.Markup(vecty.Class("breadcrumb"), l,
				vecty.Attribute("aria-label", "breadcrumbs"),
			),
			vecty.Markup(applyer...),
			elem.UnorderedList(vecty.List(c)),
		)
	}
}
