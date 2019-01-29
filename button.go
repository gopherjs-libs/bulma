package bulma

import (
	"github.com/gopherjs/vecty"
)

func Buttons(style ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &ButtonsComponent{Markup: vecty.Markup(style...), Slot: vecty.List(c)}
	}
}

func Btns(l vecty.ClassMap, style ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return Buttons(vecty.Markup(append(style, l)...))
}

func SButtons(style ...string) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return Buttons(vecty.Class(style...))
}

func Button(style ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &ButtonComponent{Markup: vecty.Markup(style...), Slot: vecty.List(c)}
	}
}
func Btn(l vecty.ClassMap, style ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return Button(vecty.Markup(append(style, l)...))
}

func SButton(style ...string) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return Button(vecty.Class(style...))
}

func IconButton(l vecty.ClassMap, style ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &IconButtonComponent{Slot: c, Markup: vecty.Markup(append(style, l)...)}
	}
}

func AlignmentButton(l vecty.ClassMap, style ...vecty.Applyer) func(ps ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return Btns(l, vecty.Markup(append(style, vecty.Class("has-addons"))...))
}

func GroupButton(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &GroupButtonComponent{Slot: vecty.List(c)}
}

func AddonsButton(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &AddonsButtonComponent{Slot: vecty.List(c)}
}
