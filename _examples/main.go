package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/AhmadWaleed/go-scrapper"
	"log"
	"net/http"
	"time"
)

var port = flag.Int("port", 8080, "port which you want to start server on.")

func main() {
	flag.Parse()

	go TestServer(fmt.Sprintf(":%d", *port))

	ctx := context.Background();
	ctx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()

	web, err := goscrapper.NewContextScrapper(ctx, fmt.Sprintf("http://localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	//web := goscrapper.NewScrapper(fmt.Sprintf("http://localhost:%d", *port)) // without context

	// scrape headers info
	fmt.Println(web.Title())
	fmt.Println(web.CSRFToken())
	fmt.Println(web.ContentType())

	// scrape all headers
	fmt.Println(web.Headers())

	// scrape paragraphs
	fmt.Println(web.Paragraphs())
	fmt.Println(web.CleanParagraphs())

	// scrape images and links and commonly interesting details
	fmt.Println(web.Links())
	fmt.Println(web.InternalLinks())
	fmt.Println(web.ExternalLinks())
	fmt.Println(web.LinksWithDetails())
	fmt.Println(web.Images())
	fmt.Println(web.ImagesWithDetails())

	// scrape emails
	fmt.Println(web.Emails())

	// scrape using custom query
	quotes := web.Query(goscrapper.Query{Name: "Quotes", Selector: "quotes p"})
	for _, q := range quotes {
		fmt.Printf("Attributes: %v, Value: %v\n", q.Attr, q.Text)
	}
}

func TestServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, indexHTML)
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}

const indexHTML = `<!doctype html>
<html>
<head>
  <title>example</title>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
</head>
<body>
	<img src="https://test-page.com/assets/cat.jpg" alt="absolute path">
    <a href="https://placekitten.com/432/287" rel="nofollow">external kitten</a>
    <a href="mailto:jane@example.com"></a>
    <p>The quick brown fox jumps over the lazy dog.</p>
	<div class="quotes">
		<p>The greatest glory in living lies not in never falling, but in rising every time we fall.</p>
		<p>The way to get started is to quit talking and begin doing.</p>
	</div>
</body>
</html>`
