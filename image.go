package main

import (
	"github.com/PuerkitoBio/goquery"
)

func (w *Web) Images() []string {
	var images []string
	w.Doc.Find("img").Each(func(i int, img *goquery.Selection) {
		url, ok := img.Attr("src")
		if ok {
			images = append(images, url)
		}
	})

	return images
}

func (w *Web) ImagesWithDetails() []map[string]interface{} {
	var images []map[string]interface{}

	nodes := w.Doc.Find("img").Nodes
	for _, node := range nodes {
		details := make(map[string]interface{})
		for _, n := range node.Attr {
			if allowed(n.Key) {
				details[n.Key] = n.Val
			}
		}
		images = append(images, details)
	}
	return images
}

func allowed(attr string) bool {
	switch attr {
	case
		"alt",
		"src",
		"width",
		"height":
		return true
	}
	return false
}