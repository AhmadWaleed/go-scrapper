package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// Fetch the title from head, if a tag wasn't found because
// it's missing in the source HTML, empty string will be returned.
//
// <title>Lorem Ipsum</title>
// w.Title() -> Lorem Ipsum
func (w *Web) Title() string {
	return strings.TrimSpace(w.Doc.Find("title").Text())
}

// Fetch the charset meta info from head, if a tag wasn't found because
// it's missing in the source HTML, empty string will be returned.
//
// <meta charset="utf-8" />
// w.Title() -> utf-8
func (w *Web) Charset() string {
	return w.Doc.Find("meta").AttrOr("charset", "")
}

// Fetch viewport meta info from head
//
// <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
// w.Viewport().Val -> ['width=device-width', 'initial-scale=1', 'maximum-scale=1', 'user-scalable=no']
// w.Viewport().String() -> 'width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no'
func (w *Web) Viewport() *Viewport {
	return NewViewport(w.Doc)
}

// Fetch content type meta info from head, slice will be returned
//
// <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
// w.ContentType() -> [text/html, utf-8]
func (w *Web) ContentType() []string {
	var types []string
	n := w.Doc.Find("meta[http-equiv='Content-type']").Get(0)
	for _, a := range n.Attr {
		types = append(types, a.Val)
	}

	return types
}

// Fetch canonical meta url from head
//
// <link rel="canonical" href="https://test-pages.goscrapper.de/page.html" />
// w.Canonical() -> https://test-pages.goscrapper.de/page.html
func (w *Web) Canonical() string {
	return w.Doc.Find("link[rel='canonical']").AttrOr("content", "")
}


// Fetch meta info of csrf token from head
//
// <meta name="csrf-token" content="token" />
// w.CSRFToken() -> token
func (w *Web) CSRFToken() string {
	return w.Doc.Find("meta[name='csrf-token']").AttrOr("content", "")
}

// get the header collected as an slice
func (w *Web) Headers() map[string]interface{} {
	headers := make(map[string]interface{}, 4)
	headers["charset"] = w.Charset()
	headers["contentType"] = w.ContentType()
	headers["csrfToken"] = w.CSRFToken()
	headers["viewport"] = w.Viewport().Val
	return headers
}

type Viewport struct {
	Val []string
}

func NewViewport(doc *goquery.Document) *Viewport {
	v := Viewport{}
	v.Fetch(doc)
	return &v
}

func (v *Viewport) Fetch(doc *goquery.Document) {
	n := doc.Find("meta[name='viewport']").Get(0)
	for _, a := range n.Attr{
		if a.Val == "viewport" {
			continue
		}
		v.Val = append(v.Val, a.Val)
	}
}

// String representation of viewport
func (v *Viewport) String() string {
	if len(v.Val) == 0 {
		return ""
	}
	return strings.Join(v.Val, ", ")
}