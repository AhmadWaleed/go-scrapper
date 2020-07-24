package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
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
			<h1 class=".post-title">
				Dee Snider Talks About the Role Musicians Play During the Pandemic
			</h1>
			<p class=".post-body">
				I’ve gotten perspective on what I do, entertainment, art, and I realize that we really are a distraction, we’re not that important.
			</p>
		</article>
<article>
			<h1 class=".post-title">
				Grinding Through the Nonexistent Summer
			</h1>
			<p class=".post-body">
				A look at the state of the metal world, how lives have so drastically changed during the pandemic, and behind-the-scenes at MetalSucks.
			</p>
		</article>
	</body>
</html>`

	query := Query{Name: "Post title", Selector: ".post-title"}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("could not create doc reader %v", err)
	}

	web := Web{Doc: doc}
	results := web.Query(query)
	fmt.Println(results)
	// TODO: complete test
}
