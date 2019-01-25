package bulma

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

var ImageSize = struct {
	Is16x16   string
	Is24x24   string
	Is32x32   string
	Is48x48   string
	Is64x64   string
	Is96x96   string
	Is128x128 string
	IsSquare  string
	Is1by1    string
	Is5by4    string
	Is4by3    string
	Is3by2    string
	Is5by3    string
	Is16by9   string
	Is2by1    string
	Is3by1    string
	Is4by5    string
	Is3by4    string
	Is2by3    string
	Is3by5    string
	Is9by16   string
	Is1by2    string
	Is1by3    string
}{
	Is16x16:   "is-16x16",
	Is24x24:   "is-24x24",
	Is32x32:   "is-32x32",
	Is48x48:   "is-48x48",
	Is64x64:   "is-64x64",
	Is96x96:   "is-96x96",
	Is128x128: "is-128x128",
	IsSquare:  "is-square",
	Is1by1:    "is-1by1",
	Is5by4:    "is-5by4",
	Is4by3:    "is-4by3",
	Is3by2:    "is-3by2",
	Is5by3:    "is-5by3",
	Is16by9:   "is-16by9",
	Is2by1:    "is-2by1",
	Is3by1:    "is-3by1",
	Is4by5:    "is-4by5",
	Is3by4:    "is-3by4",
	Is2by3:    "is-2by3",
	Is3by5:    "is-3by5",
	Is9by16:   "is-9by16",
	Is1by2:    "is-1by2",
	Is1by3:    "is-1by3",
}

type Image struct {
	vecty.Core

	Url       string
	Size      string
	IsRounded bool

	Markup vecty.MarkupList
	Slot   vecty.List
}

func (t *Image) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("image"),
			ClassMap(t.Size),
		),
		elem.Image(
			vecty.Markup(
				vecty.ClassMap{
					"is-rounded": t.IsRounded,
				},
				prop.Src(t.Url),
			),
		), t.Markup, t.Slot)
}
