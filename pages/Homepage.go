package pages

import (
	"fmt"
	"strings"

	"dac.ac/components"
	"dac.ac/middleware"
	"dac.ac/structs"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	// return the full page
	return components.PageHTML(c, HomepageContentHTML(middleware.Lang(c)))
}

func HomepageContent(c *fiber.Ctx) error {
	lang := middleware.Lang(c)
	return c.SendString(HomepageContentHTML(lang).Render())
}

func HomepageSocialButton(network string, link string) *elem.Element {
	return elem.A(
		attrs.Props{
			attrs.Href: link,
		},
		elem.Img(
			attrs.Props{
				attrs.Src:   fmt.Sprintf("/assets/icons/%s.svg", strings.ToLower(network)),
				attrs.Alt:   network,
				attrs.Class: components.Rt("w-7 lg:w-9 2xl:w-10 h-7 lg:h-9 2xl:h-10 inline-block align-middle ml-4"),
			},
		),
	)
}

func HomepageContentHTML(lang string) *elem.Element {
	// Get content from YAML
	text := structs.GetLanguagesLoc(lang)

	// Basic props
	// UlProps := attrs.Props{attrs.Class: "list-disc p-4"}
	HeadingProps2 := attrs.Props{attrs.Class: components.RtH("font-black pt-3")}

	// Left Column
	leftColumn := elem.Div(
		attrs.Props{
			attrs.Class: "w-full sm:w-1/2 sm:pr-8 sm:mb-8",
		},
		// About me basic text
		elem.Div(
			attrs.Props{
				attrs.Class: "pb-1 sm:pb-3",
			},
			elem.H2(
				attrs.Props{
					attrs.ID:    "about-me-heading",
					attrs.Class: components.RtH("font-black"),
				},
				elem.Text(text["about_title"]),
			),
			elem.P(
				attrs.Props{
					attrs.Class: "text-left py-2",
				},
				elem.Text(text["about_me_short"]),
			),
		),
		// Contact section
		elem.Div(
			attrs.Props{
				attrs.Class: "py-1 sm:py-3",
			},
			elem.H2(
				HeadingProps2,
				elem.Text(text["contact_title"]),
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "flex flex-row py-2",
				},
				HomepageSocialButton("Email", "/profiles/email"),
				HomepageSocialButton("LinkedIn", "/profiles/linkedin"),
				HomepageSocialButton("GitHub", "/profiles/github"),
				HomepageSocialButton("X", "/profiles/x"),
			),
		),
		// Latest 5 articles I wrote
		elem.Div(
			attrs.Props{
				attrs.Class: "py-1 sm:py-3",
			},
			elem.H2(
				HeadingProps2,
				elem.Text(text["blog_title"]),
			),
			elem.P(
				attrs.Props{
					attrs.Class: "text-justify py-2",
				},
				elem.Text(text["latest_articles_text"]),
			),
			elem.Div(
				nil,
				ArticleList(5)...,
			),
			elem.A(
				attrs.Props{
					attrs.Href:  "/blog",
					attrs.Class: components.RtFooter("text-links hover:text-links-hover"),
				},
				elem.Span(
					nil,
					elem.Text(text["blog_all_articles"]),
				),
			),
		),
	)

	// Right Column
	rightColumn := elem.Div(
		attrs.Props{
			attrs.Class: "w-full sm:w-1/2 mb-4 sm:pl-8",
		},
		// Cool links
		elem.Div(
			attrs.Props{
				attrs.Class: "pb-1 sm:pb-3",
			},
			elem.H2(
				HeadingProps2,
				elem.Text(text["cool_links_title"]),
			),
			elem.P(
				attrs.Props{
					attrs.Class: "text-justify py-2",
				},
				elem.Text(text["latest_links_text"]),
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "pb-1",
				},
				CoolLinksList("wikipedia", 5)...,
			),
			elem.A(
				attrs.Props{
					attrs.Href:  "/cool-links",
					attrs.Class: components.RtFooter("text-links hover:text-links-hover"),
				},
				elem.Span(
					nil,
					elem.Text(text["links_all"]),
				),
			),
		),
		// Featured playlist
		elem.Div(
			attrs.Props{
				attrs.Class: "pb-2 py-1 sm:py-3",
			},
			elem.H2(
				HeadingProps2,
				elem.Text(text["playlists_title"]),
			),
			elem.P(
				attrs.Props{
					attrs.Class: "text-justify py-2",
				},
				elem.Text(text["latest_playlist_text"]),
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "pb-1 flex justify-left items-left",
				},
				PlaylistDiv("ambient"),
			),
			elem.A(
				attrs.Props{
					attrs.Href:  "/playlists",
					attrs.Class: components.RtFooter("text-links hover:text-links-hover"),
				},
				elem.Span(
					nil,
					elem.Text(text["playlists_all"]),
				),
			),
		),
	)

	// Main div
	mainDiv := elem.Div(
		attrs.Props{
			attrs.Class: "container mx-auto text-left pt-3 md:pt-4 xl:pt-5",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "flex flex-wrap",
			},
			leftColumn,
			rightColumn,
		),
	)

	// return
	return mainDiv
}
