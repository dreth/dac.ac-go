package components

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func HTMLProps() attrs.Props {
	return attrs.Props{
		attrs.Class: "bg-background text-text fonts",
	}
}

func HTML(Title string, Description string, lang string, Article bool, content *elem.Element) *elem.Element {
	return elem.Html(
		HTMLProps(),
		Head(Title, Description, Article),
		Body(content, lang),
	)
}

func HTMLForTailwind() *elem.Element {
	return elem.Html(
		HTMLProps(),
		Head("Daniel Alonso", "Daniel Alonso's website", false),
		Body(Content("/cv"), "en"),
	)
}
