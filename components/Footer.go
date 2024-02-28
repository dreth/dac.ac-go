package components

import (
	"fmt"
	"strings"

	"dac.ac/middleware"
	"dac.ac/structs"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
)

func FooterLangSwitcher(lang string) *elem.Element {
	// Get the new language path and button text
	text := structs.GetLanguagesLoc(lang)
	langText := map[string]string{
		"es": text["lang_button_en"],
		"en": text["lang_button_es"],
	}[lang]

	// JSON data to send the next language as part of the request
	jsonData := fmt.Sprintf(`{"lang": "%s"}`, middleware.NextLang(lang))

	return elem.A(
		attrs.Props{
			attrs.Href:            "#", // Prevent default navigation
			attrs.Class:           Rt("inline-block no-underline text-links m-0 cursor-pointer hover:text-links-hover pl-5 transform xs:translate-y-px sm:translate-y-0.5"),
			htmx.HXPost:           "/toggle-lang", // Post to the server to change the language
			htmx.HXVals:           jsonData,
			htmx.HXBoost:          "true",
			htmx.HXOnAfterRequest: "window.location.reload()",
		},
		elem.Text(langText),
	)
}

func Footer(lang string) *elem.Element {
	return elem.Footer(
		attrs.Props{
			attrs.Class: RtFooter("footer text-center text-text p-2 flex justify-center items-center fixed z-10 w-full bottom-0 bg-almost-black h-11  bg-opacity-60"),
		},
		elem.Div(
			nil,
			elem.Text("© Daniel Alonso"),
			FooterSocialButton("Email", "/profiles/email"),
			FooterSocialButton("LinkedIn", "/profiles/linkedin"),
			FooterSocialButton("GitHub", "/profiles/github"),
			FooterSocialButton("X", "/profiles/x"),
			FooterLangSwitcher(lang),
		),
	)
}

func FooterSocialButton(network string, link string) *elem.Element {
	return elem.A(
		attrs.Props{
			attrs.Href:  link,
			attrs.Class: RtFooter("hover:text-gray-500"),
		},
		elem.Img(
			attrs.Props{
				attrs.Src:   fmt.Sprintf("/assets/icons/%s.svg", strings.ToLower(network)),
				attrs.Alt:   network,
				attrs.Class: RtFooter("w-5 h-auto inline-block align-middle ml-4 hover:text-gray-500"),
			},
		),
	)
}
