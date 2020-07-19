package main

import (
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
	"testing"
)

func TestTitle(t *testing.T) {
	tt := []struct {
		name  string
		html  string
		title string
	}{
		{
			name: "with title",
			html: `
<html>
	<head>
		<title>Lorem Ipsum</title>
	</head>
	<body>
	</body>
</html>`,
			title: "Lorem Ipsum",
		},
		{
			name: "no title",
			html: `
<html>
	<head>
	</head>
	<body>
	</body>
</html>`,
			title: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.html))
			if err != nil {
				t.Fatalf("could not create doc reader %v", err)
			}

			web := Web{Doc: doc}
			if title := web.Title(); title != tc.title {
				t.Errorf("title should be %v got %v", tc.title, title)
			}
		})
	}

}

func TestCharset(t *testing.T) {
	tt := []struct {
		name    string
		html    string
		charset string
	}{
		{
			name: "with charset",
			html: `
<html>
	<head>
		<meta charset="utf-8" />
	</head>
	<body>
	</body>
</html>`,
			charset: "utf-8",
		},
		{
			name: "no charset",
			html: `
<html>
	<head>
	</head>
	<body>
	</body>
</html>`,
			charset: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.html))
			if err != nil {
				t.Fatalf("could not create doc reader %v", err)
			}

			web := Web{Doc: doc}
			if charset := web.Charset(); charset != tc.charset {
				t.Errorf("charset should be %v got %v", tc.charset, charset)
			}
		})
	}
}

func TestCanonical(t *testing.T) {
	tt := []struct {
		name      string
		html      string
		canonical string
	}{
		{
			name: "with canonical url",
			html: `
<html>
	<head>
		<link rel="canonical" href="https://test-pages.goscrapper.de/page.html" />
	</head>
	<body>
	</body>
</html>`,
			canonical: "https://test-pages.goscrapper.de/page.html",
		},
		{
			name: "no canonical url",
			html: `
<html>
	<head>
	</head>
	<body>
	</body>
</html>`,
			canonical: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.html))
			if err != nil {
				t.Fatalf("could not create doc reader %v", err)
			}

			web := Web{Doc: doc}
			if canonical := web.Canonical(); canonical != tc.canonical {
				t.Errorf("canonical should be %v got %v", tc.canonical, canonical)
			}
		})
	}
}

func TestCSRFToken(t *testing.T) {
	tt := []struct {
		name  string
		html  string
		token string
	}{
		{
			name: "with csrf token",
			html: `
<html>
	<head>
		<meta name="csrf-token" content="token" />
	</head>
	<body>
	</body>
</html>`,
			token: "token",
		},
		{
			name: "no csrf token",
			html: `
<html>
	<head>
	</head>
	<body>
	</body>
</html>`,
			token: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.html))
			if err != nil {
				t.Fatalf("could not create doc reader %v", err)
			}

			web := Web{Doc: doc}
			if token := web.CSRFToken(); token != tc.token {
				t.Errorf("canonical should be %v got %v", tc.token, token)
			}
		})
	}
}

func TestViewport(t *testing.T) {
	tt := []struct {
		name     string
		html     string
		viewport map[string]string
	}{
		{
			name: "with viewport",
			html: `
<html>
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1" />
	</head>
	<body>
	</body>
</html>`,
			viewport: map[string]string{
				"width":         "device-width",
				"initial-scale": "1",
			},
		},
		{
			name: "no viewport",
			html: `
<html>
	<head>
	</head>
	<body>
	</body>
</html>`,
			viewport: make(map[string]string),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.html))
			if err != nil {
				t.Fatalf("could not create doc reader %v", err)
			}

			web := Web{Doc: doc}
			viewport := web.Viewport().Val
			viewportStr := web.Viewport().String()
			eq := reflect.DeepEqual(viewport, tc.viewport)
			if len(viewport) > 0 && viewportStr != "width=device-width, initial-scale=1" {
				t.Errorf("viewport should be width=device-width, initial-scale=1 got %v", viewportStr)
			}
			if !eq {
				t.Errorf("viewport should be %v got %v", tc.viewport, viewport)
			}
		})
	}
}
