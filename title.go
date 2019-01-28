package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Title(css vecty.ClassMap, value string) vecty.ComponentOrHTML {
	return elem.Paragraph(vecty.Markup(vecty.Class("title"),
		css), vecty.Text(value))
}

func Subtitle(css vecty.ClassMap, value string) vecty.ComponentOrHTML {
	return elem.Paragraph(vecty.Markup(vecty.Class("subtitle"),
		css), vecty.Text(value))
}
