package goscrapper

import (
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
	"testing"
)

func TestQuery(t *testing.T) {
	html := `
<html>
	<head>
		<title>My document</title>
	</head>
	<body>
		<article>
			<h1 class="post-title">
				Dee Snider Talks About the Role Musicians Play During the Pandemic
			</h1>
		</article>
<article>
			<h1 class="post-title">
				Grinding Through the Nonexistent Summer
			</h1>
		</article>
	</body>
</html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("could not create doc reader %v", err)
	}

	exp := []QueryResult{
		{
			Attr: map[string]interface{}{"class": "post-title"},
			Text: "Dee Snider Talks About the Role Musicians Play During the Pandemic",
		},
		{
			Attr: map[string]interface{}{"class": "post-title"},
			Text: "Grinding Through the Nonexistent Summer",
		},
	}

	query := Query{Name: "Post title", Selector: ".post-title"}

	web := Web{Doc: doc}
	got := web.Query(query)
	if eq := reflect.DeepEqual(got, exp); !eq {
		t.Errorf("query result should be: %v, got: %v", exp, got)
	}
}
