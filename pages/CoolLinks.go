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
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func CoolLinks(c *fiber.Ctx) error {
	// return the full page
	return components.PageHTML(c, CoolLinksContentHTML(middleware.Lang(c)))
}

func CoolLinksContent(c *fiber.Ctx) error {
	// return nil
	lang := middleware.Lang(c)
	return c.SendString(CoolLinksContentHTML(lang).Render())
}

func CoolLinksList(kind string, amount int) []elem.Node {
	// array to contain the links
	var linksBlock []elem.Node

	// get the links
	links := structs.GetLinks()
	var linksList orderedmap.OrderedMap[string, string]
	var linkConstructor func(string) string

	// link constructor for each value
	// Also assign the linksList
	switch strings.ToLower(kind) {
	// Wikipedia
	case "wikipedia":
		linksList = links.Wikipedia
		linkConstructor = WikipediaLinkConstructor
	// Wiktionary
	case "wiktionary":
		linksList = links.Wiktionary
		linkConstructor = WiktionaryLinkConstructor
	// Other sites
	case "other":
		linksList = links.OtherSites
		linkConstructor = OtherSitesLinkConstructor
	// People I admire
	case "people":
		linksList = links.PeopleIAdmire
		linkConstructor = PeopleIAdmireLinkConstructor
	}

	// values of the linksList
	listKeys := make([]string, 0, amount)
	listValues := make([]string, 0, amount)

	cnt := 0
	for pair := linksList.Newest(); pair != nil && cnt < amount; pair = pair.Prev() {
		listKeys = append(listKeys, pair.Key)
		listValues = append(listValues, pair.Value)

		cnt++
	}

	// add links
	for i := 0; i < amount; i++ {
		// contruct the link
		link := linkConstructor(listValues[i])

		linksBlock = append(linksBlock, elem.Li(
			attrs.Props{
				attrs.Class: "custom-li-marker overflow-hidden break-words py-1",
			},
			elem.A(
				attrs.Props{
					attrs.Class:  components.Rt("no-underline  text-links m-0 mx-2.5  cursor-pointer hover:text-links-hover"),
					attrs.Href:   link,
					attrs.Target: "_blank",
				},
				elem.Text(listKeys[i]),
			)))
	}

	return linksBlock
}

func CoolLinksContentHTML(lang string) *elem.Element {
	// Get content from YAML
	text := structs.GetLanguagesLoc(lang)

	// Basic props
	UlProps := attrs.Props{attrs.Class: "list-disc pb-4 pl-4 pr-4 pt-2"}
	HeadingProps := attrs.Props{attrs.Class: components.RtH("font-black")}

	// length of arrays of links
	links := structs.GetLinks()
	wikipediaLinksCount := links.Wikipedia.Len()
	wiktionaryLinksCount := links.Wiktionary.Len()
	otherSitesLinksCount := links.OtherSites.Len()
	peopleIAdmireLinksCount := links.PeopleIAdmire.Len()

	// About page
	content := elem.Div(
		attrs.Props{
			attrs.Class: "container mx-auto text-left pt-3 md:pt-4 xl:pt-5",
		},
		elem.H2(
			attrs.Props{
				attrs.Class: components.RtH("font-black mb-2"),
			},
			elem.Text(text["cool_links_title"]),
		),
		elem.P(
			nil,
			elem.Text(text["cool_links_introduction"]),
		),
		elem.Br(nil),
		elem.Br(nil),
		elem.Div(
			attrs.Props{
				attrs.Class: "flex flex-wrap",
			},
			elem.Div(
				attrs.Props{
					attrs.Class: "w-full sm:w-1/2 overflow-hidden",
				},
				elem.H2(
					HeadingProps,
					elem.Text("Wikipedia"),
				),
				elem.Ul(
					UlProps,
					CoolLinksList("wikipedia", wikipediaLinksCount)...,
				),
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "w-full sm:w-1/2 mb-4 overflow-hidden",
				},
				elem.H2(
					HeadingProps,
					elem.Text("Wiktionary"),
				),
				elem.Ul(
					UlProps,
					CoolLinksList("wiktionary", wiktionaryLinksCount)...,
				),
				elem.H2(
					HeadingProps,
					elem.Text("Other sites"),
				),
				elem.Ul(
					UlProps,
					CoolLinksList("other", otherSitesLinksCount)...,
				),
				elem.H2(
					HeadingProps,
					elem.Text("People I admire"),
				),
				elem.Ul(
					UlProps,
					CoolLinksList("people", peopleIAdmireLinksCount)...,
				),
			),
		),
	)

	// return nil
	return content
}

// WikipediaLinkConstructor takes a rune and constructs a Wikipedia link.
func WikipediaLinkConstructor(name string) string {
	const wikipediaBaseLink = "wikipedia.org/wiki/"
	if len(name) < 3 {
		return ""
	}
	articleLang := name[:2]
	articleName := name[3:]
	return fmt.Sprintf("https://%s.%s%s", articleLang, wikipediaBaseLink, articleName)
}

// WiktionaryLinkConstructor takes a rune and constructs a Wiktionary link.
func WiktionaryLinkConstructor(name string) string {
	const wiktionaryBaseLink = "wiktionary.org/wiki/"
	if len(name) < 3 {
		return ""
	}
	articleLang := name[:2]
	articleName := name[3:]
	return fmt.Sprintf("https://%s.%s%s", articleLang, wiktionaryBaseLink, articleName)
}

// OtherSitesLinkConstructor takes a rune and returns it (for direct links).
func OtherSitesLinkConstructor(link string) string {
	return link
}

// PeopleIAdmireLinkConstructor takes a rune and returns it (for direct links).
func PeopleIAdmireLinkConstructor(link string) string {
	return link
}
