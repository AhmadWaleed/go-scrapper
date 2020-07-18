package main

import "fmt"

func main() {
	url := "https://google.com";

	web := NewScrapper(url)

	fmt.Println(web.Title())
}