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

func (w *Web) WithHeadingOption(lvl HeadingLevel) *Web {
	w.HeadingLevel = lvl
	return w
}

// Fetch slice of heading text, default level is h1
// you pass different heading level -> NewHeadingOption("h2")
//
// <h1>Heading 1</h1>
// w.Heading()[0] -> Heading 1
func (w *Web) Heading(opt ...HeadingLevel) []string {
	var levels []HeadingLevel
	levels = append(levels, opt...)
	if len(levels) == 0 {
		// set default level if not provided any
		levels = append(levels, H1)
	}

	headings := make([]string, 0)
	for _, lvl := range levels {
		w.Doc.Find(string(lvl)).Each(func(i int, heading *goquery.Selection) {
			headings = append(headings, heading.Text())
		})
	}

	return headings
}

// Fetch slice of all the heading tags text (h1, h2, h3, h4, h5, h6)
//
// <h1>Heading 1</h1>
// <h2>Heading 2</h2>
// <h3>Heading 3</h3>
// w.Headings() -> [[Heading 1, Heading 2, Heading]]
func (w *Web) Headings() [][]string {
	opts := []HeadingLevel{H1, H2, H3, H4, H5, H6}

	var headings [][]string
	for _, opt := range opts {
		h := w.WithHeadingOption(opt).Heading()
		headings = append(headings, h)
	}

	return headings
}
