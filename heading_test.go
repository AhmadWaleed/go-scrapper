package main

import (
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
	"testing"
)

func TestHeading(t *testing.T) {
	tt := []struct {
		name    string
		levels  []HeadingLevel
		html    string
		heading [][]string
	}{
		{
			name: "fetches single h1 default heading",
			html: `
<html>
	<head>
	</head>
	<body>
		<h1>heading 1</h1>
	</body>
</html>`,
			heading: [][]string{{"heading 1"}},
		},
		{
			name: "fetches multiple h1 headings",
			html: `
<html>
	<head>
	</head>
	<body>
		<h1>heading 1</h1>
		<h1>heading 2</h1>
	</body>
</html>`,
			heading: [][]string{{"heading 1", "heading 2"}},
		},
		{
			name:   "fetches only given heading levels",
			levels: []HeadingLevel{H2},
			html: `
<html>
	<head>
	</head>
	<body>
		<h1>heading 1</h1>
		<h2>heading h2</h2>
	</body>
</html>`,
			heading: [][]string{{"heading h2"}},
		},
		{
			name:   "fetches multiple given heading levels",
			levels: []HeadingLevel{H2, H3},
			html: `
<html>
	<head>
	</head>
	<body>
		<h1>heading 1</h1>
		<h2>heading h2</h2>
		<h3>heading h3</h3>
	</body>
</html>`,
			heading: [][]string{{"heading h2"}, {"heading h3"}},
		},
		{
			name: "no heading",
			html: `
<html>
	<head>
	</head>
	<body>]
	</body>
</html>`,
			heading: make([][]string, 0),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.html))
			if err != nil {
				t.Fatalf("could not create doc reader %v", err)
			}

			web := Web{Doc: doc}
			heading := web.Heading(tc.levels...)
			if eq := reflect.DeepEqual(heading, tc.heading); !eq {
				t.Errorf("heading should be: %v, got: %v", tc.heading, heading)
			}
		})
	}
}

func TestHeadings(t *testing.T) {
	html := `
<html>
	<head>
	</head>
	<body>
		<h1>heading 1</h1>
		<h2>heading 2</h2>
		<h3>heading 3</h3>
		<h4>heading 4</h4>
		<h5>heading 5</h5>
		<h6>heading 6</h6>
	</body>
</html>`

	expected := [][]string{
		{"heading 1"},
		{"heading 2"},
		{"heading 3"},
		{"heading 4"},
		{"heading 5"},
		{"heading 6"},
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("could not create doc reader %v", err)
	}

	web := Web{Doc: doc}
	got := web.Headings()
	if eq := reflect.DeepEqual(got, expected); !eq {
		t.Errorf("got should be: %v, got: %v", got, expected)
	}
}
