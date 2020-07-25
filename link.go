package goscrapper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/weppos/publicsuffix-go/publicsuffix"
	"net/url"
	"strings"
)

// get slice of all links on the page as absolute URLs
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

// get all links on the page with commonly interesting details
//
// Example:
//
// html: `<a href="https://placekitten.com/432/287" rel="nofollow">external kitten</a>`
// Result: [
//     'url' => 'https://placekitten.com/432/287',
//     'text' => 'external kitten',
//     'title' => null,
//     'target' => null,
//     'rel' => 'nofollow',
//     'isNofollow' => true,
//     'isUGC' => false,
//     'isNoopener' => false,
//     'isNoreferrer' => false,
// ]
func (w *Web) LinksWithDetails() []map[string]interface{} {
	var links []map[string]interface{}
	w.Doc.Find("a").Each(func(i int, a *goquery.Selection) {
		details := make(map[string]interface{})
		link, ok := a.Attr("href")
		if ok {
			details["url"] = link

			title, ok := a.Attr("title")
			if ok {
				details["title"] = title
			}

			target, ok := a.Attr("target")
			if ok {
				details["target"] = target
			}

			rel, ok := a.Attr("rel")
			if ok {
				details["rel"] = rel
				if details["isUGC"] = true; strings.Contains(rel, "ugc") {
					details["isUGC"] = false
				}
				if details["isNoopener"] = true; strings.Contains(rel, "noopener") {
					details["isNoopener"] = false
				}
				if details["isNofollow"] = true; strings.Contains(rel, "nofollow") {
					details["isNofollow"] = false
				}
				if details["isNoreferrer"] = true; strings.Contains(rel, "noreferrer") {
					details["isNoreferrer"] = false
				}
			} else {
				details["isUGC"] = false
				details["isNoopener"] = false
				details["isNofollow"] = false
				details["isNoreferrer"] = false
			}
		}
		links = append(links, details)
	})

	return links
}

// get all internal links (same root or sub-domain) on the page as absolute URLs
func (w *Web) InternalLinks() []string {
	var links []string
	for _, link := range w.Links() {
		if internalLink(link, w.URL) {
			links = append(links, link)
		}
	}
	return links
}

// get all external links on the page as absolute URLs
func (w *Web) ExternalLinks() []string {
	var links []string
	for _, link := range w.Links() {
		if !internalLink(link, w.URL) {
			links = append(links, link)
		}
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

func internalLink(link string, currentURL string) bool {
	cDomain, err := publicsuffix.Parse(currentURL)
	if err != nil {
		return false
	}

	lDomain, err := publicsuffix.Parse(link)
	if err != nil {
		return false
	}

	if cDomain.SLD == lDomain.SLD {
		return true
	}

	return false
}
