package bulma

import (
	"github.com/gopherjs/vecty"
)

type Table TableComponent

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

func Abbr(title []string, c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	_t := ""
	_v := ""
	if len(title) > 0 {
		_v = title[0]
	}

	if len(title) > 1 {
		_t = title[1]
	}

	return &AbbrComponent{Markup: vecty.Markup(vecty.MarkupIf(_t != "", vecty.Attribute("title", title))), Slot: append(c, Text(_v))}
}
