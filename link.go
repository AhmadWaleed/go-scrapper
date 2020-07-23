package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

func (w *Web) Links() []string {
	var links []string
	w.Doc.Find("a").Each(func(i int, a *goquery.Selection) {
		link := a.AttrOr("href", "")
		if validUrl(link) {
			links = append(links, link)
		}
	})

	return links
}

func (w *Web) LinksWithDetails() []map[string]interface{} {
	var links []map[string]interface{}
	for _, node := range w.Doc.Find("a").Nodes {

		details := make(map[string]interface{})
		for _, a := range node.Attr {
			switch a.Key {
			case "href":
				if !validUrl(a.Val) {
					continue
				}
				details["url"] = a.Val
			case "title":
				details["title"] = a.Val
			case "target":
				details["target"] = a.Val
			case "rel":
				details["rel"] = strings.ToLower(a.Val)
			}

			details["isNofollow"] = false
			if strings.Contains(a.Val, "nofollow") {
				details["isNofollow"] = true
			}

			details["isUGC"] = false
			if strings.Contains(a.Val, "ugc") {
				details["isUGC"] = true
			}

			details["isNoopener"] = false
			if strings.Contains(a.Val, "noopener") {
				details["isNoopener"] = true
			}

			details["isNoreferrer"] = false
			if strings.Contains(a.Val, "noreferrer") {
				details["isNoreferrer"] = true
			}
		}

		links = append(links, details)
	}

	return links
}

func validUrl(link string) bool {
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}

	u, err := url.Parse(link)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
