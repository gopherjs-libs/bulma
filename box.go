package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func SimpleBox(style ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Div(vecty.Markup(vecty.Class("box")),
			vecty.Markup(style...),
			Components(c...))
	}
}

type Box struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *Box) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("box")),
		t.Markup,
		t.Slot)
}
