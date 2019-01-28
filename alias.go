package bulma

import "github.com/gopherjs/vecty"

var MarkupIf = vecty.MarkupIf
var Text = vecty.Text

func If(cond bool, children ...vecty.ComponentOrHTML) vecty.List {
	if cond {
		return vecty.List(children)
	}
	return nil
}
