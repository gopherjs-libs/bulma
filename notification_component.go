package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type NotificationComponent struct {
	vecty.Core

	Css    vecty.ClassMap
	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *NotificationComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("notification"), t.Css), Del(t.Markup), t.Slot)
}
