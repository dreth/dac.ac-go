package components

import (
	"fmt"

	"dac.sg/structs"
	"dac.sg/utils"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
)

// NAVBAR BUTTONS
func NavbarButton(title string, link string, id string) *elem.Element {
	return elem.A(
		attrs.Props{
			attrs.ID:              fmt.Sprintf("%s-navbar-button", id),
			attrs.Href:            link,
			attrs.Class:           Rt("navbar-link inline-block no-underline  text-links m-0 mx-3 sm:mx-4 py-2 cursor-pointer hover:text-links-hover font-black"),
			htmx.HXGet:            fmt.Sprintf("%s/content", link),
			htmx.HXTarget:         "#content",
			htmx.HXPushURL:        link,
			htmx.HXOnAfterRequest: "window.scrollTo(0, 0);",
		},
		elem.Div(
			nil,
			elem.H3(nil, elem.Text(title)),
		),
	)
}

// NAVBAR
func Navbar(lang string) *elem.Element {
	// Get the langs data so that we can use it in the navbar
	text_en := structs.GetLanguagesLoc("en")
	text := structs.GetLanguagesLoc(lang)

	navbar := elem.Nav(
		attrs.Props{
			attrs.Class:  "navbar flex justify-center items-center text-center fixed bg-almost-black text-links hover:text-links-hover py-2 w-full z-50 shadow-md h-12 sm:h-14 bg-opacity-80",
			htmx.HXBoost: "true",
		},
		NavbarButton(text["nav_home"], "/home", utils.RSWD(text_en["nav_home"])),
		NavbarButton(text["nav_cv"], "/cv", utils.RSWD(text_en["nav_cv"])),
		NavbarButton(text["nav_blog"], "/blog", utils.RSWD(text_en["nav_blog"])),
		NavbarButton(text["nav_links"], "/cool-links", utils.RSWD(text_en["nav_links"])),
		NavbarButton(text["nav_projects"], "/projects", utils.RSWD(text_en["nav_projects"])),
		NavbarButton(text["nav_playlists"], "/playlists", utils.RSWD(text_en["nav_playlists"])),
	)

	return navbar
}
