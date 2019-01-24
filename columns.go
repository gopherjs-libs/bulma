package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Columns struct {
	vecty.Core
}

func (*Columns) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("columns"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("column"),
			),
			vecty.Text("First column"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("column"),
			),
			vecty.Text("Second column"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("column"),
			),
			vecty.Text("Third column"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("column"),
			),
			vecty.Text("Fourth column"),
		),
	)

}

type columnSizes struct {
}

//var ColumnSizes

type Column struct {
	vecty.Core

	columnSizes
}

func (*Column) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("column", "is-two-thirds"),
		),
		vecty.Text("is-two-thirds"),
	)
}
