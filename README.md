# [GO Scraper]()

An oppinated & limited way to access the web using GO.

## Examples

Here are a few impressions on the way the library works. More examples are on the project go docs.

For complete working example please refer to `_examples` project directory.
```go
web := goscrapper.NewScrapper(fmt.Sprintf("http://localhost:%d", *port))

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