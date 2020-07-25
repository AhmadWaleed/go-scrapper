package goscrapper

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Query struct {
	Name     string
	Selector string
}

type QueryResult struct {
	Attr map[string]interface{}
	Text string
}

// get attributes and value of given query selector.
// return slice of result in case of multiple existence of an element
func (w *Web) Query(query Query) []QueryResult {
	var results []QueryResult
	w.Doc.Find(query.Selector).Each(func(i int, selection *goquery.Selection) {
		n := selection.Get(0)

		attr := make(map[string]interface{})
		for _, a := range n.Attr {
			attr[a.Key] = a.Val
		}

		result := QueryResult{
			Attr: attr,
			Text: strings.TrimSpace(selection.First().Text()),
		}

		results = append(results, result)
	})

	return results
}
