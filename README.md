# [GO Scraper](https://godoc.org/github.com/AhmadWaleed/go-scrapper) [![PkgGoDev](https://pkg.go.dev/badge/github.com/AhmadWaleed/go-scrapper?tab=doc)](https://pkg.go.dev/github.com/AhmadWaleed/go-scrapper?tab=doc)

An oppinated & limited way to access the web using GO.

## Examples

Here are a few impressions on the way the library works. More examples are on the project go docs.

For complete working example please refer to `_examples` project directory.

### Initialization
```go
// with context
ctx := context.Background();
ctx, cancel := context.WithTimeout(ctx, time.Second * 5)
defer cancel()

web := goscrapper.NewContextScrapper(ctx, "https://www.domain.com")
if err != nil {
    log.Fatal(err)
}

// without context
web := goscrapper.NewScrapper("https://www.domain.com")
if err != nil {
    log.Fatal(err)
}
```

### Usage
```go
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
```

See the [full documentation](https://godoc.org/github.com/AhmadWaleed/go-scrapper) for more information and examples.
