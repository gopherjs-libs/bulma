package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func Table(l vecty.ClassMap, applyer ...vecty.Applyer) func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return func(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
		return elem.Table(vecty.Markup(vecty.Class("table"), l), vecty.Markup(applyer...), vecty.List(c))
	}
}

func TBody(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.TableBody(vecty.List(c))
}

func THead(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.TableHead(elem.TableRow(vecty.List(c)))
}

func TFoot(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.TableFoot(elem.TableRow(vecty.List(c)))
}

func Td(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.TableData(vecty.List(c))
}

func Th(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.TableHeader(vecty.List(c))
}

func TR(c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.TableRow(vecty.List(c))
}

func Abbr(title string, c ...vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	return elem.Abbreviation(
		vecty.Text(title),
		vecty.List(c),
	)
}
