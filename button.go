package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Button struct {
	vecty.Core

	Size  string
	State string
	Color string

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *Button) Render() vecty.ComponentOrHTML {
	_css := vecty.ClassMap{}
	if t.Size != "" {
		_css[t.Size] = true
	}

	if t.State != "" {
		_css[t.State] = true
	}

	if t.Color != "" {
		_css[t.Color] = true
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("button"),
			_css,
		),
		t.Markup, t.Slot,
	)
}

type GroupButton struct {
	vecty.Core

	Size  string
	State string
	Color string

	Markup vecty.MarkupList
	Slot   []*Button
}

func (t *GroupButton) Render() vecty.ComponentOrHTML {

	var _ss vecty.List
	for _, s := range t.Slot {
		_ss = append(_ss, elem.Paragraph(
			vecty.Markup(
				vecty.Class("control"),
			), s,
		))
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("field"),
			ClassMap(t.Size, t.State, t.Color),
		),
		_ss,
	)
}
