package bulma

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type Progress struct {
	vecty.Core

	Color         string
	Size          string
	Indeterminate bool
	Max           string
	Value         string
}

func (t *Progress) Render() vecty.ComponentOrHTML {
	return elem.Progress(vecty.Markup(
		vecty.Class("progress"),
		vecty.MarkupIf(!t.Indeterminate && t.Value != "", prop.Value(t.Value)),
		vecty.MarkupIf(t.Max != "", vecty.Attribute("max", t.Max)),
		ClassMap(t.Color, t.Size)),
		vecty.Text(fmt.Sprintf("%s%", t.Value)))
}
