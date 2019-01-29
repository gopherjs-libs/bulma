package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Label(c string) vecty.ComponentOrHTML {
	return elem.Label(vecty.Markup(vecty.Class("label")),
		Text(c),
	)
}
