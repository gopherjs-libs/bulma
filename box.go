package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Box struct {
	vecty.Core
	Slot vecty.List
}

func (t *Box) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(
		vecty.Class("box"),
	), t.Slot)
}
