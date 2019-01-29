package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type TagComponent struct {
	vecty.Core

	css    vecty.ClassMap
	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *TagComponent) Render() vecty.ComponentOrHTML {
	return elem.Span(vecty.Markup(vecty.Class("tag"), t.css), t.Markup, t.Slot)
}

type TagsComponent struct {
	vecty.Core

	css    vecty.ClassMap
	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *TagsComponent) Render() vecty.ComponentOrHTML {
	return elem.Span(vecty.Markup(vecty.Class("tags"), t.css), t.Markup, t.Slot)
}
