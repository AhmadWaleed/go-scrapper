package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

func TestTitle(t *testing.T) {
	tt := []struct{
		name string
		html string
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
				t.Fatalf("count not create doc reader %v", err)
			}

			web := Web{Doc: doc}
			if title := web.Title(); title != tc.title {
				t.Errorf("title should be %v got %v", tc.title, title)
			}
		})
	}

}
func TestCharset(t *testing.T) {
	tt := []struct{
		name string
		html string
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
				t.Fatalf("count not create doc reader %v", err)
			}

			web := Web{Doc: doc}
			if charset := web.Charset(); charset != tc.charset {
				t.Errorf("charset should be %v got %v", tc.charset, charset)
			}
		})
	}
}
