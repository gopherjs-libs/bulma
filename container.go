package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func MultilineGrouped(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("field", "is-grouped", "is-grouped-multiline"),
		),
		MapElem(c, func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
			return elem.Div(vecty.Markup(vecty.Class("control")), c)
		}),
	)
}
