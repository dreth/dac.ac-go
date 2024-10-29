package components

import (
	"dac.sg/constants"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func Head(Title string, Description string, Article bool) *elem.Element {
	// all main nodes
	var headNodes []elem.Node

	// append the basic nodes
	headNodes = append(headNodes,
		elem.Meta(attrs.Props{
			attrs.Name:    "viewport",
			attrs.Content: "width=device-width, initial-scale=1.0",
		}),
		elem.Meta(attrs.Props{
			attrs.Name:    "description",
			attrs.Content: Description,
		}),
		elem.Meta(attrs.Props{
			attrs.Name:    "title",
			attrs.Content: Title,
		}),
		elem.Meta(attrs.Props{
			attrs.Name:    "charset",
			attrs.Content: "UTF-8",
		}),
		elem.Meta(attrs.Props{
			attrs.HTTPequiv: "Content-Type",
			attrs.Content:   "text/html; charset=UTF-8",
		}),

		// Open Graph / Facebook
		elem.Meta(attrs.Props{
			"property":    "og:type",
			attrs.Content: "website",
		}),
		elem.Meta(attrs.Props{
			"property":    "og:url",
			attrs.Content: "https://dac.sg",
		}),
		elem.Meta(attrs.Props{
			"property":    "og:title",
			attrs.Content: Title,
		}),
		elem.Meta(attrs.Props{
			"property":    "og:description",
			attrs.Content: Description,
		}),
		elem.Meta(attrs.Props{
			"property":    "og:image",
			attrs.Content: constants.ResourceWithCDN("/assets/icons/meta.png"),
		}),

		// Twitter / X
		elem.Meta(attrs.Props{
			"property":    "twitter:card",
			attrs.Content: constants.ResourceWithCDN("/assets/icons/meta.png"),
		}),
		elem.Meta(attrs.Props{
			"property":    "twitter:url",
			attrs.Content: "https://dac.sg",
		}),
		elem.Meta(attrs.Props{
			"property":    "twitter:title",
			attrs.Content: Title,
		}),
		elem.Meta(attrs.Props{
			"property":    "twitter:description",
			attrs.Content: Description,
		}),
		elem.Meta(attrs.Props{
			"property":    "twitter:image",
			attrs.Content: constants.ResourceWithCDN("/assets/icons/meta.png"),
		}),

		// Site font (ubuntu)
		elem.Link(attrs.Props{
			attrs.Rel:  "stylesheet",
			attrs.Href: "https://fonts.googleapis.com/css?family=Ubuntu",
		}),

		// Stylesheets
		elem.Link(attrs.Props{
			attrs.Rel:  "stylesheet",
			attrs.Href: "/assets/tailwind/styles.css",
		}),
		elem.Link(attrs.Props{
			attrs.Rel:  "stylesheet",
			attrs.Href: "/assets/tailwind/custom.css",
		}),

		// Favicon
		elem.Link(attrs.Props{
			attrs.Rel:  "icon",
			attrs.Type: "image/x-icon",
			attrs.Href: constants.ResourceWithCDN("/favicon.ico"),
		}),

		// HTMX
		elem.Script(attrs.Props{
			attrs.Src: constants.ResourceWithCDN("/assets/js/libs/htmx.min.js"),
		}),

		// site title
		elem.Title(nil, elem.Text(Title)),
	)

	if Article {
		// If it's an article, make sure syntax highlighting is enabled
		headNodes = append(headNodes,
			elem.Link(attrs.Props{
				attrs.Rel:  "stylesheet",
				attrs.Href: constants.ResourceWithCDN("/assets/css/libs/sunburst-prettify.css"),
			}),
			elem.Script(attrs.Props{
				attrs.Src: constants.ResourceWithCDN("/assets/js/libs/prettify.js"),
			}),
			// and latex related stuff too
			elem.Link(
				attrs.Props{
					attrs.Rel:         "stylesheet",
					attrs.Href:        constants.ResourceWithCDN("/assets/css/libs/katex.min.css"),
					attrs.Crossorigin: "anonymous",
				},
			),
			elem.Script(
				attrs.Props{
					attrs.Defer:       "true",
					attrs.Src:         constants.ResourceWithCDN("/assets/js/libs/katex.min.js"),
					attrs.Crossorigin: "anonymous",
				},
			),
			elem.Script(
				attrs.Props{
					attrs.Defer: "true",
					attrs.Src:   constants.ResourceWithCDN("/assets/js/katexRenderSingleDollarSign.js"),
				},
			),
			elem.Script(
				attrs.Props{
					attrs.Defer: "true",
					attrs.Src:   constants.ResourceWithCDN("/assets/js/libs/auto-render.min.js"),
					"onload":    "renderMathInElement(document.body);",
				},
			),
		)
	}

	// Basic HTML5 tags
	return elem.Head(nil,
		headNodes...,
	)
}
