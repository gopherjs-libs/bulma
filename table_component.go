package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type TableComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *TableComponent) Render() vecty.ComponentOrHTML {
	return elem.Table(vecty.Markup(vecty.Class("table")), t.Markup, t.Slot)
}

type TBodyComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *TBodyComponent) Render() vecty.ComponentOrHTML {
	return elem.TableBody(t.Markup, t.Slot)
}

type THeadComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *THeadComponent) Render() vecty.ComponentOrHTML {
	return elem.TableHead(t.Markup, elem.TableRow(t.Slot))
}

type TFootComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *TFootComponent) Render() vecty.ComponentOrHTML {
	return elem.TableFoot(t.Markup, elem.TableRow(t.Slot))
}

type TdComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *TdComponent) Render() vecty.ComponentOrHTML {
	return elem.TableData(t.Slot)
}

type ThComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *ThComponent) Render() vecty.ComponentOrHTML {
	return elem.TableHeader(t.Slot)
}

type TRComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *TRComponent) Render() vecty.ComponentOrHTML {
	return elem.TableRow(t.Markup, t.Slot)
}

type AbbrComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *AbbrComponent) Render() vecty.ComponentOrHTML {
	return elem.Abbreviation(t.Markup, t.Slot)
}
