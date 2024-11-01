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
	"github.com/chasefleming/elem-go/htmx"
	"github.com/gofiber/fiber/v2"
)

func CV(c *fiber.Ctx) error {
	// return the full page
	return components.PageHTML(c, CVContentHTML(middleware.Lang(c), CVDetailState(c)))
}

func CVContent(c *fiber.Ctx) error {
	lang := middleware.Lang(c) // Ensure you handle the potential absence of this value
	state := CVDetailState(c)  // Dynamically determine the state

	return c.SendString(CVContentHTML(lang, state).Render())
}

func CVFilename(lang string) string {
	return fmt.Sprintf("Daniel_CV_%s.pdf", lang)
}

func CVDetailState(c *fiber.Ctx) string {
	// Valid states
	validStates := map[string]bool{"collapsed": true, "expanded": true}

	// First, check for an overridden state in the request's Locals (for immediate changes)
	if overriddenState, ok := c.Locals("Switched-CV-Detail").(string); ok && validStates[overriddenState] {
		return overriddenState
	}

	// Next, check the cookie (for persisted state)
	if cookieState := utils.GetCVDetailCookie(c); validStates[cookieState] {
		return cookieState
	}

	// Default state if none is set
	return "collapsed"
}

func CVDetail(c *fiber.Ctx) error {
	currentState := CVDetailState(c)
	newState := "collapsed"
	if currentState == "collapsed" {
		newState = "expanded"
	}

	// Update the state in both the cookie and Locals for immediate effect
	utils.SetCVDetailCookie(c, newState)
	c.Locals("Switched-CV-Detail", newState)

	// No need to pass the state to CVContent; it will determine the state itself
	return CVContent(c)
}

func CVContentHTML(lang string, state string) *elem.Element {
	// Get content from YAML
	cv := structs.GetCurriculumLoc(lang)
	text := structs.GetLanguagesLoc(lang)

	// Basic props
	UlProps := attrs.Props{attrs.Class: "list-disc p-4"}
	HeadingProps := attrs.Props{attrs.Class: components.RtH("font-black")}
	TitleProps := attrs.Props{attrs.Class: "font-bold italic"}

	// Main div
	topDiv := elem.Div(
		attrs.Props{
			attrs.Class: "rounded-lg bg-almost-black-2 p-4",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "flex flex-row pb-2 justify-between",
			},
			elem.Div(
				attrs.Props{
					attrs.Class: "flex-1",
				},
				elem.H2(
					attrs.Props{
						attrs.Class: components.RtH("font-black text-left"),
					},
					elem.Text(text["cv_full_title"]),
				),
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "flex-auto",
				},
			),
			elem.Div(
				attrs.Props{
					attrs.Class: "flex-1 text-right",
				},
				elem.A(
					attrs.Props{
						attrs.Href:  constants.ResourceWithCDN(fmt.Sprintf("/cv/%s", CVFilename(lang))),
						attrs.Class: components.RtS("text-links hover:text-links-hover"),
					},
					elem.Span(
						nil,
						elem.Text(text["cv_download"]),
					),
				),
			),
		),
		elem.P(
			nil,
			elem.Text(text["cv_introduction"]),
		),
	)

	// Full CV (left) JOBS
	var jobsBlock []elem.Node

	// add title to jobs block
	cvDetailToggleText := map[string]string{
		"collapsed": text["cv_detail_toggle_show"],
		"expanded":  text["cv_detail_toggle_hide"],
	}[state]
	jobsBlock = append(jobsBlock,
		elem.Div(
			nil,
			elem.H2(
				HeadingProps,
				elem.Text(cv.Work.Heading),
				elem.A(
					attrs.Props{
						attrs.ID:      "cv-detail-toggle",
						attrs.Href:    "#",
						attrs.Class:   components.RtFooter("text-links-hover"),
						htmx.HXGet:    "/cv-detail",
						htmx.HXTarget: "#content",
					},
					elem.Text(fmt.Sprintf("&nbsp; %s", cvDetailToggleText)),
				),
			),
		),
	)

	// add jobs to jobsBlock
	for _, job := range cv.Work.List {
		// check if the state is expanded to add details to the jobsUl
		var jobDetail *elem.Element
		if state == "expanded" {
			jobDetail = elem.Div(
				attrs.Props{
					attrs.Class: components.RtFooterClasses,
				},
				elem.Br(nil),
				elem.P(nil, elem.Text(job.Detail)),
			)
		} else {
			jobDetail = elem.Span(nil)
		}

		// add jobs to jobsBlock
		jobsUl := elem.Ul(
			UlProps,
			elem.Li(
				nil,
				elem.Span(
					TitleProps,
					elem.Text(job.Title),
				),
			),
			elem.Span(
				nil,
				elem.Text(job.Institution),
				elem.Br(nil),
				elem.Text(job.Dates),
				elem.Br(nil),
				elem.Text(job.Location),
				elem.Br(nil),
				elem.Text(job.Schedule),
				elem.Br(nil),
			),
			jobDetail,
		)
		jobsBlock = append(jobsBlock, jobsUl)
	}

	leftCV := elem.Div(
		attrs.Props{
			attrs.Class: "left-column sm:pb-6 sm:pl-0 sm:pr-0 py-2 ",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "sm:mb-10 rounded-lg bg-almost-black-2 p-4 py-2",
			},
			jobsBlock...,
		),
	)

	// Full CV (right) EDUCATION + OTHERS
	var edBlock []elem.Node

	// add title to jobs block
	edBlock = append(edBlock,
		elem.Div(
			nil,
			elem.H2(
				HeadingProps,
				elem.Text(cv.Education.Heading),
			),
		),
	)

	// add jobs to edBlock
	for _, ed := range cv.Education.List {
		// check if the state is expanded to add details to the jobsUl
		var edDetail *elem.Element
		if state == "expanded" {
			edDetail = elem.Div(
				attrs.Props{
					attrs.Class: components.RtFooterClasses,
				},
				elem.Br(nil),
				elem.P(nil, elem.Text(ed.Detail)),
			)
		} else {
			edDetail = elem.Span(nil)
		}

		edUl := elem.Ul(
			UlProps,
			elem.Li(
				nil,
				elem.Span(
					TitleProps,
					elem.Text(ed.Title),
				),
			),
			elem.Span(
				nil,
				elem.Text(ed.Institution),
				elem.Br(nil),
				elem.Text(ed.Dates),
				elem.Br(nil),
				elem.Text(ed.Location),
				elem.Br(nil),
				elem.Text(ed.Schedule),
				elem.Br(nil),
			),
			edDetail,
		)
		edBlock = append(edBlock, edUl)
	}

	// Append languages and other skills to the edBlock
	var skillsBlock []elem.Node

	// add title to skills block
	skillsBlock = append(skillsBlock, elem.Div(
		nil,
		elem.H2(
			HeadingProps,
			elem.Text(cv.Other.Skills.Heading),
		),
	))

	// add skills to the skillsBlock
	skillsBlock = append(skillsBlock, elem.Ul(
		attrs.Props{attrs.Class: "list-disc p-4 space-y-4"},
		components.LiSpan(elem.Text(cv.Other.Skills.List.NoLevel.Languages)),
		components.LiSpan(elem.Text(cv.Other.Skills.List.NoLevel.Data)),
		components.LiSpan(elem.Text(cv.Other.Skills.List.NoLevel.Tools1)),
		components.LiSpan(elem.Text(cv.Other.Skills.List.NoLevel.Tools2)),
		components.LiSpan(elem.Text(cv.Other.Skills.List.NoLevel.Tools3)),
		components.LiSpan(elem.Text(cv.Other.Skills.List.NoLevel.Other1)),
		components.LiSpan(elem.Text(cv.Other.Skills.List.NoLevel.Other2)),
	))

	// Languages Block
	languagesBlock := elem.Div(
		nil,
		elem.Div(
			nil,
			elem.H2(
				HeadingProps,
				elem.Text(cv.Other.SkillsLangs.Heading),
			),
		),

		elem.Div(
			attrs.Props{
				attrs.Class: "pt-4",
			},
			elem.Span(
				nil,
				elem.Text(cv.Other.SkillsLangs.Level.Native),
			),
			elem.Ul(
				UlProps,
				components.LiSpan(elem.Text(strings.Join(cv.Other.SkillsLangs.List.Native, elem.Br(nil).Render()))),
			),
			elem.Span(
				nil,
				elem.Text(cv.Other.SkillsLangs.Level.Basic),
			),
			elem.Ul(
				UlProps,
				components.LiSpan(elem.Text(strings.Join(cv.Other.SkillsLangs.List.Basic, elem.Br(nil).Render()))),
			),
		),
	)

	// Right side of the CV
	rightCV := elem.Div(
		attrs.Props{
			attrs.Class: "right-column py-2 pb-6 sm:pl-0 sm:pr-0 space-y-2",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "rounded-lg bg-almost-black-2 py-2 p-4",
			},
			edBlock...,
		),
		elem.Div(
			attrs.Props{
				attrs.Class: "rounded-lg bg-almost-black-2 py-2 p-4",
			},
			skillsBlock...,
		),
		elem.Div(
			attrs.Props{
				attrs.Class: "rounded-lg bg-almost-black-2 py-2 p-4",
			},
			languagesBlock,
		),
	)

	// Main div
	mainDiv := elem.Div(
		attrs.Props{
			attrs.Class: "container mx-auto text-left pt-4",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "sm:pl-0 sm:pr-0",
			},
			topDiv,
		),
		elem.Div(
			attrs.Props{
				attrs.Class: "grid grid-cols-1 sm:grid-cols-2 sm:gap-2",
			},
			leftCV,
			rightCV,
		),
	)

	// return nil
	return mainDiv
}
