package bulma

import (
	"fmt"
	"github.com/gopherjs/jquery"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/kooksee/bulma/vars"
	"reflect"
	"sort"
	"strings"
)

var jQuery = jquery.NewJQuery

type TableComponent struct {
	vecty.Core
	SelectedNum int
	Header      [][]string

	Body []map[string]interface{}

	_headerSorted []string
	header        vecty.ComponentOrHTML
	footer        vecty.ComponentOrHTML

	Markup vecty.MarkupList
	Slot   vecty.List

	isSorted       bool
	onHandleHeader func(hs []string) vecty.ComponentOrHTML
	onHandleBody   func(k string, v interface{}) vecty.ComponentOrHTML
}

func (t *TableComponent) HandleHeader(fn func(hs []string) vecty.ComponentOrHTML) *TableComponent {
	t.onHandleHeader = fn
	return t
}

func (t *TableComponent) HandleBody(fn func(k string, v interface{}) vecty.ComponentOrHTML) *TableComponent {
	t.onHandleBody = fn
	return t
}

func (t *TableComponent) Render() vecty.ComponentOrHTML {
	if t.onHandleHeader != nil {
		var _hs vecty.List
		for _, hs := range t.Header {
			_hs = append(_hs, t.onHandleHeader(hs))
		}

		if len(_hs) > 0 {
			t.header = elem.TableHead(t.Markup, elem.TableRow(vecty.Markup(event.Click(func(i *vecty.Event) {
				t.isSorted = !t.isSorted

				if i.Target.Get("tagName").String() == "ABBR" {
					name := i.Target.Get("title").String()
					sort.Slice(t.Body, func(i, j int) bool {
						a := t.Body[i][name]
						b := t.Body[j][name]

						if a == nil {
							return true
						}

						if b == nil {
							return true
						}

						if reflect.TypeOf(a).Kind() == reflect.String {
							if t.isSorted {
								return strings.Compare(a.(string), b.(string)) > 0
							} else {
								return strings.Compare(a.(string), b.(string)) < 0
							}
						}

						if reflect.TypeOf(a).Kind() == reflect.Int {
							if t.isSorted {
								return a.(int) > b.(int)
							} else {
								return a.(int) < b.(int)
							}
						}

						return true
					})

					vecty.Rerender(t)
				}
			})), _hs))
		}
	}

	if t.onHandleBody != nil {
		var _bd vecty.List
		fmt.Println(t.Body)

		for _i, _tr := range t.Body {

			var _td vecty.List
			for k, v := range _tr {
				_td = append(_td, t.onHandleBody(k, v))
			}

			_bd = append(_bd, elem.TableRow(vecty.Markup(vecty.Data("k_num", fmt.Sprintf("%d", _i)), vecty.ClassMap{vars.IsSelected: t.SelectedNum == _i}, event.Click(func(i *vecty.Event) {
				t.SelectedNum = i.Target.Get("parentNode").Get("dataset").Get("k_num").Int()
				vecty.Rerender(t)
			})), _td))

			fmt.Println(_tr)
		}

		return elem.Table(vecty.Markup(vecty.Class("table")), t.Markup, t.header, _bd, t.Slot)
	}

	return nil
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

	IsSelected bool
	OnClick    func(int)
	Num        int
	Markup     vecty.MarkupList
	Slot       vecty.List
}

func (t *TRComponent) Render() vecty.ComponentOrHTML {
	return elem.TableRow(vecty.Markup(vecty.ClassMap{vars.IsSelected: t.IsSelected}, event.Click(func(i *vecty.Event) {
		if t.OnClick != nil {
			t.OnClick(t.Num)
		}
	})), t.Markup, t.Slot)
}

type AbbrComponent struct {
	vecty.Core

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *AbbrComponent) Render() vecty.ComponentOrHTML {
	return elem.Abbreviation(t.Markup, t.Slot)
}
