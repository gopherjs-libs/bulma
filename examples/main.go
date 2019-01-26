package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	. "github.com/kooksee/bulma"
)

// MainView is the parent level view.
type MainView struct {
	vecty.Core
}

func (pv *MainView) Render() vecty.ComponentOrHTML {
	_bsy := ButtonModel.Style
	_bc := ButtonModel.Color

	return elem.Body(
		SimpleBox()(article()),
		SimpleButtons(ButtonModel.Size.AreSmall)(
			SimpleButton()(vecty.Text("SimpleButton")),
			SimpleButton(ButtonModel.Color.IsWhite)(vecty.Text("SimpleButton")),
			SimpleButton(ButtonModel.Color.IsDanger)(vecty.Text("SimpleButton")),
			SimpleButton(ButtonModel.Color.IsDark)(vecty.Text("SimpleButton")),
			SimpleButton(ButtonModel.Color.IsPrimary)(vecty.Text("SimpleButton"))),

		SimpleButtons(ButtonModel.Size.AreLarge)(
			SimpleButton()(vecty.Text("SimpleButton")),
			SimpleButton(ButtonModel.Color.IsWhite)(vecty.Text("SimpleButton")),
			SimpleButton(ButtonModel.Color.IsDanger)(vecty.Text("SimpleButton")),
			SimpleButton(ButtonModel.Color.IsDark)(vecty.Text("SimpleButton")),
			SimpleButton(ButtonModel.Color.IsPrimary)(vecty.Text("SimpleButton"))),

		SimpleButton(_bc.IsInfo, _bsy.IsInverted)(
			vecty.Text("SimpleButton")),

		SimpleButton(_bsy.IsRounded, _bc.IsDanger)(
			vecty.Text("SimpleButton")),

		GroupButton(
			SimpleButton(_bsy.IsRounded, _bc.IsDanger)(
				vecty.Text("SimpleButton")),
			SimpleButton(_bsy.IsRounded, _bc.IsDark)(
				vecty.Text("SimpleButton")),
		),
	)
}

func main() {
	vecty.SetTitle("test bulma")
	vecty.RenderBody(&MainView{})
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
