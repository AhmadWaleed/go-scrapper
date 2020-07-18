package main

import (
	"github.com/PuerkitoBio/goquery"
)

// Levels: h1, h2, h3, h4, h5, h6
type HeadingOption struct {
	Level string
}

func DefaultHeadingOption() HeadingOption {
	return HeadingOption{Level: "h1"}
}

func (w *Web) WithHeadingOption(opt HeadingOption) *Web {
	w.HeadingOption = opt
	return w
}

func NewHeadingOption(level string) HeadingOption {
	return HeadingOption{Level: level}
}

// Fetch slice of heading text, default level is h1
// you pass different heading level -> NewHeadingOption("h2")
//
// <h1>Heading 1</h1>
// w.Heading()[0] -> Heading 1
func (w *Web) Heading() []string {
	var headings []string
	w.Doc.Find(w.HeadingOption.Level).Next().Each(func(i int, heading *goquery.Selection) {
		headings = append(headings, heading.Text())
	})
	return headings
}

// Fetch slice of all the heading tags text (h1, h2, h3, h4, h5, h6)
//
// <h1>Heading 1</h1>
// <h2>Heading 2</h2>
// <h3>Heading 3</h3>
// w.Headings() -> [[Heading 1, Heading 2, Heading]]
func (w *Web) Headings() [][]string {
	opts := []HeadingOption{
		NewHeadingOption("h1"),
		NewHeadingOption("h2"),
		NewHeadingOption("h3"),
		NewHeadingOption("h4"),
		NewHeadingOption("h5"),
		NewHeadingOption("h6"),
	}

	var headings [][]string
	for _, opt := range opts {
		h := w.WithHeadingOption(opt).Heading()
		headings = append(headings, h)
	}

	return headings
}
