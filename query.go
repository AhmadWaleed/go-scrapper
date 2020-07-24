package main

import "github.com/PuerkitoBio/goquery"

type QueryResult struct {
	Attr map[string]interface{}
	Text string
}

type Query struct {
	Name     string
	Selector string
}

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
			Text: selection.First().Text(),
		}

		results = append(results, result)
	})

	return results
}
