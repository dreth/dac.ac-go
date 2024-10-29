package pages

import (
	"fmt"
	"strings"

	"dac.sg/components"
	"dac.sg/constants"
	"dac.sg/middleware"
	"dac.sg/structs"
	"dac.sg/utils"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/gofiber/fiber/v2"
)

func Playlists(c *fiber.Ctx) error {
	// return the full page
	return components.PageHTML(c, PlaylistsContentHTML(middleware.Lang(c)))
}

func SpotifyLinkConstructor(id string) string {
	return fmt.Sprintf("%s%s", "https://open.spotify.com/playlist/", id)
}

func PlaylistsContent(c *fiber.Ctx) error {
	lang := middleware.Lang(c)
	return c.SendString(PlaylistsContentHTML(lang).Render())
}

func PlaylistDiv(playlistName string) *elem.Element {
	// get the playlists
	playlists := structs.GetPlaylists()
	playlistId, ok := playlists.Get(playlistName)
	utils.HandleErrBool(ok, false, "Playlist not found")

	// arrangement (center or left)

	// contruct the link
	link := SpotifyLinkConstructor(playlistId)
	playlistImagePath := constants.ResourceWithCDN(fmt.Sprintf("/assets/playlist_images/%s.png", strings.ToLower(playlistName)))

	playlistBlock := elem.Div(
		attrs.Props{
			attrs.Class: "w-2/3 h-auto relative", // Sets the size and positions relatively for the anchor and overlay
		},
		elem.A(
			attrs.Props{
				attrs.Class:  components.Rt("group inline-block no-underline sm:text-links cursor-pointer sm:hover:text-links-hover"),
				attrs.Href:   link,
				attrs.Target: "_blank",
			},
			elem.Span(
				attrs.Props{
					attrs.Class: "displayblock sm:hidden text-playlists-mobile",
				},
				elem.Text(playlistName),
			),
			elem.Img(
				attrs.Props{
					attrs.Src:   playlistImagePath,
					attrs.Class: "w-full h-auto", // Image takes the full width and auto height of its container
				},
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "absolute inset-0 bg-black bg-opacity-50 flex opacity-0 group-hover:opacity-100 transition duration-300 ease-in-out justify-center items-center",
				},
				elem.Span(
					attrs.Props{
						attrs.Class: components.RtH("text-white text-center"),
					},
					elem.Text(playlistName),
				),
			),
		),
	)

	return playlistBlock
}

func PlaylistsContentHTML(lang string) *elem.Element {
	// Get content from YAML
	text := structs.GetLanguagesLoc(lang)

	// About page
	content := elem.Div(
		attrs.Props{
			attrs.Class: "container mx-auto text-left pt-4",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "rounded-lg bg-almost-black-2 p-4 sm:pb-6 pt-4",
			},
			elem.H2(
				attrs.Props{
					attrs.Class: components.RtH("font-black mb-2"),
				},
				elem.Text(text["playlists_title"]),
			),
			elem.P(
				nil,
				elem.Text(text["playlists_introduction"]),
			),
		),
		elem.Div(
			attrs.Props{
				attrs.Class: "pb-6 py-2",
			},
			elem.Div(
				attrs.Props{
					attrs.Class: "grid grid-cols-1 sm:grid-cols-2 gap-2 justify-center items-center rounded-lg bg-almost-black-2 p-4 sm:py-10 space-y-2",
				},
				elem.Div(
					attrs.Props{
						attrs.Class: "left-column justify-center flex flex-wrap space-y-4",
					},
					PlaylistDiv("ambient"),
					PlaylistDiv("math"),
					PlaylistDiv("the sky was pink"),
				),
				elem.Div(
					attrs.Props{
						attrs.Class: "right-column justify-center flex flex-wrap space-y-4",
					},
					PlaylistDiv("post-rock"),
					PlaylistDiv("loved despite of great faults"),
				),
			),
		),
	)

	// return nil
	return content
}
