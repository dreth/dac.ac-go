package components

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func Content(content *elem.Element) *elem.Element {
	return elem.Div(
		attrs.Props{
			attrs.ID:    "content",
			attrs.Class: Rt("content flex flex-col justify-center text-justify text-text max-w-4xl xl:max-w-6xl mx-auto px-4 pt-14 pb-6"),
		},
		content,
	)
}

func ContentArticle(content *elem.Element) *elem.Element {
	return elem.Div(
		attrs.Props{
			attrs.ID:    "content",
			attrs.Class: Rt("content flex flex-col justify-center text-justify text-text max-w-4xl xl:max-w-6xl mx-auto px-4 pt-14 pb-6"),
		},
		content,
	)
}

func Body(content *elem.Element, lang string) *elem.Element {
	return elem.Body(
		nil,
		Navbar(lang),
		content,
		Footer(lang),
		Analytics(),
	)
}

func Analytics() *elem.Element {
	return elem.Div(
		nil,
		elem.Script(
			attrs.Props{
				attrs.Async: "true",
				attrs.Defer: "true",
				attrs.Src:   "https://sa.dac.sg/latest.js",
			},
		),
		elem.NoScript(
			nil,
			elem.Img(
				attrs.Props{
					attrs.Src:            "https://sa.dac.sg/noscript.gif",
					attrs.Alt:            "",
					attrs.Referrerpolicy: "no-referrer-when-downgrade",
				},
			),
		),
	)
}
