package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Buttons(style ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Div(vecty.Markup(vecty.Class("buttons"),
			vecty.Markup(style...)),
			Components(c...))
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
		return elem.Anchor(vecty.Markup(append(style, vecty.Class("button"))...),
			Components(c...),
		)
	}
}
func Btn(l vecty.ClassMap, style ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return Button(vecty.Markup(append(style, l)...))
}

func SButton(style ...string) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return Button(vecty.Class(style...))
}

func IconButton(l vecty.ClassMap, style ...vecty.Applyer) func(c ...interface{}) vecty.ComponentOrHTML {
	return func(c ...interface{}) vecty.ComponentOrHTML {
		var _p []vecty.ComponentOrHTML
		for _, _c := range c {
			switch _i := _c.(type) {
			case *vecty.HTML:
				_p = append(_p, _i)
			case string:
				_p = append(_p, elem.Span(vecty.Text(_i)))
			}
		}
		return Btn(l, style...)(_p...)
	}
}

func AlignmentButton(l vecty.ClassMap, style ...vecty.Applyer) func(ps ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return Btns(l, vecty.Markup(append(style, vecty.Class("has-addons"))...))
}

func GroupButton(ps ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("field", "is-grouped"),
	), MapElem(ps, func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Paragraph(vecty.Markup(vecty.Class("control")), c)
	}))
}

func AddonsButton(ps ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("field", "has-addons"),
	), MapElem(ps, func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Paragraph(vecty.Markup(vecty.Class("control")), c)
	}))
}
