package main

import (
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
	"testing"
)

func TestParagraphs(t *testing.T) {
	tt := []struct{
		name string
		html string
		paragraph []string
	}{
		{
			name: "fetches single paragraph",
			html: `
<html>
	<head>
	</head>
	<body>
		<p>The quick brown fox jumps over the lazy dog.</p>
	</body>
</html>`,
			paragraph:  []string{"The quick brown fox jumps over the lazy dog."},
		},
		{
			name: "fetches multiple paragraph",
			html: `
<html>
	<head>
	</head>
	<body>
		<p>The quick brown fox jumps over the lazy dog.</p>
		<p>Nymphs blitz quick vex dwarf jog.</p>
	</body>
</html>`,
			paragraph:  []string{
				"The quick brown fox jumps over the lazy dog.",
				"Nymphs blitz quick vex dwarf jog.",
			},
		},
		{
			name: "no paragraph return empty slice",
			html: `
<html>
	<head>
	</head>
	<body>
	</body>
</html>`,
			paragraph:  make([]string, 0),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.html))
			if err != nil {
				t.Fatalf("could not create doc reader %v", err)
			}

			web := Web{Doc: doc}
			paragraph := web.Paragraphs()
			if eq := reflect.DeepEqual(paragraph, tc.paragraph); !eq {
				t.Errorf("paragraph should be: %v, got: %v", tc.paragraph, paragraph)
			}
		})
	}
}

func TestCleanParagraph(t *testing.T) {
	html := `
<html>
	<head>
	</head>
	<body>
		<p>The quick brown fox jumps over the lazy dog.</p>
		<p></p>
	</body>
</html>`

	exp := []string{"The quick brown fox jumps over the lazy dog."}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("could not create doc reader %v", err)
	}

	web := Web{Doc: doc}
	got := web.CleanParagraphs()
	if eq := reflect.DeepEqual(got, exp); !eq {
		t.Errorf("paragraph should be: %v, got: %v", exp, got)
	}
}