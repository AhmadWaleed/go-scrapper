package main

import "fmt"

func main() {
	//url := "https://google.com";
	url := "https://www.metalsucks.net/"

	web := NewScrapper(url)

	fmt.Println(web.Images()[0])
}
