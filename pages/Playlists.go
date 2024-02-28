package pages

import (
	"fmt"
	"strings"

	"dac.ac/components"
	"dac.ac/middleware"
	"dac.ac/structs"
	"dac.ac/utils"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/gofiber/fiber/v2"
)

func Playlists(c *fiber.Ctx) error {
	// return HTML
	c.Type("html")

	// content
	normalizedPath := strings.TrimSuffix(c.Path(), "/")
	content := components.Content(normalizedPath)

	// Playlists page
	html := components.HTML("Daniel Alonso", "Daniel Alonso", middleware.Lang(c), false, content).Render()

	// return nil
	return c.SendString(html)
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
	playlistImagePath := fmt.Sprintf("/assets/playlist_images/%s.png", strings.ToLower(playlistName))

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
			attrs.Class: "container mx-auto text-left pt-3 md:pt-4 xl:pt-5",
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
		elem.Br(nil),
		elem.Br(nil),
		elem.Div(
			attrs.Props{
				attrs.Class: "flex flex-wrap justify-center items-center",
			},
			elem.Div(
				attrs.Props{
					attrs.Class: "flex flex-wrap justify-center w-full sm:w-1/2 sm:mb-10  space-y-2",
				},
				PlaylistDiv("ambient"),
				PlaylistDiv("math"),
				PlaylistDiv("the sky was pink"),
				PlaylistDiv("clusterfuck"),
				PlaylistDiv("nederlands"),
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "flex flex-wrap justify-center w-full sm:w-1/2 mb-10  space-y-2",
				},
				PlaylistDiv("post-rock"),
				PlaylistDiv("loved despite of great faults"),
				PlaylistDiv("pumped up walking up a tiny hill that isn't even particularly steep tbh"),
				PlaylistDiv("crying songs"),
				PlaylistDiv("when life was easier"),
			),
		),
	)

	// return nil
	return content
}
