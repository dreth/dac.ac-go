package main

import (
	"os"
	"path/filepath"

	"dac.ac/components"
	"dac.ac/pages"
	"dac.ac/structs"
)

func RenderComponents() {
	// components directory
	comps_dir := filepath.Join("generated")
	os.MkdirAll(comps_dir, 0755)

	os.WriteFile(filepath.Join(comps_dir, "index.html"), []byte(components.HTMLForTailwind().Render()), 0644)
	os.WriteFile(filepath.Join(comps_dir, "about-content.html"), []byte(pages.CVContentHTML("en", "expanded").Render()), 0644)
	os.WriteFile(filepath.Join(comps_dir, "homepage-content.html"), []byte(pages.HomepageContentHTML("en").Render()), 0644)
	os.WriteFile(filepath.Join(comps_dir, "cool-links-content.html"), []byte(pages.CoolLinksContentHTML("en").Render()), 0644)
	os.WriteFile(filepath.Join(comps_dir, "blog-content.html"), []byte(pages.BlogContentHTML("en").Render()), 0644)
	os.WriteFile(filepath.Join(comps_dir, "playlists-content.html"), []byte(pages.PlaylistsContentHTML("en").Render()), 0644)

	// articles
	articles := structs.GetArticles()

	for pair := articles.Newest(); pair != nil; pair = pair.Prev() {
		article := pair.Key
		os.WriteFile(filepath.Join(comps_dir, article+".html"), []byte(pages.ArticleContentHTML(article).Render()), 0644)
	}
}
