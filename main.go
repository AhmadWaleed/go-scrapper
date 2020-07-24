package main

import "fmt"

func main() {
	//url := "https://google.com";
	url := "https://www.metalsucks.net/"

	web := NewScrapper(url)

	q := Query{Name: "Title", Selector: ".post-title"}

	fmt.Println(web.Query(q)[0])
}
