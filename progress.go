package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Progress(css vecty.ClassMap, value string, max string, indeterminate bool) vecty.ComponentOrHTML {
	return elem.Progress(vecty.Markup(vecty.Class("progress"),
		vecty.MarkupIf(!indeterminate, vecty.Attribute("value", value)),
		vecty.MarkupIf(max != "", vecty.Attribute("max", max)),
		css), vecty.Text(value+"%"))

}
