package main

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	. "github.com/kooksee/bulma"
	"github.com/kooksee/bulma/attrs"
	"github.com/kooksee/bulma/vars"
)

func main() {
	vecty.SetTitle("test bulma")
	vecty.RenderBody(&MainView{
		isShow:    true,
		isTagShow: true,
		isDown:    false,
	})
}

// MainView is the parent level view.
type MainView struct {
	vecty.Core
	isShow    bool
	isTagShow bool
	isDown    bool
}

func (pv *MainView) Mount() {
	fmt.Println("Mount")
}

func (pv *MainView) Render() vecty.ComponentOrHTML {
	fmt.Println("111111")
	return elem.Body(
		Box()(Text("SButton")),


		SButtons(vars.AreSmall)(
			SButton()(
				Text("SButton")),
			SButton(vars.IsSuccess)(
				Text("SButton")),
			SButton(vars.IsDanger)(
				Text("SButton")),
			SButton(vars.IsDark, vars.IsLarge)(
				Text("SButton")),
			SButton(vars.IsPrimary)(
				Text("SButton"))),

		SButtons(vars.AreLarge)(
			SButton()(Text("SButton")),
			SButton(vars.IsWhite)(Text("SButton")),
			SButton(vars.IsDanger)(Text("SButton")),
			SButton(vars.IsDark)(Text("SButton")),
			SButton(vars.IsPrimary)(Text("SButton"))),

		SButton(vars.IsInfo, vars.IsInverted)(
			Text("SButton")),

		SButton(vars.IsRounded, vars.IsDanger)(
			Text("SButton")),

		GroupButton(
			SButton(vars.IsRounded, vars.IsDanger)(
				Text("SButton")),
			SButton(vars.IsRounded, vars.IsDark)(
				Text("SButton"))),

		AddonsButton(
			Btn(Css(vars.IsLoading))(
				Text("SButton")),
			Btn(Css(vars.IsRounded, vars.IsDark, vars.IsStatic))(
				Text("SButton"))),

		Icon(Css(vars.IsSmall))(I("fa-bold")),

		IconButton(Css())(
			Icon(Css(vars.IsSmall))(I("fa-bold")),
			Text("Bold")),

		AddonsButton(
			IconButton(Css())(
				Icon(Css(vars.IsSmall))(I("fa-bold")),
				Text("Bold")),
			IconButton(Css())(
				Icon(Css(vars.IsSmall))(I("fa-italic")),
				Text("Italic")),
			IconButton(Css())(
				Icon(Css(vars.IsSmall))(I("fa-underline")),
				Text("Underline")),
		),

		AlignmentButton(Css(vars.IsCentered))(
			IconButton(Css())(
				Icon(Css(vars.IsSmall))(I("fa-bold")),
				Text("Bold")),
			IconButton(Css())(
				Icon(Css(vars.IsSmall))(I("fa-italic")),
				Text("Italic")),
			IconButton(Css(vars.IsDanger, vars.IsSelected))(
				Icon(Css(vars.IsSmall))(I("fa-underline")),
				Text("Underline"))),

		Cnt(Css())(
			OrderedList(Css())(
				ListItem(Text("hello1")),
				ListItem(Text("hello1")),
				ListItem(Text("hello1")),
				ListItem(Text("hello1")),
			),

			OrderedList(Css(vars.IsLowerAlpha))(
				ListItem(Text("hello1")),
				ListItem(Text("hello1")),
				ListItem(Text("hello1")),
				ListItem(Text("hello1")),
			),

			UnorderedList(Css())(
				ListItem(Text("hello1")),
				ListItem(Text("hello1")),
				ListItem(Text("hello1")),
				ListItem(Text("hello1")),
			),

			Del(),
		),

		Image(Css("is-128x128"))("https://bulma.io/images/placeholders/128x128.png"),
		RoundedImage(Css("is-128x128"))("https://bulma.io/images/placeholders/128x128.png"),

		vecty.If(pv.isShow, Notification(Css(vars.IsInfo), event.Click(func(i *vecty.Event) {
			pv.isShow = false
			fmt.Println("load111")
			vecty.Rerender(pv)
		}))(Text("Primar lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor."))),

		Progress(Css(vars.IsSmall), "30", "100", false),
		Progress(Css(vars.IsInfo), "30", "100", true),

		Tag(Css())(Text("ss")),
		Tags(Css())(
			Tag(Css())(Text("Tag")),
			Tag(Css(vars.IsRounded))(Text("Tag")),
			If(pv.isTagShow, Tag(Css(vars.IsSuccess))(Text("Tag"), Del(event.Click(func(i *vecty.Event) {
				pv.isTagShow = false
				vecty.Rerender(pv)
			})))),
		),

		If(pv.isTagShow, Tags(Css(vars.HasAddons))(
			Tag(Css(vars.IsDanger))(Text("Tag")),
			Tag(Css(vars.IsRounded, vars.IsDelete), event.Click(func(i *vecty.Event) {
				pv.isTagShow = false
				vecty.Rerender(pv)
			}))(),
		)),

		MultilineGrouped(If(pv.isTagShow, Tags(Css(vars.HasAddons))(
			Tag(Css(vars.IsDanger))(Text("Tag")),
			Tag(Css(vars.IsRounded, vars.IsDelete), event.Click(func(i *vecty.Event) {
				pv.isTagShow = false
				vecty.Rerender(pv)
			}))(),
		)), If(pv.isTagShow, Tags(Css(vars.HasAddons))(
			Tag(Css(vars.IsDanger))(Text("Tag")),
			Tag(Css(vars.IsRounded, vars.IsDelete), event.Click(func(i *vecty.Event) {
				pv.isTagShow = false
				vecty.Rerender(pv)
			}))(),
		)), If(pv.isTagShow, Tags(Css(vars.HasAddons))(
			Tag(Css(vars.IsDanger))(Text("Tag")),
			Tag(Css(vars.IsRounded, vars.IsDelete), event.Click(func(i *vecty.Event) {
				pv.isTagShow = false
				vecty.Rerender(pv)
			}))(),
		))),

		MultilineGrouped(
			ListItem(Text("hello1")),
			ListItem(Text("hello1")),
			ListItem(Text("hello1")),
			ListItem(Text("hello1"))),

		Table(Css("is-fullwidth"))(
			THead(Th(Abbr("Position")), Th(Abbr("Played")), Th(Abbr("Team"))),
			TFoot(Th(Abbr("Position")), Th(Abbr("Played")), Th(Abbr("Team"))),
			TBody(TR(Css())(
				Th(Text("1")),
				Td(Text("2")),
				Td(Text("3")),
			), TR(Css(vars.IsSelected))(
				Th(Text("1")),
				Td(Text("2")),
				Td(Text("3")),
			), TR(Css())(
				Th(Text("1")),
				Td(Text("2")),
				Td(A(Css(),
					attrs.Href("https://en.wikipedia.org/wiki/Manchester_United_F.C."),
					attrs.Title("Manchester United F.C."),
				)(Text("test"))),
			)),
		),

		elem.Button(vecty.Markup(event.Click(func(i *vecty.Event) {
			fmt.Println("Button Click")
			vecty.Rerender(pv)
		}), vecty.Class(vars.IsInfo, "button")), vecty.Text("测试")),
	)
}
