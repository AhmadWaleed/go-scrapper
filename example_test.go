package goscrapper_test

import (
	"fmt"
	"goscrapper"
)

func Example() {
	web := goscrapper.NewScrapper("https://www.domain.com")

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

func ExampleWeb_Title() {
	web := goscrapper.NewScrapper("https://www.metalsucks.net/")

	fmt.Println(web.Title())
	// Output: MetalSucks | Metal News, Tour Dates, Reviews and Videos
}

func ExampleWeb_Query() {
	web := goscrapper.NewScrapper("https://www.metalsucks.net/")

	metaResult := web.Query(goscrapper.Query{Name: "Meta Info", Selector: "meta[property='og:locale']"})
	fmt.Println(metaResult[0].Attr)
	//Output: map[content:en_US property:og:locale]
}
