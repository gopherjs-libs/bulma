package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type Button struct {
	vecty.Core

	Size    []string
	Color   string
	State   string
	Style   []string
	Slot    vecty.List
	OnClick EventHandle
	Markup  vecty.MarkupList
}

func (t *Button) Render() vecty.ComponentOrHTML {
	t.Style = append(t.Style, t.Size...)
	t.Style = append(t.Style, t.Color, t.State)

	return elem.Anchor(vecty.Markup(vecty.Class("button"),
		ClassMap(t.Style...),
		vecty.MarkupIf(t.OnClick != nil, event.Click(t.OnClick))),
		t.Slot,
		t.Markup)
}

type Buttons struct {
	vecty.Core
	Size string

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *Buttons) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("buttons"),
		ClassMap(t.Size)),
		t.Markup,
		t.Slot)
}

func SimpleButtons(style ...string) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Div(vecty.Markup(vecty.Class("buttons"),
			ClassMap(style...)),
			Components(c...))
	}
}

func SimpleButton(style ...string) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Anchor(vecty.Markup(vecty.Class("button"),
			ClassMap(style...)),
			Components(c...),
		)
	}
}

func AttachButtons(alignment ...string) func(ps ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(ps ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &Buttons{
			Slot:   Components(ps...),
			Markup: vecty.Markup(ClassMap(append(alignment, "has-addons")...)),
		}
	}
}

func GroupButton(ps ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("field", "is-grouped"),
	), MapElem(ps, func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Paragraph(vecty.Markup(vecty.Class("control")), c)
	}))
}

func AddonsButtons(ps ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("field", "has-addons"),
	), MapElem(ps, func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Paragraph(vecty.Markup(vecty.Class("control")), c)
	}))
}
