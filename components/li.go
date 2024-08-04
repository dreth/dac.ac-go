package components

import (
	"github.com/chasefleming/elem-go"
)

func LiSpan(text elem.TextNode) *elem.Element {
	return elem.Li(
		nil,
		elem.Span(nil, text),
	)
}
