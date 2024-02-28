package pages

import (
	"fmt"
	"os"
	"path"
	"strings"

	"dac.ac/components"
	"dac.ac/structs"
	"dac.ac/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/gofiber/fiber/v2"
	"github.com/russross/blackfriday/v2"
	"github.com/sourcegraph/syntaxhighlight"
)

func Article(c *fiber.Ctx) error {
	// return HTML
	c.Type("html")

	// get the article parameter from the URL
	article := c.Params("article")

	// content
	articleHtml := ArticleContentHTML(article)
	content := components.ContentArticle(articleHtml)

	// article list to get the title from
	articles := structs.GetArticles()
	articleData, present := articles.Get(article)
	var articleTitle string
	if present {
		articleTitle = articleData.Title
	}

	// Blog page
	html := components.HTML(articleTitle, articleTitle, utils.GetLang(c), true, content).Render()

	// return nil
	return c.SendString(html)
}

func ArticleContentHTML(article string) *elem.Element {
	// Read the article content from the file
	content, err := os.ReadFile(path.Join("static", "articles", article, "article.md"))
	utils.HandleErr(err, false, fmt.Sprintf("error reading %s", article))

	// additional extensions for rendering
	blackfridayExtensions := blackfriday.WithExtensions(blackfriday.CommonExtensions | blackfriday.AutoHeadingIDs)

	// convert the markdown to HTML
	mdToHTML := blackfriday.Run(content, blackfridayExtensions)

	// apply styling
	htmlContent := ArticleHTMLStyling(mdToHTML)

	return elem.Div(
		attrs.Props{
			attrs.Class: "p-4 overflow-wrap break-words overflow-x-auto",
		},
		elem.Text(htmlContent),
	)
}

func ArticleHTMLStyling(mdConversion []byte) string {
	// Parse the HTML content using goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(mdConversion)))
	utils.HandleErr(err, false, "error parsing HTML content")

	// HEADINGS AND OTHER TAGS
	// Mapping of tags to classes
	tagClasses := map[string]string{
		"h1":         components.H1Classes,
		"h2":         components.H2Classes,
		"h3":         components.H3Classes,
		"h4":         components.H4Classes,
		"h5":         components.H5Classes,
		"a":          components.AClasses,
		"p":          components.PClasses,
		"ul":         components.UlClasses,
		"ol":         components.OlClasses,
		"li":         components.LiClasses,
		"hr":         components.HrClasses,
		"pre":        components.PreClasses,
		"blockquote": components.BlockquoteClasses,
		"figure":     components.FigureClasses,
		"figcaption": components.FigcaptionClasses,
	}

	// Loop through H1-H5 elements and apply classes
	for tag, class := range tagClasses {
		doc.Find(tag).Each(func(j int, s *goquery.Selection) {
			s.AddClass(class)
		})
	}

	// IMAGE TAGS
	// Find images outside figures without the smaller class
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		// Check if the image is inside a figure
		insideFigure := s.ParentsFiltered("figure").Length() > 0

		// Check if the image has the "smaller" class
		hasSmallerClass := s.HasClass("smaller")

		// Remove "smaller" class if it exists
		if hasSmallerClass {
			s.RemoveClass("smaller")
		}

		// Add appropriate classes based on conditions
		switch {
		case insideFigure && hasSmallerClass:
			s.AddClass(components.ImgClassesSmallerInsideFigure)
		case insideFigure && !hasSmallerClass:
			s.AddClass(components.ImgClassesInsideFigure)
		case !insideFigure && hasSmallerClass:
			s.AddClass(components.ImgClassesSmallerOutsideFigure)
		case !insideFigure && !hasSmallerClass:
			s.AddClass(components.ImgClassesOutsideFigure)
		}
	})

	// FIGURE TAGS
	// Add figure classes
	doc.Find("figure").Each(func(i int, s *goquery.Selection) {
		s.AddClass(components.FigureClasses)
	})

	// CODE RELATED QUERIES
	// Code highlighting
	doc.Find("code[class*=\"language-\"]").Each(func(i int, s *goquery.Selection) {
		oldCode := s.Text()
		formatted, err := syntaxhighlight.AsHTML([]byte(oldCode))
		utils.HandleErr(err, true, "error formatting code")
		s.SetHtml(string(formatted))
	})

	// FINALIZE CODE
	// Get the modified HTML content as a string
	htmlContent, err := doc.Html()
	utils.HandleErr(err, false, "error getting modified HTML content")

	// Remove unnecessary tags
	htmlContent = strings.Replace(htmlContent, "<html><head></head><body>", "", 1)
	htmlContent = strings.Replace(htmlContent, "</body></html>", "", 1)

	return htmlContent
}

func ArticleList(amount int) []elem.Node {
	// array to contain the links
	var articlesBlock []elem.Node

	// get the article data
	articles := structs.GetArticles()

	// values of the articlesList
	listKeys := make([]string, 0, amount)
	listValues := make([]structs.Article, 0, amount)

	// iterate over article
	cnt := 0
	for pair := articles.Oldest(); pair != nil && cnt < amount; pair = pair.Next() {
		listKeys = append(listKeys, pair.Key)
		listValues = append(listValues, pair.Value)

		cnt++
	}

	// iterate over the list
	for i := 0; i < len(listKeys); i++ {
		// get the article info
		article := listValues[i]
		articleTopInfo := fmt.Sprintf("%s - %s", article.Date, article.Lang)
		articleHref := fmt.Sprintf("/blog/%s", listKeys[i])
		articleTitle := fmt.Sprintf("%s %s", article.Emoji, article.Title)

		// append the link
		articlesBlock = append(articlesBlock,
			elem.Div(
				attrs.Props{
					attrs.Class: "pb-3 sm:pb-5",
				},
				elem.P(
					attrs.Props{
						attrs.Class: components.RtFooter("pt-2 md:pt-1"),
					},
					elem.Text(articleTopInfo),
				),
				elem.A(
					attrs.Props{
						attrs.Href:  articleHref,
						attrs.Class: components.Rt("text-blog-links hover:text-blog-links-hover"),
					},
					elem.Text(articleTitle),
				),
			),
		)
	}

	// return the articleBlock
	return articlesBlock
}
