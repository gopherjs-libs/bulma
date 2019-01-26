package bulma

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)



type Table struct {
	vecty.Core

	HasFooter bool

	Header      []TableHeader
	headerMap   []string
	Body        []MapComponent
	selectedNum int
	Modifier    string
}

func (t *Table) onSelected(rowNum int) {
	t.selectedNum = rowNum
	vecty.Rerender(t)
}

func (t *Table) onSort(title string) {
	fmt.Println(title)
}

func (t *Table) Render() vecty.ComponentOrHTML {
	var _header vecty.List
	for _, _h := range t.Header {
		if _h.Title != "" {
			t.headerMap = append(t.headerMap, _h.Title)
		}

		_header = append(_header, elem.TableHeader(vecty.Markup(event.Click(func(_ *vecty.Event) { t.onSort(_h.Title) })),
			vecty.If(_h.Title != "", elem.Abbreviation(vecty.Markup(vecty.Attribute("title", _h.Title)), vecty.Text(_h.Value))),
			vecty.If(_h.Title == "", vecty.Text(_h.Value))))
	}

	var _body vecty.List
	for i, _h := range t.Body {
		if _h == nil {
			continue
		}

		var _th vecty.List
		for _, _k := range t.headerMap {
			_th = append(_th, elem.TableData(_h[_k]))
		}

		_body = append(_body, &TableRow{Slot: _th, OnClick: t.onSelected, IsSelected: t.selectedNum == i})
	}

	return elem.Table(
		vecty.Markup(
			vecty.Class("table"),
			ClassMap(t.Modifier),
		),
		vecty.If(len(_header) > 0, elem.TableHead(elem.TableRow(_header))),
		vecty.If(t.HasFooter && len(_header) > 0, elem.TableFoot(elem.TableRow(_header))),
		vecty.If(len(_body) > 0, elem.TableBody(_body)),
	)
}

type TableRow struct {
	vecty.Core

	RowNum     int
	IsSelected bool
	Markup     vecty.MarkupList
	Slot       vecty.List
	OnClick    func(int)
}

func (t *TableRow) onClick(e *vecty.Event) {
	if t.OnClick != nil {
		t.OnClick(t.RowNum)
	}
}

func (t *TableRow) Render() vecty.ComponentOrHTML {
	return elem.TableRow(
		vecty.Markup(
			vecty.ClassMap{
				"is-selected": t.IsSelected,
			},
			event.Click(t.onClick),
		),
	)
}

func TH(title, value string) TableHeader {
	return TableHeader{Title: title, Value: value}
}

type TableHeader struct {
	Title string
	Value string
}
