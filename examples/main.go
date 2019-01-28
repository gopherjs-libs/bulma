package main

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	. "github.com/kooksee/bulma"
	"github.com/kooksee/bulma/attrs"
	"github.com/kooksee/bulma/vars"
)

func main() {
	vecty.SetTitle("test bulma")
	vecty.RenderBody(&MainView{
		isShow:    true,
		isTagShow: true,
	})
}

// MainView is the parent level view.
type MainView struct {
	vecty.Core
	isShow    bool
	isTagShow bool
}

func (pv *MainView) Render() vecty.ComponentOrHTML {

	return elem.Body(
		Box()(
			article()),

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
			SButton()(vecty.Text("SButton")),
			SButton(vars.IsWhite)(vecty.Text("SButton")),
			SButton(vars.IsDanger)(vecty.Text("SButton")),
			SButton(vars.IsDark)(vecty.Text("SButton")),
			SButton(vars.IsPrimary)(vecty.Text("SButton"))),

		SButton(vars.IsInfo, vars.IsInverted)(
			vecty.Text("SButton")),

		SButton(vars.IsRounded, vars.IsDanger)(
			vecty.Text("SButton")),

		GroupButton(
			SButton(vars.IsRounded, vars.IsDanger)(
				vecty.Text("SButton")),
			SButton(vars.IsRounded, vars.IsDark)(
				vecty.Text("SButton")),
		),

		AddonsButton(
			Btn(Css(vars.IsLoading))(
				vecty.Text("SButton")),
			Btn(Css(vars.IsRounded, vars.IsDark, vars.IsStatic))(
				vecty.Text("SButton"))),

		AddonsButton(
			IconButton(Css())(
				Icon(Css(vars.IsSmall))(I("fa-bold")),
				"Bold"),
			IconButton(Css())(
				Icon(Css(vars.IsSmall))(I("fa-italic")),
				"Italic"),
			IconButton(Css())(
				Icon(Css(vars.IsSmall))(I("fa-underline")),
				"Underline"),
		),

		AlignmentButton(Css(vars.IsCentered))(
			IconButton(Css())(
				Icon(Css(vars.IsSmall))(I("fa-bold")),
				"Bold"),
			IconButton(Css())(
				Icon(Css(vars.IsSmall))(I("fa-italic")),
				"Italic"),
			IconButton(Css(vars.IsDanger, vars.IsSelected))(
				Icon(Css(vars.IsSmall))(I("fa-underline")),
				"Underline"),
		),

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

		//Image(Css("is-128x128"))("https://bulma.io/images/placeholders/128x128.png"),
		//RoundedImage(Css("is-128x128"))("https://bulma.io/images/placeholders/128x128.png"),

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
			TBody(TR(
				Th(Text("1")),
				Td(Text("2")),
				Td(Text("3")),
			), TR(
				Th(Text("1")),
				Td(Text("2")),
				Td(Text("3")),
			), TR(
				Th(Text("1")),
				Td(Text("2")),
				Td(A(Css(),
					attrs.Href("https://en.wikipedia.org/wiki/Manchester_United_F.C."),
					attrs.Title("Manchester United F.C."),
				)(vecty.Text("test"))),
			)),
		),
	)
}

func article() vecty.ComponentOrHTML {
	return elem.Article(
		vecty.Markup(
			vecty.Class("media"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("media-left"),
			),
			elem.Figure(
				vecty.Markup(
					vecty.Class("image", "is-64x64"),
				),
				elem.Image(
					vecty.Markup(
						prop.Src("https://bulma.io/images/placeholders/128x128.png"),
						vecty.Attribute("alt", "Image"),
					),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("media-content"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("content"),
				),
				elem.Paragraph(
					elem.Strong(
						vecty.Text("John Smith"),
					),
					elem.Small(
						vecty.Text("@johnsmith"),
					),
					elem.Small(
						vecty.Text("31m"),
					),
					elem.Break(),
					vecty.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis."),
				),
			),
			elem.Navigation(
				vecty.Markup(
					vecty.Class("level", "is-mobile"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("level-left"),
					),
					elem.Anchor(
						vecty.Markup(
							vecty.Class("level-item"),
							vecty.Attribute("aria-label", "reply"),
						),
						elem.Span(
							vecty.Markup(
								vecty.Class("icon", "is-small"),
							),
							elem.Italic(
								vecty.Markup(
									vecty.Class("fas", "fa-reply"),
									vecty.Attribute("aria-hidden", "true"),
								),
							),
						),
					),
					elem.Anchor(
						vecty.Markup(
							vecty.Class("level-item"),
							vecty.Attribute("aria-label", "retweet"),
						),
						elem.Span(
							vecty.Markup(
								vecty.Class("icon", "is-small"),
							),
							elem.Italic(
								vecty.Markup(
									vecty.Class("fas", "fa-retweet"),
									vecty.Attribute("aria-hidden", "true"),
								),
							),
						),
					),
					elem.Anchor(
						vecty.Markup(
							vecty.Class("level-item"),
							vecty.Attribute("aria-label", "like"),
						),
						elem.Span(
							vecty.Markup(
								vecty.Class("icon", "is-small"),
							),
							elem.Italic(
								vecty.Markup(
									vecty.Class("fas", "fa-heart"),
									vecty.Attribute("aria-hidden", "true"),
								),
							),
						),
					),
				),
			),
		),
	)
}
