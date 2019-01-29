package bulma

import (
	"github.com/gopherjs/vecty"
)

func Table(l vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &TableComponent{Markup: vecty.Markup(append(applyer, l)...), Slot: c}
	}
}

func TBody(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &TBodyComponent{Slot: c}
}

func THead(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &THeadComponent{Slot: c}
}

func TFoot(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &TFootComponent{Slot: c}
}

func Td(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &TdComponent{Slot: c}
}

func Th(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &ThComponent{Slot: c}
}

func TR(l vecty.ClassMap, style ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return &TRComponent{Markup: vecty.Markup(append(style, l)...), Slot: c}
	}
}

func Abbr(title string, c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return &AbbrComponent{Markup: vecty.Markup(vecty.Attribute("title", title)), Slot: append(c, Text(title))}
}
