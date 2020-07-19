package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func NewScrapper(url string) *Web {
	web := Web{}
	web.URL = url

	if err := web.Fetch(); err != nil {
		log.Fatal(err)
	}

	return &web
}

type Web struct {
	URL          string
	Doc          *goquery.Document
	HeadingLevel HeadingLevel
}

func (w *Web) Fetch() error {
	res, err := http.Get(w.URL)
	if err != nil {
		return fmt.Errorf("could not get %s: %v", w.URL, err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusTooManyRequests {
			return fmt.Errorf("you are being rate limited")
		}

		return fmt.Errorf("bad response from server: %s", res.Status)
	}

	w.Doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("could not parse: %v", err)
	}

	return nil
}