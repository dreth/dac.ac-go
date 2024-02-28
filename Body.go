package components

import (
	"fmt"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
)

func Content(defaultHxGet string) *elem.Element {
	return elem.Div(
		attrs.Props{
			attrs.ID:       "content",
			attrs.Class:    Rt("content flex flex-col justify-center text-justify text-text max-w-4xl xl:max-w-6xl mx-auto px-4 pt-14 pb-6"),
			htmx.HXGet:     fmt.Sprintf("%s/content", defaultHxGet),
			htmx.HXTrigger: "load",
		},
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
				attrs.Src:   "https://sa.dac.ac/latest.js",
			},
		),
		elem.NoScript(
			nil,
			elem.Img(
				attrs.Props{
					attrs.Src:            "https://sa.dac.ac/noscript.gif",
					attrs.Alt:            "",
					attrs.Referrerpolicy: "no-referrer-when-downgrade",
				},
			),
		),
	)
}
