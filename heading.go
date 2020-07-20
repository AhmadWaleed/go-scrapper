package main

import (
	"github.com/PuerkitoBio/goquery"
)

type HeadingLevel string

const (
	H1 HeadingLevel = "h1"
	H2 HeadingLevel = "h2"
	H3 HeadingLevel = "h3"
	H4 HeadingLevel = "h4"
	H5 HeadingLevel = "h5"
	H6 HeadingLevel = "h6"
)

// Fetch slice of heading text, default level is h1
// you pass different heading level -> NewHeadingOption("h2")
//
// <h1>Heading 1</h1>
// w.Heading()[0] -> Heading 1
func (w *Web) Heading(opt ...HeadingLevel) [][]string {
	var levels []HeadingLevel
	levels = append(levels, opt...)
	if len(levels) == 0 {
		// set default level if not provided any
		levels = append(levels, H1)
	}

	headings := make([][]string, 0)
	for _, lvl := range levels {
		var heading []string
		w.Doc.Find(string(lvl)).Each(func(i int, h *goquery.Selection) {
			heading = append(heading, h.Text())
		})

		if len(heading) > 0 {
			headings = append(headings, heading)
		}
	}

	return headings
}

// Fetch slice of all the heading tags text (h1, h2, h3, h4, h5, h6)
//
// <h1>Heading 1</h1>
// <h1>Heading 1</h1>
// <h2>Heading 2</h2>
// <h2>Heading 2</h2>
// w.Headings() -> [[Heading 1, heading 1], [Heading 2, Heading 2]]
func (w *Web) Headings() [][]string {
	opts := []HeadingLevel{H1, H2, H3, H4, H5, H6}
	return w.Heading(opts...)
}
