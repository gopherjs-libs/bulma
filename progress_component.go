package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type ProgressComponent struct {
	vecty.Core

	css           vecty.ClassMap
	value         string
	max           string
	indeterminate bool
	Markup        vecty.MarkupList
	Slot          vecty.List
}

func (t *ProgressComponent) Render() vecty.ComponentOrHTML {
	return elem.Progress(vecty.Markup(vecty.Class("progress"),
		vecty.MarkupIf(!t.indeterminate, vecty.Attribute("value", t.value)),
		vecty.MarkupIf(t.max != "", vecty.Attribute("max", t.max)),
		t.css), Text(t.value+"%"))
}
