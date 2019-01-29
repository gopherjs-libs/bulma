package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Card(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("card")), vecty.List(c))
}

func CardHeader(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Header(vecty.Markup(vecty.Class("card-header")), vecty.List(c))
}

func CardHeaderTitle(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Paragraph(vecty.Markup(vecty.Class("card-header-title")), vecty.List(c))
}

func CardHeaderIcon(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Anchor(vecty.Markup(vecty.Class("card-header-icon")), vecty.List(c))
}

func CardContent(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("card-content")),
		vecty.List(c),
	)
}

func CardFooter(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Footer(vecty.Markup(vecty.Class("card-footer")), vecty.List(c))
}

func CardFooterItem(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Paragraph(vecty.Markup(vecty.Class("card-footer-item")), vecty.List(c))
}
