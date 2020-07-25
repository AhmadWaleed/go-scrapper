package goscrapper

import "github.com/PuerkitoBio/goquery"

// Fetch all the paragraphs (<p>) on a website
func (w *Web) Paragraphs() []string {
	paragraphs := make([]string, 0)
	w.Doc.Find("p").Each(func(i int, p *goquery.Selection) {
		paragraphs = append(paragraphs, p.Text())
	})
	return paragraphs
}

// Empty p-tags would lead to empty strings in the returned array.
// To avoid this you can call w.CleanParagraphs() instead.
// This will filter empty paragraphs and only return those with content.
func (w *Web) CleanParagraphs() []string {
	var cp []string
	for _, p := range w.Paragraphs() {
		if p != "" {
			cp = append(cp, p)
		}
	}
	return cp
}
