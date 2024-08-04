package pages

import (
	"fmt"
	"strings"

	"dac.ac/components"
	"dac.ac/constants"
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
				attrs.Src:   constants.ResourceWithCDN(fmt.Sprintf("/assets/icons/%s.svg", strings.ToLower(network))),
				attrs.Alt:   network,
				attrs.Title: network,
				attrs.Class: components.Rt("w-7 lg:w-9 2xl:w-10 h-7 lg:h-9 2xl:h-10 inline-block align-middle ml-4 transform transition-transform duration-300 hover:scale-110 hover:shadow-lg"),
			},
		),
	)
}

func HomepageContentHTML(lang string) *elem.Element {
	// Get content from YAML
	text := structs.GetLanguagesLoc(lang)

	// Basic props
	// UlProps := attrs.Props{attrs.Class: "list-disc p-4"}
	HeadingProps2 := attrs.Props{attrs.Class: components.RtH("font-black")}

	// Left Column
	leftColumn := elem.Div(
		attrs.Props{
			attrs.Class: "left-column sm:pb-6 pt-4 sm:pl-0 sm:pr-0 space-y-2",
		},
		// About me basic text
		elem.Div(
			attrs.Props{
				attrs.Class: "rounded-lg bg-almost-black-2 p-4",
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
					attrs.Class: "text-left pt-2",
				},
				elem.Text(text["about_me_short"]),
			),
		),
		// Contact section
		elem.Div(
			attrs.Props{
				attrs.Class: "rounded-lg bg-almost-black-2 pt-2 pb-4 p-4",
			},
			elem.H2(
				HeadingProps2,
				elem.Text(text["contact_title"]),
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "flex flex-row pt-2",
				},
				HomepageSocialButton("Email", "/profiles/email"),
				HomepageSocialButton("LinkedIn", "/profiles/linkedin"),
				HomepageSocialButton("GitHub", "/profiles/github"),
			),
		),
		// Latest 5 articles I wrote
		elem.Div(
			attrs.Props{
				attrs.Class: "py-2 rounded-lg bg-almost-black-2 p-4",
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
				attrs.Props{
					attrs.Class: "pb-1 space-y-2",
				},
				ArticleList(6)...,
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
			attrs.Class: "right-column pb-8 sm:pb-6 sm:pt-4  space-y-2",
		},
		// Featured project
		elem.Div(
			attrs.Props{
				attrs.Class: "py-2 rounded-lg bg-almost-black-2 p-4",
			},
			elem.H2(
				HeadingProps2,
				elem.Text(text["projects_title"]),
			),
			elem.P(
				attrs.Props{
					attrs.Class: "text-justify py-2",
				},
				elem.Text(text["latest_project_text"]),
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "pb-1 flex justify-left items-center",
				},
				ProjectList(lang, 1)...,
			),
			elem.A(
				attrs.Props{
					attrs.Href:  "/projects",
					attrs.Class: components.RtFooter("text-links hover:text-links-hover"),
				},
				elem.Span(
					nil,
					elem.Text(text["projects_all"]),
				),
			),
		),
		// Cool links
		elem.Div(
			attrs.Props{
				attrs.Class: "py-2 rounded-lg bg-almost-black-2 p-4",
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
				CoolLinksList("wikipedia", 6)...,
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
				attrs.Class: "py-2 rounded-lg bg-almost-black-2 p-4",
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
					attrs.Class: "pb-1 flex justify-left items-center",
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
			attrs.Class: "container mx-auto text-left grid grid-cols-1 sm:grid-cols-2 gap-2",
		},
		leftColumn,
		rightColumn,
	)

	// return
	return mainDiv
}
