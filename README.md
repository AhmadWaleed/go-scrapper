# [GO Scraper]()

An oppinated & limited way to access the web using GO.

## Examples

Here are a few impressions on the way the library works. More examples are on the project go docs.

Fetching the meta info of a web page:
```go
web := NewScrapper("https://www.google.com/")

// prints "google"
fmt.Println(web.Title())

// <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
// return ["text/html utf-8]"
fmt.Println(web.ContentType())
```