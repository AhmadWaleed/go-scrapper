package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func NewScrapper(url string) *Web {
	web := Web{}
	web.URL = url
	web.HeadingOption = DefaultHeadingOption()
	return &web
}

type Web struct {
	URL           string
	Doc           *goquery.Document
	HeadingOption HeadingOption
}

func (w *Web) Fetch(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not get %w: %v", url, err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusTooManyRequests {
			return fmt.Errorf("you are being rate limited")
		}

		return fmt.Errorf("bad response from server: %w", res.Status)
	}

	w.Doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("could not parse: %v", err)
	}

	return nil
}