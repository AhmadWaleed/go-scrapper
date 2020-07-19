package main

import (
	"fmt"
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

// Fetch content type meta info from head
//
// <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
// w.ContentType() -> [text/html, utf-8]
func (w *Web) ContentType() string {
	var values []string
	content, ok := w.Doc.Find("meta[http-equiv='Content-type']").Attr("content")
	if !ok {
		return ""
	}

	items := strings.Split(content, " ")
	for _, v := range items {
		if strings.Contains(v, "=") {
			kv := strings.Split(v, "=")
			values = append(values, strings.TrimSpace(strings.TrimRight(kv[1], ";")))
			continue
		}
		values = append(values, strings.TrimSpace(strings.TrimRight(v, ";")))
	}

	return strings.Join(values, ", ")
}

// Fetch canonical meta url from head
//
// <link rel="canonical" href="https://test-pages.goscrapper.de/page.html" />
// w.Canonical() -> https://test-pages.goscrapper.de/page.html
func (w *Web) Canonical() string {
	return w.Doc.Find("link[rel='canonical']").AttrOr("href", "")
}

// Fetch meta info of csrf token from head
//
// <meta name="csrf-token" content="token" />
// w.CSRFToken() -> token
func (w *Web) CSRFToken() string {
	return w.Doc.Find("meta[name='csrf-token']").AttrOr("content", "")
}

// get the header collected as an slice
func (w *Web) Headers() map[string]string {
	headers := make(map[string]string, 4)
	headers["charset"] = w.Charset()
	headers["canonical"] = w.Canonical()
	headers["contentType"] = w.ContentType()
	headers["csrfToken"] = w.CSRFToken()
	headers["viewport"] = w.Viewport().String()
	return headers
}

type Viewport struct {
	Val map[string]string
}

func NewViewport(doc *goquery.Document) *Viewport {
	v := Viewport{}
	v.Fetch(doc)
	return &v
}

func (v *Viewport) Fetch(doc *goquery.Document) {
	vp := make(map[string]string)
	values, ok := doc.Find("meta[name='viewport']").Attr("content")
	if ok {
		for _, item := range strings.Split(values, ",") {
			kv := strings.Split(item, "=")
			vp[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	v.Val = vp
}

// String representation of viewport
func (v *Viewport) String() string {
	if len(v.Val) == 0 {
		return ""
	}

	var vp []string
	for k, v := range v.Val {
		vp = append(vp, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(vp, ", ")
}
