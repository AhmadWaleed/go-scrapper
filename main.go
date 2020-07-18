package main

func main() {
	//url := "https://stackoverflow.com/questions/32171498/how-to-get-value-of-attribute-href-value-in-go-language";
	url := "https://godoc.org/github.com/PuerkitoBio/goquery";

	web := NewScrapper(url)

	web.Fetch(url)
}